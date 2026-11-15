using System.ComponentModel.DataAnnotations;

namespace EventsService.Models;

public class Event
{
    [Key]
    public Guid Id { get; set; } = Guid.NewGuid();

    [Required]
    [MaxLength(100)]
    public string Title { get; set; } = string.Empty;

    [MaxLength(2000)]
    public string Description { get; set; } = string.Empty;

    [Required]
    public DateTime DateTime { get; set; }

    [Required]
    [MaxLength(255)]
    public string Location { get; set; } = string.Empty;

    [Required]
    [Range(1, int.MaxValue)]
    public int Capacity { get; set; }

    [Required]
    [Range(0, int.MaxValue)]
    public int AvailableTickets { get; set; }

    [Required]
    public Guid OrganiserId { get; set; }
}
