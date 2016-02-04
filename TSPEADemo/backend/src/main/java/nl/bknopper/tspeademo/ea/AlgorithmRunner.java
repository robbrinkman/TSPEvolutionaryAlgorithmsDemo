package nl.bknopper.tspeademo.ea;


import nl.bknopper.tspeademo.gui.TSPEADemo;

public class AlgorithmRunner {

    private static Algorithm algorithm;

    private AlgorithmRunner() {
    }

    public synchronized static void startAlgorithm(TSPEADemo demo, AlgorithmOptions options) {
        algorithm = new Algorithm(demo, options.getMutationProbability(),
                options.getPopulationSize(),
                options.getNrOfGenerations(),
                options.getFitnessThreshold(),
                options.getParentSelectionSize(),
                options.getParentPoolSize());
        algorithm.startAlgorithm();
    }

    public synchronized static void stopAlgorithm() {
        algorithm.stopAlgorithm();
    }

    public synchronized static CandidateSolution getCurrentBest(boolean forceRetrieval) throws IllegalStateException {
        if(forceRetrieval || isStillRunning()) {
            return algorithm.getCurrentBest();
        }
        throw new IllegalStateException("No Algorithm running at this point in time. Please start one.");
    }

    public synchronized static Boolean isStillRunning() {
        if(algorithm == null) {
            throw new IllegalStateException("No Algorithm running at this point in time. Please start one.");
        }
        return algorithm.isStillRunning();
    }
}
