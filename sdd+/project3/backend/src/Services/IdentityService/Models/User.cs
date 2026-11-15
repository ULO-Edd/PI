using System.ComponentModel.DataAnnotations;

namespace IdentityService.Models;

public class User
{
    [Key]
    public Guid Id { get; set; } = Guid.NewGuid();

    [Required]
    [EmailAddress]
    [MaxLength(255)]
    public string Email { get; set; } = string.Empty;

    [Required]
    [MaxLength(500)]
    public string PasswordHash { get; set; } = string.Empty;

    [Required]
    [MaxLength(50)]
    public string Role { get; set; } = "Attendee"; // Default: Attendee

    [Required]
    public DateTime CreatedAt { get; set; } = DateTime.UtcNow;
}
