using System;
using RabbitMQ.Client;
using System.Text;

namespace SDD.Shared
{
    public static class CorrelationContext
    {
        private const string CorrelationIdHeader = "X-Correlation-ID";

        /// <summary>
        /// Extracts the CorrelationId from AMQP message headers.
        /// </summary>
        public static Guid GetCorrelationId(IBasicProperties properties)
        {
            if (properties?.Headers != null && properties.Headers.ContainsKey(CorrelationIdHeader))
            {
                var bytes = (byte[])properties.Headers[CorrelationIdHeader];
                var idString = Encoding.UTF8.GetString(bytes);
                if (Guid.TryParse(idString, out var guid))
                {
                    return guid;
                }
            }
            return Guid.Empty;
        }

        /// <summary>
        /// Injects the CorrelationId into AMQP message headers.
        /// </summary>
        public static void InjectCorrelationId(IBasicProperties properties, Guid correlationId)
        {
            if (properties.Headers == null)
            {
                properties.Headers = new System.Collections.Generic.Dictionary<string, object>();
            }
            properties.Headers[CorrelationIdHeader] = Encoding.UTF8.GetBytes(correlationId.ToString());
        }
    }
}
