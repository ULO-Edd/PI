using EventsService.Data;
using Microsoft.EntityFrameworkCore;

var builder = WebApplication.CreateBuilder(args);

// Add DbContext with SQLite datastore
builder.Services.AddDbContext<EventsDbContext>(options =>
    options.UseSqlite(builder.Configuration.GetConnectionString("DefaultConnection")));

builder.Services.AddHealthChecks();
builder.Services.AddControllers();

var app = builder.Build();

// Automatically execute pending migrations on startup (FR-005)
using (var scope = app.Services.CreateScope())
{
    var context = scope.ServiceProvider.GetRequiredService<EventsDbContext>();
    try
    {
        context.Database.Migrate();
        Console.WriteLine("Events database migrated successfully.");
    }
    catch (Exception ex)
    {
        Console.WriteLine($"Error occurred during database migration: {ex.Message}");
    }
}

app.MapHealthChecks("/health");
app.MapControllers();

app.Run();
