using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using SDD.Services.DeadLetterService.Services;
using SDD.Services.DeadLetterService.Data;

namespace SDD.Services.DeadLetterService.Controllers
{
    [ApiController]
    [Route("api/admin/dlq")]
    public class DeadLetterController : ControllerBase
    {
        private readonly IFailureManager _failureManager;

        public DeadLetterController(IFailureManager failureManager)
        {
            _failureManager = failureManager;
        }

        [HttpGet]
        public async Task<IActionResult> GetFailedMessages([FromQuery] string? originalQueueName, [FromQuery] int page = 1, [FromQuery] int pageSize = 20)
        {
            var (items, total) = await _failureManager.GetFailures(originalQueueName, page, pageSize);
            return Ok(new { totalItems = total, page, pageSize, items });
        }

        [HttpGet("{id}")]
        public async Task<IActionResult> GetFailedMessageById(int id)
        {
            var entry = await _failureManager.GetFailureById(id);
            if (entry == null) return NotFound();
            return Ok(entry);
        }

        [HttpPost("{id}/replay")]
        public async Task<IActionResult> ReplayFailedMessage(int id)
        {
            var success = await _failureManager.MarkAsReplayed(id);
            if (!success) return NotFound("Message not found or failed to replay.");
            
            var entry = await _failureManager.GetFailureById(id);
            return Ok(new { success = true, message = "Message successfully replayed", correlationId = entry?.CorrelationId });
        }

        [HttpDelete("{id}/purge")]
        public async Task<IActionResult> PurgeFailedMessage(int id)
        {
            var success = await _failureManager.MarkAsPurged(id);
            if (!success) return NotFound();
            return Ok(new { success = true, message = "Message successfully purged" });
        }
    }
}
