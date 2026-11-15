using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Microsoft.EntityFrameworkCore;
using RabbitMQ.Client;
using SDD.Services.DeadLetterService.Data;
using SDD.Shared;

namespace SDD.Services.DeadLetterService.Services
{
    public interface IFailureManager
    {
        Task SaveFailure(DeadLetterEntry entry);
        Task<(List<DeadLetterEntry> Items, int Total)> GetFailures(string? queueName, int page, int pageSize);
        Task<DeadLetterEntry?> GetFailureById(int id);
        Task<bool> MarkAsReplayed(int id);
        Task<bool> MarkAsPurged(int id);
    }

    public class FailureManager : IFailureManager
    {
        private readonly DeadLetterDbContext _context;
        private readonly IConnection _rabbitConnection;

        public FailureManager(DeadLetterDbContext context, IConnection rabbitConnection)
        {
            _context = context;
            _rabbitConnection = rabbitConnection;
        }

        public async Task SaveFailure(DeadLetterEntry entry)
        {
            _context.DeadLetterEntries.Add(entry);
            await _context.SaveChangesAsync();
        }

        public async Task<(List<DeadLetterEntry> Items, int Total)> GetFailures(string? queueName, int page, int pageSize)
        {
            IQueryable<DeadLetterEntry> query = _context.DeadLetterEntries;
            if (!string.IsNullOrEmpty(queueName))
            {
                query = query.Where(e => e.OriginalQueueName == queueName);
            }

            int total = await query.CountAsync();
            var items = await query
                .OrderByDescending(e => e.FailureTimestamp)
                .Skip((page - 1) * pageSize)
                .Take(pageSize)
                .ToListAsync();

            return (items, total);
        }

        public async Task<DeadLetterEntry?> GetFailureById(int id)
        {
            return await _context.DeadLetterEntries.FindAsync(id);
        }

        public async Task<bool> MarkAsReplayed(int id)
        {
            var entry = await _context.DeadLetterEntries.FindAsync(id);
            if (entry == null) return false;

            try
            {
                using var channel = _rabbitConnection.CreateModel();
                
                // Use the shared CorrelationContext to ensure the ID is injected into headers
                var properties = channel.CreateBasicProperties();
                CorrelationContext.InjectCorrelationId(properties, entry.CorrelationId);
                properties.Persistent = true;

                var body = Encoding.UTF8.GetBytes(entry.Payload);
                
                channel.BasicPublish(
                    exchange: "", // Default exchange for direct routing to queue
                    routingKey: entry.OriginalQueueName,
                    basicProperties: properties,
                    body: body
                );

                entry.Status = DeadLetterStatus.Replayed;
                await _context.SaveChangesAsync();
                return true;
            }
            catch (Exception)
            {
                return false;
            }
        }

        public async Task<bool> MarkAsPurged(int id)
        {
            var entry = await _context.DeadLetterEntries.FindAsync(id);
            if (entry == null) return false;

            entry.Status = DeadLetterStatus.Purged;
            await _context.SaveChangesAsync();
            return true;
        }
    }
}
