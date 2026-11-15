using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;
using IdentityService.Data;
using IdentityService.Models;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Microsoft.IdentityModel.Tokens;

namespace IdentityService.Controllers;

[ApiController]
[Route("api/auth")]
public class AuthController : ControllerBase
{
    private readonly IdentityDbContext _context;
    private readonly IConfiguration _configuration;

    public AuthController(IdentityDbContext context, IConfiguration configuration)
    {
        _context = context;
        _configuration = configuration;
    }

    [HttpPost("register")]
    public async Task<IActionResult> Register([FromBody] RegisterInput input)
    {
        if (string.IsNullOrWhiteSpace(input.Email) || string.IsNullOrWhiteSpace(input.Password))
        {
            return BadRequest("Email and Password are required.");
        }

        var normalizedEmail = input.Email.Trim().ToLowerInvariant();
        var exists = await _context.Users.AnyAsync(u => u.Email == normalizedEmail);
        if (exists)
        {
            return BadRequest("Email is already registered.");
        }

        var user = new User
        {
            Email = normalizedEmail,
            PasswordHash = BCrypt.Net.BCrypt.HashPassword(input.Password),
            Role = input.Role == "Organiser" || input.Role == "Admin" ? input.Role : "Attendee"
        };

        _context.Users.Add(user);
        await _context.SaveChangesAsync();

        return CreatedAtAction(nameof(GetSession), new { }, new { user = new { user.Id, user.Email, user.Role } });
    }

    [HttpPost("login")]
    public async Task<IActionResult> Login([FromBody] LoginInput input)
    {
        if (string.IsNullOrWhiteSpace(input.Email) || string.IsNullOrWhiteSpace(input.Password))
        {
            return BadRequest("Email and Password are required.");
        }

        var normalizedEmail = input.Email.Trim().ToLowerInvariant();
        var user = await _context.Users.FirstOrDefaultAsync(u => u.Email == normalizedEmail);

        if (user == null || !BCrypt.Net.BCrypt.Verify(input.Password, user.PasswordHash))
        {
            // Execute dummy verification to prevent email enumeration timing attacks
            BCrypt.Net.BCrypt.HashPassword("dummy_password");
            return Unauthorized("Invalid email or password.");
        }

        var token = GenerateJwtToken(user);

        // Append token as a secure, HTTP-only cookie
        Response.Cookies.Append("authToken", token, new CookieOptions
        {
            HttpOnly = true,
            Secure = false, // Set to true in prod environments
            SameSite = SameSiteMode.Lax,
            Expires = DateTimeOffset.UtcNow.AddDays(1)
        });

        return Ok(new { user = new { user.Id, user.Email, user.Role }, token });
    }

    [HttpPost("logout")]
    public IActionResult Logout()
    {
        Response.Cookies.Delete("authToken");
        return Ok(new { message = "Logged out successfully" });
    }

    [HttpGet("session")]
    public IActionResult GetSession()
    {
        // Check for token in cookies or Authorization header
        var token = Request.Cookies["authToken"];
        if (string.IsNullOrEmpty(token) && Request.Headers.TryGetValue("Authorization", out var authHeader))
        {
            var headerValue = authHeader.ToString();
            if (headerValue.StartsWith("Bearer ", StringComparison.OrdinalIgnoreCase))
            {
                token = headerValue.Substring(7);
            }
        }

        if (string.IsNullOrEmpty(token))
        {
            return Unauthorized("No active session found.");
        }

        try
        {
            var tokenHandler = new JwtSecurityTokenHandler();
            var jwtSecret = _configuration["JwtSettings:Secret"] ?? "SuperSecretPerimeterAPIKey2026!YourStrongSecurityKeyHere";
            var key = Encoding.UTF8.GetBytes(jwtSecret);

            tokenHandler.ValidateToken(token, new TokenValidationParameters
            {
                ValidateIssuer = true,
                ValidateAudience = true,
                ValidateLifetime = true,
                ValidateIssuerSigningKey = true,
                ValidIssuer = _configuration["JwtSettings:Issuer"] ?? "SDDProject3Identity",
                ValidAudience = _configuration["JwtSettings:Audience"] ?? "SDDProject3Gateway",
                IssuerSigningKey = new SymmetricSecurityKey(key)
            }, out SecurityToken validatedToken);

            var jwtToken = (JwtSecurityToken)validatedToken;
            var userId = jwtToken.Claims.First(x => x.Type == ClaimTypes.NameIdentifier).Value;
            var email = jwtToken.Claims.First(x => x.Type == ClaimTypes.Email).Value;
            var role = jwtToken.Claims.First(x => x.Type == ClaimTypes.Role).Value;

            return Ok(new { user = new { id = userId, email, role } });
        }
        catch
        {
            return Unauthorized("Session is invalid or expired.");
        }
    }

    private string GenerateJwtToken(User user)
    {
        var tokenHandler = new JwtSecurityTokenHandler();
        var jwtSecret = _configuration["JwtSettings:Secret"] ?? "SuperSecretPerimeterAPIKey2026!YourStrongSecurityKeyHere";
        var key = Encoding.UTF8.GetBytes(jwtSecret);

        var tokenDescriptor = new SecurityTokenDescriptor
        {
            Subject = new ClaimsIdentity(new[]
            {
                new Claim(ClaimTypes.NameIdentifier, user.Id.ToString()),
                new Claim(ClaimTypes.Email, user.Email),
                new Claim(ClaimTypes.Role, user.Role)
            }),
            Expires = DateTime.UtcNow.AddDays(1),
            Issuer = _configuration["JwtSettings:Issuer"] ?? "SDDProject3Identity",
            Audience = _configuration["JwtSettings:Audience"] ?? "SDDProject3Gateway",
            SigningCredentials = new SigningCredentials(new SymmetricSecurityKey(key), SecurityAlgorithms.HmacSha256Signature)
        };

        var token = tokenHandler.CreateToken(tokenDescriptor);
        return tokenHandler.WriteToken(token);
    }
}

public record RegisterInput(string Email, string Password, string Role);
public record LoginInput(string Email, string Password);
