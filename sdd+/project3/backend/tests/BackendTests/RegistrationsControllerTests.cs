using System;
using System.Net;
using System.Net.Http;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Moq;
using Moq.Protected;
using RegistrationsService.Controllers;
using RegistrationsService.Data;
using RegistrationsService.Services;
using Xunit;

namespace BackendTests;

public class RegistrationsControllerTests
{
    [Fact]
    public async Task BookTicket_WhenCapacityExhausted_ShouldReturnBadRequestAndNotSave()
    {
        // 1. Arrange: Setup in-memory sqlite registrations DB context
        var options = new DbContextOptionsBuilder<RegistrationsDbContext>()
            .UseInMemoryDatabase(databaseName: "RegistrationsDb_Exhausted")
            .Options;

        using var context = new RegistrationsDbContext(options);

        // Mock HttpClient to return 400 Bad Request for capacity reserve call
        var mockMessageHandler = new Mock<HttpMessageHandler>();
        mockMessageHandler.Protected()
            .Setup<Task<HttpResponseMessage>>(
                "SendAsync",
                ItExpr.IsAny<HttpRequestMessage>(),
                ItExpr.IsAny<System.Threading.CancellationToken>()
            )
            .ReturnsAsync(new HttpResponseMessage
            {
                StatusCode = HttpStatusCode.BadRequest,
                Content = new StringContent("{ \"message\": \"Capacity exhausted\" }")
            });

        var httpClient = new HttpClient(mockMessageHandler.Object);
        var mockConfiguration = new Mock<Microsoft.Extensions.Configuration.IConfiguration>();
        mockConfiguration.Setup(c => c["EventsService:BaseUrl"]).Returns("http://localhost:5002");

        var eventsClient = new EventsClient(httpClient, mockConfiguration.Object);
        var controller = new RegistrationsController(context, eventsClient);

        var eventId = Guid.NewGuid();
        var attendeeId = Guid.NewGuid();
        var input = new BookInput(eventId, attendeeId, "corr-test-123");

        // 2. Act
        var result = await controller.BookTicket(input);

        // 3. Assert
        var badRequestResult = Assert.IsType<BadRequestObjectResult>(result);
        Assert.Contains("Capacity exhausted", badRequestResult.Value?.ToString() ?? string.Empty);

        // Verify no registrations were added locally
        var registrationsCount = await context.Registrations.CountAsync();
        Assert.Equal(0, registrationsCount);
    }

    [Fact]
    public async Task BookTicket_WhenDatabaseWriteFails_ShouldTriggerCapacityRollback()
    {
        // 1. Arrange: Setup database context that throws exceptions on SaveChanges
        var options = new DbContextOptionsBuilder<RegistrationsDbContext>()
            .UseInMemoryDatabase(databaseName: "RegistrationsDb_FailWrite")
            .Options;

        // Creating a custom subclass of DbContext that triggers exceptions on save to simulate crashes
        var mockContext = new TestRegistrationsDbContext(options);

        // Mock HttpClient to return 200 OK for reserve and release calls
        var mockMessageHandler = new Mock<HttpMessageHandler>();
        mockMessageHandler.Protected()
            .Setup<Task<HttpResponseMessage>>(
                "SendAsync",
                ItExpr.IsAny<HttpRequestMessage>(),
                ItExpr.IsAny<System.Threading.CancellationToken>()
            )
            .ReturnsAsync(new HttpResponseMessage
            {
                StatusCode = HttpStatusCode.OK,
                Content = new StringContent("{ \"message\": \"Success\" }")
            });

        var httpClient = new HttpClient(mockMessageHandler.Object);
        var mockConfiguration = new Mock<Microsoft.Extensions.Configuration.IConfiguration>();
        mockConfiguration.Setup(c => c["EventsService:BaseUrl"]).Returns("http://localhost:5002");

        var eventsClient = new EventsClient(httpClient, mockConfiguration.Object);
        var controller = new RegistrationsController(mockContext, eventsClient);

        var eventId = Guid.NewGuid();
        var attendeeId = Guid.NewGuid();
        var input = new BookInput(eventId, attendeeId, "corr-test-456");

        // 2. Act
        var result = await controller.BookTicket(input);

        // 3. Assert
        var badRequestResult = Assert.IsType<BadRequestObjectResult>(result);
        Assert.Contains("Ticket booking transaction failed", badRequestResult.Value?.ToString() ?? string.Empty);

        // Assert that HTTP Release was called successfully to rollback/release capacity (T015)
        mockMessageHandler.Protected().Verify(
            "SendAsync",
            Times.Exactly(2), // 1st call: reserve, 2nd call: release/rollback
            ItExpr.Is<HttpRequestMessage>(req => req.Method == HttpMethod.Post),
            ItExpr.IsAny<System.Threading.CancellationToken>()
        );
    }
}

public class TestRegistrationsDbContext : RegistrationsDbContext
{
    public TestRegistrationsDbContext(DbContextOptions<RegistrationsDbContext> options) : base(options)
    {
    }

    public override Task<int> SaveChangesAsync(System.Threading.CancellationToken cancellationToken = default)
    {
        // Force database write exceptions to simulate database/transaction failures
        throw new DbUpdateException("Simulated SQLite write restriction constraint violation exception.");
    }
}
