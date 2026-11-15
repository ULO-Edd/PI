package pt.ipleiria.estg.events.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import pt.ipleiria.estg.events.entity.Event;
import pt.ipleiria.estg.events.entity.User;
import java.util.List;

public interface EventRepository extends JpaRepository<Event, Long> {
    List<Event> findByStatus(Event.EventStatus status);
    List<Event> findByOrganiser(User organiser);

    @Query("SELECT e FROM Event e WHERE (:status IS NULL OR e.status = :status) AND (:organiserId IS NULL OR e.organiser.id = :organiserId)")
    List<Event> findByStatusAndOrganiserId(Event.EventStatus status, Long organiserId);
}
