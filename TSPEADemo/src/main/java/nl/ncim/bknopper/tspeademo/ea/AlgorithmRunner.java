package nl.ncim.bknopper.tspeademo.ea;


import nl.ncim.bknopper.tspeademo.gui.TSPEADemo;

public class AlgorithmRunner {

    private static Algorithm algorithm;

    private AlgorithmRunner() {
    }

    public synchronized static void startAlgorithm(TSPEADemo demo, int mutationProbability, int populationSizeSlide, int nrOfGenerationsSlide,
                                                   int fitnessThresholdsSlide, int parentSelectionSizeSlide, int parentPoolSizeSlide) {
        algorithm = new Algorithm(demo, mutationProbability,
                populationSizeSlide,
                nrOfGenerationsSlide,
                fitnessThresholdsSlide,
                parentSelectionSizeSlide,
                parentPoolSizeSlide);
        algorithm.startAlgorithm();
    }

    public synchronized static void stopAlgorithm() {
        algorithm.stopAlgorithm();
    }
}
