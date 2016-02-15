package nl.bknopper.tspeademo.ea;

import lombok.Data;

@Data
public class AlgorithmOptions {

    private int mutationProbability;
    private int populationSize;
    private int nrOfGenerations;
    private int fitnessThreshold;
    private int parentSelectionSize;
    private int parentPoolSize;
    private String algorithmStyle;

    public AlgorithmOptions() {
        // empty one needed for Jackson
    }

    public AlgorithmOptions(int mutationProbability, int populationSize, int nrOfGenerations, int fitnessThresholds,
                            int parentSelectionSize, int parentPoolSize, String algorithmStyle) {
        this.mutationProbability = mutationProbability;
        this.populationSize = populationSize;
        this.nrOfGenerations = nrOfGenerations;
        this.fitnessThreshold = fitnessThresholds;
        this.parentSelectionSize = parentSelectionSize;
        this.parentPoolSize = parentPoolSize;
        this.algorithmStyle = algorithmStyle;
    }
}
