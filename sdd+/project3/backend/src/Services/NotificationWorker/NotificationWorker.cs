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

namespace SDD.Services.NotificationWorker
{
    public class NotificationWorker : BackgroundService
    {
        private readonly ILogger<NotificationWorker> _logger;
        private readonly IConfiguration _configuration;
        private readonly IRabbitMQConnection _rabbitConnection;
        private readonly NotificationConsumer _consumer;
        private AsyncRetryPolicy _retryPolicy;

        public NotificationWorker(ILogger<NotificationWorker> logger, IConfiguration configuration, IRabbitMQConnection rabbitConnection)
        {
            _logger = logger;
            _configuration = configuration;
            _rabbitConnection = rabbitConnection;
            
            _consumer = new NotificationConsumer(
                _configuration["NotificationApi:ApiKey"] ?? "", 
                _configuration["NotificationApi:BaseUrl"] ?? ""
            );

            // T020: Implement Polly retry policy (3 attempts, exponential backoff)
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
            _logger.LogInformation("Notification Worker started");

            try
            {
                var connection = _rabbitConnection.GetConnection();
                using var channel = connection.CreateModel();

                var queueName = _configuration["RabbitMQ:QueueName"] ?? "notification-queue";
                
                // T018: Low prefetch count for rate-limited API
                channel.BasicQos(0, 1, false);

                var consumer = new EventingBasicConsumer(channel);
                consumer.Received += async (model, ea) =>
                {
                    var correlationId = CorrelationContext.GetCorrelationId(ea.BasicProperties);
                    _logger.LogInformation($"Processing notification message {ea.BasicProperties.MessageId} with CorrelationId {correlationId}");

                    try
                    {
                        // T019 & T020: Process event with Polly retry
                        await _retryPolicy.ExecuteAsync(async () => 
                        {
                            await _consumer.ProcessEvent(ea, correlationId.ToString());
                        });

                        channel.BasicAck(ea.DeliveryTag, false);
                    }
                    catch (Exception ex)
                    {
                        _logger.LogError($"Permanent failure processing notification {ea.BasicProperties.MessageId}: {ex.Message}");
                        
                        // T021: DLX routing on final failure
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
                _logger.LogCritical($"Fatal error in NotificationWorker: {ex.Message}");
                throw;
            }
        }
    }
}
