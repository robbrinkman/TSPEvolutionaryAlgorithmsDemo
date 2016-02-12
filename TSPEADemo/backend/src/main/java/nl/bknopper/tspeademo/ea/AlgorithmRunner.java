package nl.bknopper.tspeademo.ea;

/**
 * Created by bknopper on 12/02/16.
 */
public interface AlgorithmRunner {
    void startAlgorithm(AlgorithmOptions options);

    void stopAlgorithm();

    CandidateSolution getCurrentBest(boolean forceRetrieval) throws IllegalStateException;

    Boolean isStillRunning();
}
