package nl.bknopper.tspeademo.ea;

public class AlgorithmOptions {

    private int mutationProbability;
    private int populationSize;
    private int nrOfGenerations;
    private int fitnessThreshold;
    private int parentSelectionSize;
    private int parentPoolSize;

    public AlgorithmOptions() {
        // empty one needed for Jackson
    }

    public AlgorithmOptions(int mutationProbability, int populationSize, int nrOfGenerations, int fitnessThresholds,
                            int parentSelectionSize, int parentPoolSize) {
        this.mutationProbability = mutationProbability;
        this.populationSize = populationSize;
        this.nrOfGenerations = nrOfGenerations;
        this.fitnessThreshold = fitnessThresholds;
        this.parentSelectionSize = parentSelectionSize;
        this.parentPoolSize = parentPoolSize;
    }

    public int getMutationProbability() {
        return mutationProbability;
    }

    public void setMutationProbability(int mutationProbability) {
        this.mutationProbability = mutationProbability;
    }

    public int getPopulationSize() {
        return populationSize;
    }

    public void setPopulationSize(int populationSize) {
        this.populationSize = populationSize;
    }

    public int getNrOfGenerations() {
        return nrOfGenerations;
    }

    public void setNrOfGenerations(int nrOfGenerations) {
        this.nrOfGenerations = nrOfGenerations;
    }

    public int getFitnessThreshold() {
        return fitnessThreshold;
    }

    public void setFitnessThreshold(int fitnessThreshold) {
        this.fitnessThreshold = fitnessThreshold;
    }

    public int getParentSelectionSize() {
        return parentSelectionSize;
    }

    public void setParentSelectionSize(int parentSelectionSize) {
        this.parentSelectionSize = parentSelectionSize;
    }

    public int getParentPoolSize() {
        return parentPoolSize;
    }

    public void setParentPoolSize(int parentPoolSize) {
        this.parentPoolSize = parentPoolSize;
    }
}
