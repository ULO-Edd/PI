package pt.ipleiria.estg.events;

import org.springframework.boot.CommandLineRunner;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Component;
import pt.ipleiria.estg.events.entity.Role;
import pt.ipleiria.estg.events.entity.User;
import pt.ipleiria.estg.events.repository.UserRepository;

@Component
public class DataInitializer implements CommandLineRunner {

    private final UserRepository userRepository;
    private final PasswordEncoder passwordEncoder;

    public DataInitializer(UserRepository userRepository, PasswordEncoder passwordEncoder) {
        this.userRepository = userRepository;
        this.passwordEncoder = passwordEncoder;
    }

    @Override
    public void run(String... args) {
        if (userRepository.count() == 0) {
            userRepository.save(new User("admin", passwordEncoder.encode("admin"), Role.ADMIN));
            userRepository.save(new User("organiser1", passwordEncoder.encode("password"), Role.ORGANISER));
            userRepository.save(new User("organiser2", passwordEncoder.encode("password"), Role.ORGANISER));
            userRepository.save(new User("attendee1", passwordEncoder.encode("password"), Role.ATTENDEE));
            userRepository.save(new User("attendee2", passwordEncoder.encode("password"), Role.ATTENDEE));
            System.out.println("Sample users initialized.");
        }
    }
}
