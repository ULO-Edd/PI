using System.Security.Claims;
using System.Text;
using System.Threading.RateLimiting;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.AspNetCore.RateLimiting;
using Microsoft.IdentityModel.Tokens;

var builder = WebApplication.CreateBuilder(args);

// Add HttpClient for proxy forwarding
builder.Services.AddHttpClient();

// Load Jwt Configuration
var jwtSecret = builder.Configuration["JwtSettings:Secret"] ?? "SuperSecretPerimeterAPIKey2026!YourStrongSecurityKeyHere";
var jwtIssuer = builder.Configuration["JwtSettings:Issuer"] ?? "SDDProject3Identity";
var jwtAudience = builder.Configuration["JwtSettings:Audience"] ?? "SDDProject3Gateway";

builder.Services.AddAuthentication(options =>
{
    options.DefaultAuthenticateScheme = JwtBearerDefaults.AuthenticationScheme;
    options.DefaultChallengeScheme = JwtBearerDefaults.AuthenticationScheme;
})
.AddJwtBearer(options =>
{
    options.TokenValidationParameters = new TokenValidationParameters
    {
        ValidateIssuer = true,
        ValidateAudience = true,
        ValidateLifetime = true,
        ValidateIssuerSigningKey = true,
        ValidIssuer = jwtIssuer,
        ValidAudience = jwtAudience,
        IssuerSigningKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(jwtSecret))
    };
});

builder.Services.AddAuthorization();

// Implement Role-Based Rate Limiting Policies globally using PartitionedRateLimiter (FR-003)
builder.Services.AddRateLimiter(options =>
{
    options.RejectionStatusCode = StatusCodes.Status429TooManyRequests;

    options.GlobalLimiter = PartitionedRateLimiter.Create<HttpContext, string>(context =>
    {
        var role = context.User?.FindFirst(ClaimTypes.Role)?.Value ?? "";
        var isOrganiser = role == "Organiser" || role == "Admin";
        var clientIp = context.Connection.RemoteIpAddress?.ToString() ?? "unknown";

        if (isOrganiser)
        {
            // Organiser Limits: 1000 requests per minute
            return RateLimitPartition.GetFixedWindowLimiter(
                partitionKey: $"{clientIp}-organiser",
                factory: _ => new FixedWindowRateLimiterOptions
                {
                    PermitLimit = 1000,
                    QueueLimit = 50,
                    QueueProcessingOrder = QueueProcessingOrder.OldestFirst,
                    Window = TimeSpan.FromMinutes(1)
                });
        }
        else
        {
            // Attendee/Public Limits: 100 requests per minute
            return RateLimitPartition.GetFixedWindowLimiter(
                partitionKey: $"{clientIp}-attendee",
                factory: _ => new FixedWindowRateLimiterOptions
                {
                    PermitLimit = 100,
                    QueueLimit = 10,
                    QueueProcessingOrder = QueueProcessingOrder.OldestFirst,
                    Window = TimeSpan.FromMinutes(1)
                });
        }
    });
});

var app = builder.Build();

app.UseAuthentication();
app.UseAuthorization();
app.UseRateLimiter();

// Map and forward paths under Custom Reverse Proxy middleware
app.Run(async context =>
{
    var path = context.Request.Path.Value ?? "";

    // 1. Determine target downstream service port and credentials
    string targetServiceUrl;
    if (path.StartsWith("/api/auth"))
    {
        targetServiceUrl = "http://localhost:5001";
    }
    else if (path.StartsWith("/api/events") || path.StartsWith("/api/organiser"))
    {
        targetServiceUrl = "http://localhost:5002";
    }
    else if (path.StartsWith("/api/registrations"))
    {
        targetServiceUrl = "http://localhost:5003";
    }
    else
    {
        context.Response.StatusCode = StatusCodes.Status404NotFound;
        await context.Response.WriteAsync("Service route not found at Gateway.");
        return;
    }

    // 2. Enforce Authentication Guards on secured endpoints (FR-002)
    var isPublicEndpoint = path == "/api/auth/login" || path == "/api/auth/register" || (context.Request.Method == "GET" && path == "/api/events");
    if (!isPublicEndpoint)
    {
        var authenticated = context.User?.Identity?.IsAuthenticated ?? false;
        if (!authenticated)
        {
            context.Response.StatusCode = StatusCodes.Status401Unauthorized;
            await context.Response.WriteAsync("Unauthorized: Invalid or missing token.");
            return;
        }

        // Role restriction checks for Organiser routes
        if (path.StartsWith("/api/organiser"))
        {
            var userRole = context.User?.FindFirst(ClaimTypes.Role)?.Value ?? "";
            if (userRole != "Organiser" && userRole != "Admin")
            {
                context.Response.StatusCode = StatusCodes.Status403Forbidden;
                await context.Response.WriteAsync("Forbidden: Restricted to organisers.");
                return;
            }
        }
    }

    // 3. Copy and forward HTTP Request Message
    var httpClientFactory = context.RequestServices.GetRequiredService<IHttpClientFactory>();
    using var client = httpClientFactory.CreateClient();

    var targetUri = new Uri($"{targetServiceUrl}{context.Request.Path}{context.Request.QueryString}");
    using var proxyRequest = new HttpRequestMessage(new HttpMethod(context.Request.Method), targetUri);

    // Copy request body
    if (context.Request.ContentLength > 0 || context.Request.Headers.ContainsKey("Transfer-Encoding"))
    {
        var streamContent = new StreamContent(context.Request.Body);
        proxyRequest.Content = streamContent;
        if (context.Request.ContentType != null)
        {
            proxyRequest.Content.Headers.ContentType = System.Net.Http.Headers.MediaTypeHeaderValue.Parse(context.Request.ContentType);
        }
    }

    // Copy request headers
    foreach (var header in context.Request.Headers)
    {
        if (header.Key.StartsWith("Content-") && proxyRequest.Content != null)
        {
            proxyRequest.Content.Headers.TryAddWithoutValidation(header.Key, header.Value.ToArray());
        }
        else
        {
            proxyRequest.Headers.TryAddWithoutValidation(header.Key, header.Value.ToArray());
        }
    }

    // Forward the request and capture response
    try
    {
        using var responseMessage = await client.SendAsync(proxyRequest, HttpCompletionOption.ResponseHeadersRead, context.RequestAborted);
        
        // Copy response headers and body back to client
        context.Response.StatusCode = (int)responseMessage.StatusCode;
        foreach (var header in responseMessage.Headers)
        {
            context.Response.Headers[header.Key] = header.Value.ToArray();
        }
        foreach (var header in responseMessage.Content.Headers)
        {
            context.Response.Headers[header.Key] = header.Value.ToArray();
        }

        // Copy body content
        await responseMessage.Content.CopyToAsync(context.Response.Body, context.RequestAborted);
    }
    catch (Exception ex)
    {
        context.Response.StatusCode = StatusCodes.Status502BadGateway;
        await context.Response.WriteAsync($"Gateway Error connecting downstream service: {ex.Message}");
    }
});

app.Run();
