package nl.bknopper.tspeademo.ea;

import org.springframework.stereotype.Component;

import java.util.ArrayList;
import java.util.List;

@Component
public class ParallelAlgorithmRunner implements AlgorithmRunner {

    private List<AlgorithmRunner> runners;

    public ParallelAlgorithmRunner() {
    }

    @Override
    public synchronized void startAlgorithm(AlgorithmOptions options) {
        runners = new ArrayList<>();
        for (int i = 0; i < 4; i++) {
            SingleThreadedAlgorithmRunner singleThreadedAlgorithmRunner = new SingleThreadedAlgorithmRunner();
            singleThreadedAlgorithmRunner.startAlgorithm(options);
            runners.add(singleThreadedAlgorithmRunner);
        }
    }

    @Override
    public synchronized void stopAlgorithm() {
        for (AlgorithmRunner runner : runners) {
            runner.stopAlgorithm();
        }
    }

    @Override
    public synchronized CandidateSolution getCurrentBest(boolean forceRetrieval) throws IllegalStateException {
        if(forceRetrieval || isStillRunning()) {
            CandidateSolution currentBest = null;
            for (AlgorithmRunner runner : runners) {
                CandidateSolution runnersBest = runner.getCurrentBest(forceRetrieval);
                if(currentBest == null) {
                    currentBest = runnersBest;
                } else if(runnersBest != null && runnersBest.compareTo(currentBest) > 0) {
                    currentBest = runnersBest;
                }
            }
            return currentBest;
        }
        throw new IllegalStateException("No Algorithm running at this point in time. Please start one.");
    }

    @Override
    public synchronized Boolean isStillRunning() {
        for (AlgorithmRunner runner : runners) {
            if(runner.isStillRunning()) {
                return true;
            }
        }
        return false;
    }
}
