using Microsoft.EntityFrameworkCore;
using RegistrationsService.Models;

namespace RegistrationsService.Data;

public class RegistrationsDbContext : DbContext
{
    public RegistrationsDbContext(DbContextOptions<RegistrationsDbContext> options) : base(options)
    {
    }

    public DbSet<Registration> Registrations => Set<Registration>();

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        base.OnModelCreating(modelBuilder);
    }
}
