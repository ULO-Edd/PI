package pt.ipleiria.estg.events.api;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestMapping;
import pt.ipleiria.estg.events.service.RegistrationService;

@RestController
@RequestMapping("/api/v1")
public class RegistrationsController implements RegistrationsApi {

    private final RegistrationService registrationService;

    public RegistrationsController(RegistrationService registrationService) {
        this.registrationService = registrationService;
    }

    @Override
    public ResponseEntity<Void> registerForEvent(Long id) {
        registrationService.registerForEvent(id);
        return new ResponseEntity<>(HttpStatus.CREATED);
    }

    @Override
    public ResponseEntity<Void> unregisterFromEvent(Long id) {
        registrationService.unregisterFromEvent(id);
        return ResponseEntity.noContent().build();
    }
}
