package pt.ipleiria.estg.events.api;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestMapping;
import pt.ipleiria.estg.events.model.Event;
import pt.ipleiria.estg.events.model.EventCreate;
import pt.ipleiria.estg.events.model.EventStatus;
import pt.ipleiria.estg.events.model.EventUpdate;
import pt.ipleiria.estg.events.service.EventService;

import java.util.List;

@RestController
@RequestMapping("/api/v1")
public class EventsController implements EventsApi {

    private final EventService eventService;

    public EventsController(EventService eventService) {
        this.eventService = eventService;
    }

    @Override
    public ResponseEntity<Event> createEvent(EventCreate eventCreate) {
        return new ResponseEntity<>(eventService.createEvent(eventCreate), HttpStatus.CREATED);
    }

    @Override
    public ResponseEntity<Void> deleteEvent(Long id) {
        eventService.deleteEvent(id);
        return ResponseEntity.noContent().build();
    }

    @Override
    public ResponseEntity<Event> getEvent(Long id) {
        return ResponseEntity.ok(eventService.getEvent(id));
    }

    @Override
    public ResponseEntity<List<Event>> listEvents(EventStatus status, Long organiserId) {
        return ResponseEntity.ok(eventService.listEvents(status, organiserId));
    }

    @Override
    public ResponseEntity<Event> updateEvent(Long id, EventUpdate eventUpdate) {
        return ResponseEntity.ok(eventService.updateEvent(id, eventUpdate));
    }
}
