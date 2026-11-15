using System;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Configuration;
using RabbitMQ.Client;
using RabbitMQ.Client.Events;
using Polly;
using Polly.Retry;
using SDD.Shared;
using SDD.Shared.Events;

namespace SDD.Services.AnalyticsWorker
{
    public class AnalyticsWorker : BackgroundService
    {
        private readonly ILogger<AnalyticsWorker> _logger;
        private readonly IConfiguration _configuration;
        private readonly IRabbitMQConnection _rabbitConnection;
        private readonly CheckInEventConsumer _consumer;
        private AsyncRetryPolicy _retryPolicy;

        public AnalyticsWorker(ILogger<AnalyticsWorker> logger, IConfiguration configuration, IRabbitMQConnection rabbitConnection)
        {
            _logger = logger;
            _configuration = configuration;
            _rabbitConnection = rabbitConnection;
            _consumer = new CheckInEventConsumer();

            // T015: Implement Polly retry policy (3 attempts, exponential backoff)
            _retryPolicy = Policy
                .Handle<Exception>()
                .WaitAndRetryAsync(3, retryAttempt => 
                    TimeSpan.FromSeconds(Math.Pow(2, retryAttempt)), 
                    (exception, timeSpan, retryCount, context) =>
                    {
                        _logger.LogWarning($"Retry {retryCount} of 3 after {timeSpan.TotalSeconds}s due to: {exception.Message}");
                    });
        }

        protected override async Task ExecuteAsync(CancellationToken stoppingToken)
        {
            _logger.LogInformation("Analytics Worker started");

            try
            {
                var connection = _rabbitConnection.GetConnection();
                using var channel = connection.CreateModel();

                var queueName = _configuration["RabbitMQ:QueueName"] ?? "analytics-queue";
                
                // T013: BasicQos prefetch count
                channel.BasicQos(0, 10, false);

                var consumer = new EventingBasicConsumer(channel);
                consumer.Received += async (model, ea) =>
                {
                    var correlationId = CorrelationContext.GetCorrelationId(ea.BasicProperties);
                    _logger.LogInformation($"Processing message {ea.BasicProperties.MessageId} with CorrelationId {correlationId}");

                    try
                    {
                        // T014 & T015: Process event with Polly retry
                        await _retryPolicy.ExecuteAsync(async () => 
                        {
                            _consumer.ProcessEvent(ea, (evt) => 
                            {
                                // Simulate analytics update
                                _logger.LogInformation($"[Analytics] Updating dashboard for attendee {evt.AttendeeId} at event {evt.EventId}");
                            });
                        });

                        channel.BasicAck(ea.DeliveryTag, false);
                    }
                    catch (Exception ex)
                    {
                        _logger.LogError($"Permanent failure processing message {ea.BasicProperties.MessageId}: {ex.Message}");
                        
                        // T016: DLX routing on final failure (Nack with requeue: false)
                        channel.BasicNack(ea.DeliveryTag, false, requeue: false);
                    }
                };

                channel.BasicConsume(queue: queueName, autoAck: false, consumer: consumer);

                while (!stoppingToken.IsCancellationRequested)
                {
                    await Task.Delay(1000, stoppingToken);
                }
            }
            catch (Exception ex)
            {
                _logger.LogCritical($"Fatal error in AnalyticsWorker: {ex.Message}");
                throw;
            }
        }
    }
}
