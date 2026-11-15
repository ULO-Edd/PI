package pt.ipleiria.estg.events.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import pt.ipleiria.estg.events.entity.User;
import java.util.Optional;

public interface UserRepository extends JpaRepository<User, Long> {
    Optional<User> findByUsername(String username);
}
