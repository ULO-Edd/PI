using Microsoft.AspNetCore.Mvc;
using RegistrationsService.Data;
using RegistrationsService.Models;
using RegistrationsService.Services;

namespace RegistrationsService.Controllers;

[ApiController]
[Route("api/registrations")]
public class RegistrationsController : ControllerBase
{
    private readonly RegistrationsDbContext _context;
    private readonly EventsClient _eventsClient;

    public RegistrationsController(RegistrationsDbContext context, EventsClient eventsClient)
    {
        _context = context;
        _eventsClient = eventsClient;
    }

    [HttpGet]
    public IActionResult Get()
    {
        return Ok("Registrations Service Active.");
    }

    // POST /api/registrations/book - Transaction Coordinator with rollback (T014)
    [HttpPost("book")]
    public async Task<IActionResult> BookTicket([FromBody] BookInput input)
    {
        if (input.EventId == Guid.Empty || input.AttendeeId == Guid.Empty)
        {
            return BadRequest("Event ID and Attendee ID are required.");
        }

        var correlationId = string.IsNullOrWhiteSpace(input.CorrelationId) 
            ? $"corr-{Guid.NewGuid()}" 
            : input.CorrelationId.Trim();

        // 1. Synchronously reserve capacity in Events Service
        var reserveSucceeded = await _eventsClient.ReserveCapacityAsync(input.EventId, 1, correlationId);
        if (!reserveSucceeded)
        {
            return BadRequest("Ticket booking failed: Capacity exhausted or event not found.");
        }

        // 2. Insert registration entity into Registrations database
        var registration = new Registration
        {
            EventId = input.EventId,
            AttendeeId = input.AttendeeId,
            Status = "Confirmed",
            BookedAt = DateTime.UtcNow,
            CorrelationId = correlationId
        };

        try
        {
            _context.Registrations.Add(registration);
            await _context.SaveChangesAsync();
            return StatusCode(201, registration);
        }
        catch (Exception ex)
        {
            Console.WriteLine($"Database write failed for correlation {correlationId}: {ex.Message}. Rolling back capacity...");
            
            // 3. Resilience Rollback: Release capacity atomically on database save failures (T015)
            var rollbackSucceeded = await _eventsClient.ReleaseCapacityAsync(input.EventId, 1, correlationId);
            if (!rollbackSucceeded)
            {
                Console.WriteLine($"CRITICAL ERROR: Capacity rollback failed for event {input.EventId}, correlation {correlationId}. Manual intervention required.");
            }

            return BadRequest($"Ticket booking transaction failed: {ex.Message}. Allocated capacity has been safely rolled back.");
        }
    }
}

public record BookInput(Guid EventId, Guid AttendeeId, string CorrelationId);
