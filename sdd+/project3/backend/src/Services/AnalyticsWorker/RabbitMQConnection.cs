using System;
using System.Text;
using RabbitMQ.Client;

namespace SDD.Services.AnalyticsWorker
{
    public interface IRabbitMQConnection
    {
        IConnection GetConnection();
    }

    public class RabbitMQConnection : IRabbitMQConnection
    {
        private readonly string _hostName;
        private readonly int _port;
        private readonly string _userName;
        private readonly string _password;
        private IConnection _connection;

        public RabbitMQConnection(string hostName, int port, string userName, string password)
        {
            _hostName = hostName;
            _port = port;
            _userName = userName;
            _password = password;
        }

        public IConnection GetConnection()
        {
            if (_connection == null || !_connection.IsOpen)
            {
                var factory = new ConnectionFactory()
                {
                    HostName = _hostName,
                    Port = _port,
                    UserName = _userName,
                    Password = _password
                };
                _connection = factory.CreateConnection();
            }
            return _connection;
        }
    }
}
