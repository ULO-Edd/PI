package pt.ipleiria.estg.events.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import pt.ipleiria.estg.events.entity.Event;
import pt.ipleiria.estg.events.entity.Registration;
import pt.ipleiria.estg.events.entity.User;
import java.util.Optional;

public interface RegistrationRepository extends JpaRepository<Registration, Long> {
    long countByEvent(Event event);
    Optional<Registration> findByEventAndAttendee(Event event, User attendee);
    boolean existsByEventAndAttendee(Event event, User attendee);
}
