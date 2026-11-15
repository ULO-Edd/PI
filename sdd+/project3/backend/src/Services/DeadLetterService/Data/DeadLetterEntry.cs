using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;

namespace SDD.Services.DeadLetterService.Data
{
    public enum DeadLetterStatus
    {
        PendingAction,
        Replayed,
        Purged
    }

    public class DeadLetterEntry
    {
        [Key]
        public int Id { get; set; }
        public Guid MessageId { get; set; }
        public Guid CorrelationId { get; set; }
        public string EventType { get; set; } = string.Empty;
        public string OriginalQueueName { get; set; } = string.Empty;
        public string TargetConsumer { get; set; } = string.Empty;
        public string FailureReason { get; set; } = string.Empty;
        public string? ErrorStackTrace { get; set; }
        public string Payload { get; set; } = string.Empty;
        public DateTime FailureTimestamp { get; set; } = DateTime.UtcNow;
        public string RetryHistory { get; set; } = "[]"; // JSON array of retry attempts
        public DeadLetterStatus Status { get; set; } = DeadLetterStatus.PendingAction;
    }
}
