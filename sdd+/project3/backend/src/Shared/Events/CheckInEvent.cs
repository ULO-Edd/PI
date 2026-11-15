using System;

namespace SDD.Shared.Events
{
    public class CheckInEvent : BaseEvent
    {
        public string AttendeeId { get; set; } = string.Empty;
        public string EventId { get; set; } = string.Empty;
        public DateTime CheckInTime { get; set; }

        public CheckInEvent()
        {
            EventType = "check_in";
        }
    }
}
