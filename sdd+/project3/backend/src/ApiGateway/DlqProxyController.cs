using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using System.Threading.Tasks;
using System.Net.Http;
using System.Net.Http.Headers;

namespace SDD.ApiGateway.Controllers
{
    [ApiController]
    [Route("api/admin/dlq")]
    [Authorize(Roles = "Admin")] // T023: Implement JWT role check (Admin access)
    public class DlqProxyController : ControllerBase
    {
        private readonly IHttpClientFactory _httpClientFactory;
        private readonly string _dlqServiceUrl = "http://dead-letter-service:8080"; // Internal K8s DNS name

        public DlqProxyController(IHttpClientFactory httpClientFactory)
        {
            _httpClientFactory = httpClientFactory;
        }

        [HttpGet]
        public async Task<IActionResult> GetDlq([FromQuery] string? originalQueueName, [FromQuery] int page = 1, [FromQuery] int pageSize = 20)
        {
            // T022: Configure Gateway routing to DeadLetterService
            var client = _httpClientFactory.CreateClient();
            var requestUrl = $"{_dlqServiceUrl}/api/admin/dlq?page={page}&pageSize={pageSize}";
            if (!string.IsNullOrEmpty(originalQueueName))
            {
                requestUrl += $"&originalQueueName={originalQueueName}";
            }

            var response = await client.GetAsync(requestUrl);
            return StatusCode((int)response.StatusCode, await response.Content.ReadAsStringAsync());
        }

        [HttpGet("{id}")]
        public async Task<IActionResult> GetDlqItem(int id)
        {
            var client = _httpClientFactory.CreateClient();
            var response = await client.GetAsync($"{_dlqServiceUrl}/api/admin/dlq/{id}");
            return StatusCode((int)response.StatusCode, await response.Content.ReadAsStringAsync());
        }

        [HttpPost("{id}/replay")]
        public async Task<IActionResult> ReplayDlqItem(int id)
        {
            var client = _httpClientFactory.CreateClient();
            var response = await client.PostAsync($"{_dlqServiceUrl}/api/admin/dlq/{id}/replay", null);
            return StatusCode((int)response.StatusCode, await response.Content.ReadAsStringAsync());
        }

        [HttpDelete("{id}/purge")]
        public async Task<IActionResult> PurgeDlqItem(int id)
        {
            var client = _httpClientFactory.CreateClient();
            var response = await client.DeleteAsync($"{_dlqServiceUrl}/api/admin/dlq/{id}/purge");
            return StatusCode((int)response.StatusCode, await response.Content.ReadAsStringAsync());
        }
    }
}
