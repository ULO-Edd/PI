package pt.ipleiria.estg.events.service;

import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import pt.ipleiria.estg.events.entity.Event;
import pt.ipleiria.estg.events.entity.Registration;
import pt.ipleiria.estg.events.entity.User;
import pt.ipleiria.estg.events.repository.EventRepository;
import pt.ipleiria.estg.events.repository.RegistrationRepository;
import pt.ipleiria.estg.events.repository.UserRepository;

import java.time.OffsetDateTime;
import java.util.NoSuchElementException;

@Service
public class RegistrationService {

    private final EventRepository eventRepository;
    private final RegistrationRepository registrationRepository;
    private final UserRepository userRepository;

    public RegistrationService(EventRepository eventRepository, RegistrationRepository registrationRepository, UserRepository userRepository) {
        this.eventRepository = eventRepository;
        this.registrationRepository = registrationRepository;
        this.userRepository = userRepository;
    }

    @Transactional
    public void registerForEvent(Long eventId) {
        Event event = eventRepository.findById(eventId)
            .orElseThrow(() -> new NoSuchElementException("Event not found"));

        if (event.getStatus() != Event.EventStatus.PUBLISHED) {
            throw new IllegalStateException("Can only register for published events");
        }

        User attendee = getCurrentUser();

        if (registrationRepository.existsByEventAndAttendee(event, attendee)) {
            throw new IllegalArgumentException("Already registered for this event");
        }

        long currentCount = registrationRepository.countByEvent(event);
        if (currentCount >= event.getCapacity()) {
            throw new IllegalStateException("Event is full");
        }

        Registration registration = new Registration(event, attendee, OffsetDateTime.now());
        registrationRepository.save(registration);
    }

    @Transactional
    public void unregisterFromEvent(Long eventId) {
        Event event = eventRepository.findById(eventId)
            .orElseThrow(() -> new NoSuchElementException("Event not found"));

        User attendee = getCurrentUser();

        Registration registration = registrationRepository.findByEventAndAttendee(event, attendee)
            .orElseThrow(() -> new NoSuchElementException("Registration not found"));

        registrationRepository.delete(registration);
    }

    private User getCurrentUser() {
        String username = SecurityContextHolder.getContext().getAuthentication().getName();
        return userRepository.findByUsername(username)
            .orElseThrow(() -> new IllegalStateException("User not found in context"));
    }
}
