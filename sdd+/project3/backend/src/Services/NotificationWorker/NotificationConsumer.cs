using System;
using System.Text;
using System.Text.Json;
using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using SDD.Shared.Events;
using SDD.Shared;

namespace SDD.Services.NotificationWorker
{
    public class NotificationConsumer
    {
        private readonly string _apiKey;
        private readonly string _baseUrl;

        public NotificationConsumer(string apiKey, string baseUrl)
        {
            _apiKey = apiKey;
            _baseUrl = baseUrl;
        }

        public async Task ProcessEvent(BasicDeliverEventArgs ea, string correlationId)
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

                // Simulate API call to notification service
                Console.WriteLine($"[Notification] Sending notification to attendee {checkInEvent.AttendeeId} for event {checkInEvent.EventId}. CorrelationId: {correlationId}");
                
                // In a real app, this would be an HttpClient call
                await Task.Delay(100); 
            }
            catch (Exception ex)
            {
                throw new Exception($"Notification consumer processing error: {ex.Message}", ex);
            }
        }
    }
}
