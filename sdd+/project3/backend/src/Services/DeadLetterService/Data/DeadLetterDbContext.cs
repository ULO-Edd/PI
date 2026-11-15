using Microsoft.EntityFrameworkCore;

namespace SDD.Services.DeadLetterService.Data
{
    public class DeadLetterDbContext : DbContext
    {
        public DeadLetterDbContext(DbContextOptions<DeadLetterDbContext> options) : base(options) { }

        public DbSet<DeadLetterEntry> DeadLetterEntries { get; set; } = null!;

        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<DeadLetterEntry>()
                .HasIndex(e => e.CorrelationId);
        }
    }
}
