using Microsoft.EntityFrameworkCore;
using RegistrationsService.Data;
using RegistrationsService.Services;

var builder = WebApplication.CreateBuilder(args);

// Add DbContext with SQLite datastore (T011)
builder.Services.AddDbContext<RegistrationsDbContext>(options =>
    options.UseSqlite(builder.Configuration.GetConnectionString("DefaultConnection")));

// Add HttpClient and register EventsClient for inter-service communication (T013)
builder.Services.AddHttpClient<EventsClient>();

builder.Services.AddHealthChecks();
builder.Services.AddControllers();

var app = builder.Build();

// Automatically execute pending migrations on startup (FR-006)
using (var scope = app.Services.CreateScope())
{
    var context = scope.ServiceProvider.GetRequiredService<RegistrationsDbContext>();
    try
    {
        context.Database.Migrate();
        Console.WriteLine("Registrations database migrated successfully.");
    }
    catch (Exception ex)
    {
        Console.WriteLine($"Error occurred during database migration: {ex.Message}");
    }
}

app.MapHealthChecks("/health");
app.MapControllers();

app.Run();
