using System;

namespace SDD.Shared.Events
{
    public abstract class BaseEvent
    {
        public Guid MessageId { get; set; } = Guid.NewGuid();
        public Guid CorrelationId { get; set; }
        public DateTime Timestamp { get; set; } = DateTime.UtcNow;
        public string EventType { get; set; } = string.Empty;
    }
}
