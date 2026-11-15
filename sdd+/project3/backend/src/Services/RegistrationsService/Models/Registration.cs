using System.ComponentModel.DataAnnotations;

namespace RegistrationsService.Models;

public class Registration
{
    [Key]
    public Guid Id { get; set; } = Guid.NewGuid();

    [Required]
    public Guid EventId { get; set; }

    [Required]
    public Guid AttendeeId { get; set; }

    [Required]
    [MaxLength(50)]
    public string Status { get; set; } = "Pending"; // Pending, Confirmed, Cancelled

    [Required]
    public DateTime BookedAt { get; set; } = DateTime.UtcNow;

    [Required]
    [MaxLength(100)]
    public string CorrelationId { get; set; } = string.Empty;
}
