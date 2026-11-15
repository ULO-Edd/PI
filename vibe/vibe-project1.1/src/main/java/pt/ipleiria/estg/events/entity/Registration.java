package pt.ipleiria.estg.events.entity;

import jakarta.persistence.*;
import java.time.OffsetDateTime;
import java.util.Objects;

@Entity
@Table(name = "registrations", uniqueConstraints = {
    @UniqueConstraint(columnNames = {"event_id", "user_id"})
})
public class Registration {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @ManyToOne(optional = false)
    @JoinColumn(name = "event_id")
    private Event event;

    @ManyToOne(optional = false)
    @JoinColumn(name = "user_id")
    private User attendee;

    @Column(nullable = false)
    private OffsetDateTime registrationTime;

    public Registration() {}

    public Registration(Event event, User attendee, OffsetDateTime registrationTime) {
        this.event = event;
        this.attendee = attendee;
        this.registrationTime = registrationTime;
    }

    public Long getId() { return id; }
    public void setId(Long id) { this.id = id; }

    public Event getEvent() { return event; }
    public void setEvent(Event event) { this.event = event; }

    public User getAttendee() { return attendee; }
    public void setAttendee(User attendee) { this.attendee = attendee; }

    public OffsetDateTime getRegistrationTime() { return registrationTime; }
    public void setRegistrationTime(OffsetDateTime registrationTime) { this.registrationTime = registrationTime; }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Registration that = (Registration) o;
        return Objects.equals(id, that.id);
    }

    @Override
    public int hashCode() {
        return Objects.hash(id);
    }
}
