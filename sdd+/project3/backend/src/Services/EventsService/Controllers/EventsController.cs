using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using EventsService.Data;
using EventsService.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;

namespace EventsService.Controllers;

[ApiController]
public class EventsController : ControllerBase
{
    private readonly EventsDbContext _context;

    public EventsController(EventsDbContext context)
    {
        _context = context;
    }

    // GET /api/events - Public catalog search and filter
    [HttpGet("api/events")]
    public async Task<IActionResult> GetPublicEvents([FromQuery] string? search)
    {
        var query = _context.Events.AsQueryable();

        // Only return events with remaining tickets
        query = query.Where(e => e.AvailableTickets > 0);

        if (!string.IsNullOrWhiteSpace(search))
        {
            var searchClean = search.Trim().ToLowerInvariant();
            query = query.Where(e => e.Title.ToLower().Contains(searchClean) || 
                                     e.Description.ToLower().Contains(searchClean));
        }

        var results = await query.ToListAsync();
        return Ok(results);
    }

    // GET /api/organiser/events - Organiser owned events list
    [HttpGet("api/organiser/events")]
    public async Task<IActionResult> GetOrganiserEvents()
    {
        var organiserId = GetUserIdFromToken();
        if (organiserId == Guid.Empty)
        {
            return Unauthorized("Unauthorized access.");
        }

        var results = await _context.Events.Where(e => e.OrganiserId == organiserId).ToListAsync();
        return Ok(results);
    }

    // POST /api/organiser/events - Create new event (with validations)
    [HttpPost("api/organiser/events")]
    public async Task<IActionResult> CreateEvent([FromBody] CreateEventInput input)
    {
        var organiserId = GetUserIdFromToken();
        if (organiserId == Guid.Empty)
        {
            return Unauthorized("Unauthorized access.");
        }

        if (string.IsNullOrWhiteSpace(input.Title))
        {
            return BadRequest("Title is required.");
        }

        if (input.Title.Length > 100)
        {
            return BadRequest("Title cannot exceed 100 characters.");
        }

        if (input.DateTime <= DateTime.UtcNow)
        {
            return BadRequest("Event Date and Time must be in the future.");
        }

        if (input.Capacity < 1)
        {
            return BadRequest("Capacity must be at least 1.");
        }

        var newEvent = new Event
        {
            Title = input.Title.Trim(),
            Description = input.Description?.Trim() ?? string.Empty,
            DateTime = input.DateTime,
            Location = input.Location?.Trim() ?? "Location TBA",
            Capacity = input.Capacity,
            AvailableTickets = input.Capacity, // Initially all tickets available
            OrganiserId = organiserId
        };

        _context.Events.Add(newEvent);
        await _context.SaveChangesAsync();

        return CreatedAtAction(nameof(GetPublicEvents), new { }, newEvent);
    }

    // PUT /api/organiser/events/{id} - Edit event details (with validations)
    [HttpPut("api/organiser/events/{id}")]
    public async Task<IActionResult> UpdateEvent(Guid id, [FromBody] UpdateEventInput input)
    {
        var organiserId = GetUserIdFromToken();
        if (organiserId == Guid.Empty)
        {
            return Unauthorized("Unauthorized access.");
        }

        var existingEvent = await _context.Events.FirstOrDefaultAsync(e => e.Id == id && e.OrganiserId == organiserId);
        if (existingEvent == null)
        {
            return NotFound("Event not found.");
        }

        if (string.IsNullOrWhiteSpace(input.Title))
        {
            return BadRequest("Title is required.");
        }

        if (input.Title.Length > 100)
        {
            return BadRequest("Title cannot exceed 100 characters.");
        }

        if (input.Capacity < 1)
        {
            return BadRequest("Capacity must be at least 1.");
        }

        // Available tickets must be adjusted in sync with capacity change
        var allocatedCount = existingEvent.Capacity - existingEvent.AvailableTickets;
        if (input.Capacity < allocatedCount)
        {
            return BadRequest("Capacity cannot be reduced below currently allocated tickets.");
        }

        existingEvent.Title = input.Title.Trim();
        existingEvent.Description = input.Description?.Trim() ?? string.Empty;
        existingEvent.DateTime = input.DateTime;
        existingEvent.Location = input.Location?.Trim() ?? "Location TBA";
        existingEvent.AvailableTickets = input.Capacity - allocatedCount;
        existingEvent.Capacity = input.Capacity;

        await _context.SaveChangesAsync();
        return Ok(existingEvent);
    }

