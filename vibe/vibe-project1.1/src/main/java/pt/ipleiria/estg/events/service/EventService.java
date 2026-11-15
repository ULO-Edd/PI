package pt.ipleiria.estg.events.service;

import org.springframework.security.access.AccessDeniedException;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import pt.ipleiria.estg.events.entity.Event;
import pt.ipleiria.estg.events.entity.Role;
import pt.ipleiria.estg.events.entity.User;
import pt.ipleiria.estg.events.model.EventCreate;
import pt.ipleiria.estg.events.model.EventUpdate;
import pt.ipleiria.estg.events.repository.EventRepository;
import pt.ipleiria.estg.events.repository.RegistrationRepository;
import pt.ipleiria.estg.events.repository.UserRepository;

import java.util.List;
import java.util.NoSuchElementException;
import java.util.stream.Collectors;

@Service
public class EventService {

    private final EventRepository eventRepository;
    private final UserRepository userRepository;
    private final RegistrationRepository registrationRepository;

    public EventService(EventRepository eventRepository, UserRepository userRepository, RegistrationRepository registrationRepository) {
        this.eventRepository = eventRepository;
        this.userRepository = userRepository;
        this.registrationRepository = registrationRepository;
    }

    public List<pt.ipleiria.estg.events.model.Event> listEvents(pt.ipleiria.estg.events.model.EventStatus status, Long organiserId) {
        Event.EventStatus entityStatus = status != null ? Event.EventStatus.valueOf(status.name()) : null;
        return eventRepository.findByStatusAndOrganiserId(entityStatus, organiserId).stream()
            .map(this::mapToDto)
            .collect(Collectors.toList());
    }

    @Transactional
    public pt.ipleiria.estg.events.model.Event createEvent(EventCreate eventCreate) {
        User organiser = getCurrentUser();
        if (organiser.getRole() == Role.ATTENDEE) {
            throw new AccessDeniedException("Attendees cannot create events");
        }

        Event event = new Event();
        event.setTitle(eventCreate.getTitle());
        event.setDescription(eventCreate.getDescription());
        event.setLocation(eventCreate.getLocation());
        event.setStartTime(eventCreate.getStartTime());
        event.setEndTime(eventCreate.getEndTime());
        event.setCapacity(eventCreate.getCapacity());
        event.setOrganiser(organiser);
        event.setStatus(Event.EventStatus.DRAFT);

        return mapToDto(eventRepository.save(event));
    }

    public pt.ipleiria.estg.events.model.Event getEvent(Long id) {
        return eventRepository.findById(id)
            .map(this::mapToDto)
            .orElseThrow(() -> new NoSuchElementException("Event not found"));
    }

    @Transactional
    public pt.ipleiria.estg.events.model.Event updateEvent(Long id, EventUpdate eventUpdate) {
        Event event = eventRepository.findById(id)
            .orElseThrow(() -> new NoSuchElementException("Event not found"));

        User currentUser = getCurrentUser();
        checkOwnershipOrAdmin(event, currentUser);

        if (eventUpdate.getTitle() != null) event.setTitle(eventUpdate.getTitle());
        if (eventUpdate.getDescription() != null) event.setDescription(eventUpdate.getDescription());
        if (eventUpdate.getLocation() != null) event.setLocation(eventUpdate.getLocation());
        if (eventUpdate.getStartTime() != null) event.setStartTime(eventUpdate.getStartTime());
        if (eventUpdate.getEndTime() != null) event.setEndTime(eventUpdate.getEndTime());
        if (eventUpdate.getCapacity() != null) event.setCapacity(eventUpdate.getCapacity());
        if (eventUpdate.getStatus() != null) {
            event.setStatus(Event.EventStatus.valueOf(eventUpdate.getStatus().name()));
        }

        return mapToDto(eventRepository.save(event));
    }

    @Transactional
    public void deleteEvent(Long id) {
        Event event = eventRepository.findById(id)
            .orElseThrow(() -> new NoSuchElementException("Event not found"));

        User currentUser = getCurrentUser();
        checkOwnershipOrAdmin(event, currentUser);

        event.setStatus(Event.EventStatus.CANCELLED);
        eventRepository.save(event);
    }

    private User getCurrentUser() {
        String username = SecurityContextHolder.getContext().getAuthentication().getName();
        return userRepository.findByUsername(username)
            .orElseThrow(() -> new IllegalStateException("User not found in context"));
    }

    private void checkOwnershipOrAdmin(Event event, User user) {
        if (user.getRole() != Role.ADMIN && !event.getOrganiser().equals(user)) {
            throw new AccessDeniedException("Only the organiser or an admin can modify this event");
        }
    }

    private pt.ipleiria.estg.events.model.Event mapToDto(Event event) {
        pt.ipleiria.estg.events.model.Event dto = new pt.ipleiria.estg.events.model.Event();
        dto.setId(event.getId());
        dto.setTitle(event.getTitle());
        dto.setDescription(event.getDescription());
        dto.setLocation(event.getLocation());
        dto.setStartTime(event.getStartTime());
        dto.setEndTime(event.getEndTime());
        dto.setOrganiserId(event.getOrganiser().getId());
        dto.setCapacity(event.getCapacity());
        dto.setStatus(pt.ipleiria.estg.events.model.EventStatus.fromValue(event.getStatus().name()));
        dto.setCurrentRegistrations((int) registrationRepository.countByEvent(event));
        return dto;
    }
}
