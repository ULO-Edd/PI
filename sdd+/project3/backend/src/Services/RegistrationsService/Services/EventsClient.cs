using System.Net.Http.Json;

namespace RegistrationsService.Services;

public class EventsClient
{
    private readonly HttpClient _httpClient;

    public EventsClient(HttpClient httpClient, IConfiguration configuration)
    {
        _httpClient = httpClient;
        // Load target Events Service URL from config, falling back to default port 5002
        _httpClient.BaseAddress = new Uri(configuration["EventsService:BaseUrl"] ?? "http://localhost:5002");
    }

    // Synchronously request EventsService to reserve ticket capacity (T013)
    public async Task<bool> ReserveCapacityAsync(Guid eventId, int ticketsCount, string correlationId)
    {
        try
        {
            var request = new HttpRequestMessage(HttpMethod.Post, $"/api/events/{eventId}/reserve");
            request.Headers.Add("X-Correlation-Id", correlationId);
            request.Content = JsonContent.Create(new { ticketsCount, correlationId });

            var response = await _httpClient.SendAsync(request);
            return response.IsSuccessStatusCode;
        }
        catch (Exception ex)
        {
            Console.WriteLine($"HTTP inter-service reserve request failed: {ex.Message}");
            return false;
        }
    }

    // Synchronously request EventsService to release/rollback ticket capacity
    public async Task<bool> ReleaseCapacityAsync(Guid eventId, int ticketsCount, string correlationId)
    {
        try
        {
            var request = new HttpRequestMessage(HttpMethod.Post, $"/api/events/{eventId}/release");
            request.Headers.Add("X-Correlation-Id", correlationId);
            request.Content = JsonContent.Create(new { ticketsCount, correlationId });

            var response = await _httpClient.SendAsync(request);
            return response.IsSuccessStatusCode;
        }
        catch (Exception ex)
        {
            Console.WriteLine($"HTTP inter-service release request failed: {ex.Message}");
            return false;
        }
    }
}
// Definition of done (T013)
