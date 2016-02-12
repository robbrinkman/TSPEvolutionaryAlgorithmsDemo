package nl.bknopper.tspeademo.ea;

import org.springframework.stereotype.Component;

@Component
public class SingleThreadedAlgorithmRunner implements AlgorithmRunner {

    private static Algorithm algorithm;

    public SingleThreadedAlgorithmRunner() {
    }

    @Override
    public synchronized void startAlgorithm(AlgorithmOptions options) {
        algorithm = new Algorithm(options.getMutationProbability(),
                options.getPopulationSize(),
                options.getNrOfGenerations(),
                options.getFitnessThreshold(),
                options.getParentSelectionSize(),
                options.getParentPoolSize());
        algorithm.startAlgorithm();
    }

    @Override
    public synchronized void stopAlgorithm() {
        algorithm.stopAlgorithm();
    }

    @Override
    public synchronized CandidateSolution getCurrentBest(boolean forceRetrieval) throws IllegalStateException {
        if(forceRetrieval || isStillRunning()) {
            return algorithm.getCurrentBest();
        }
        throw new IllegalStateException("No Algorithm running at this point in time. Please start one.");
    }

    @Override
    public synchronized Boolean isStillRunning() {
        if(algorithm == null) {
            throw new IllegalStateException("No Algorithm running at this point in time. Please start one.");
        }
        return algorithm.isStillRunning();
    }
}
