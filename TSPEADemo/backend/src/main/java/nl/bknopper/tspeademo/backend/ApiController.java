package nl.bknopper.tspeademo.backend;

import nl.bknopper.tspeademo.domain.City;
import nl.bknopper.tspeademo.ea.AlgorithmOptions;
import nl.bknopper.tspeademo.ea.AlgorithmRunner;
import nl.bknopper.tspeademo.ea.CandidateSolution;
import nl.bknopper.tspeademo.util.TSPUtils;
import org.springframework.web.bind.annotation.*;

import java.text.SimpleDateFormat;
import java.util.List;
import java.util.UUID;

@RestController
@RequestMapping("/api")
public class ApiController {
    private static final SimpleDateFormat FORMAT = new SimpleDateFormat("yyyy-MM-dd-HHmm");

    @RequestMapping(value = "message", method = RequestMethod.GET)
    public Message message() {
        Message message = new Message();
        message.message = "Hello world";

        return message;
    }

    @RequestMapping(value = "currentBest", method = RequestMethod.GET)
    public CandidateSolution currentBest() {
        return AlgorithmRunner.getCurrentBest(false);
    }

    @RequestMapping(value = "latestBest", method = RequestMethod.GET)
    public CandidateSolution getLatestBest() {
        return AlgorithmRunner.getCurrentBest(true);
    }

    @RequestMapping(value = "getCities", method = RequestMethod.GET)
    public List<City> getCities() {
        return TSPUtils.getRandomizedCities();
    }

    @RequestMapping(value = "stillRunning", method = RequestMethod.GET)
    public Boolean stillRunning() {
        return AlgorithmRunner.isStillRunning();
    }

    @RequestMapping(value = "startAlgorithm", method = RequestMethod.POST)
    public void startAlgorithm(@RequestBody AlgorithmOptions options) {
        AlgorithmRunner.startAlgorithm(options);
    }

    @RequestMapping(value = "stopAlgorithm", method = RequestMethod.POST)
    public void stopAlgorithm() {
        AlgorithmRunner.stopAlgorithm();
    }


    public static class Session {
        public String token;

        public static Session newSession() {
            Session s = new Session();
            s.token = UUID.randomUUID().toString();

            return s;
        }
    }

    public static class Message {
        public String message;
    }
}