using System;
using System.Text;
using System.Text.Json;
using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using SDD.Shared.Events;
using SDD.Shared;

namespace SDD.Services.AnalyticsWorker
{
    public class CheckInEventConsumer
    {
        public void ProcessEvent(BasicDeliverEventArgs ea, Action<CheckInEvent> onProcessed)
        {
            try
            {
                var body = ea.Body.ToArray();
                var message = Encoding.UTF8.GetString(body);
                var checkInEvent = JsonSerializer.Deserialize<CheckInEvent>(message);

                if (checkInEvent == null)
                {
                    throw new Exception("Failed to deserialize CheckInEvent payload.");
                }

                onProcessed(checkInEvent);
            }
            catch (Exception ex)
            {
                throw new Exception($"Consumer processing error: {ex.Message}", ex);
            }
        }
    }
}