    // DELETE /api/organiser/events/{id} - Delete event
    [HttpDelete("api/organiser/events/{id}")]
    public async Task<IActionResult> DeleteEvent(Guid id)
    {
        var organiserId = GetUserIdFromToken();
        if (organiserId == Guid.Empty)
        {
            return Unauthorized("Unauthorized access.");
        }

        var existingEvent = await _context.Events.FirstOrDefaultAsync(e => e.Id == id && e.OrganiserId == organiserId);
        if (existingEvent == null)
        {
            return NotFound("Event not found.");
        }

        _context.Events.Remove(existingEvent);
        await _context.SaveChangesAsync();
        return NoContent();
    }

    // POST /api/events/{id}/reserve - Synchronously validate and deduct ticket capacity atomically (T012)
    [HttpPost("api/events/{id}/reserve")]
    public async Task<IActionResult> ReserveCapacity(Guid id, [FromBody] ReserveInput input)
    {
        if (input.TicketsCount < 1)
        {
            return BadRequest("Tickets count must be at least 1.");
        }

        // Atomic DB SQL subtraction to prevent overbooking concurrency race conditions
        var rowsAffected = await _context.Events
            .Where(e => e.Id == id && e.AvailableTickets >= input.TicketsCount)
            .ExecuteUpdateAsync(s => s.SetProperty(e => e.AvailableTickets, e => e.AvailableTickets - input.TicketsCount));

        if (rowsAffected == 0)
        {
            return BadRequest("Capacity exhausted or event not found.");
        }

        return Ok(new { message = "Capacity reserved successfully" });
    }

    // POST /api/events/{id}/release - Release/Rollback capacity atomically
    [HttpPost("api/events/{id}/release")]
    public async Task<IActionResult> ReleaseCapacity(Guid id, [FromBody] ReleaseInput input)
    {
        if (input.TicketsCount < 1)
        {
            return BadRequest("Tickets count must be at least 1.");
        }

        // Atomic DB SQL addition to restore ticket capacity safely
        var rowsAffected = await _context.Events
            .Where(e => e.Id == id && e.AvailableTickets + input.TicketsCount <= e.Capacity)
            .ExecuteUpdateAsync(s => s.SetProperty(e => e.AvailableTickets, e => e.AvailableTickets + input.TicketsCount));

        if (rowsAffected == 0)
        {
            return BadRequest("Failed to release capacity: capacity overflow or event not found.");
        }

        return Ok(new { message = "Capacity released successfully" });
    }

    // Placeholder endpoint for attendee details (called by Organiser Dashboard)
    [HttpGet("api/organiser/events/{id}/attendees")]
    public IActionResult GetEventAttendees(Guid id)
    {
        return Ok(new List<object>());
    }

    private Guid GetUserIdFromToken()
    {
        var authHeader = Request.Headers["Authorization"].ToString();
        if (string.IsNullOrEmpty(authHeader))
        {
            // Fallback to custom proxy header if Jwt validation was offloaded to Gateway
            var proxyUserId = Request.Headers["X-User-Id"].ToString();
            if (Guid.TryParse(proxyUserId, out var parsedGuid))
            {
                return parsedGuid;
            }
            return Guid.Empty;
        }

        if (!authHeader.StartsWith("Bearer ", StringComparison.OrdinalIgnoreCase))
        {
            return Guid.Empty;
        }

        try
        {
            var token = authHeader.Substring(7);
            var tokenHandler = new JwtSecurityTokenHandler();
            var jwtToken = tokenHandler.ReadJwtToken(token);
            var userIdClaim = jwtToken.Claims.First(x => x.Type == ClaimTypes.NameIdentifier).Value;
            return Guid.Parse(userIdClaim);
        }
        catch
        {
            return Guid.Empty;
        }
    }
}

public record CreateEventInput(string Title, string Description, DateTime DateTime, string Location, int Capacity);
public record UpdateEventInput(string Title, string Description, DateTime DateTime, string Location, int Capacity);
public record ReserveInput(int TicketsCount, string CorrelationId);
public record ReleaseInput(int TicketsCount, string CorrelationId);
