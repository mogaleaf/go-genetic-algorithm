package evolution

import (
	"sync"
)

type EvolutionConfig struct {
	// Randomly create the genotype
	Create CreateRandomGeneFunc
	// Max iteration before stopping
	NumberIterationMax int
	// Size of the population
	PopulationSize int
	//SelectionMethod of the parents
	ParentsSelectionConfig SelectionConfig
	//Next generation Selection
	SurvivorSelectionConfig GenerationSelectionConfig
}

type SelectionConfig struct {
	// SelectionMethod method
	SelectionMethod SelectionI
	//Probability type (FPS,RANK)
	ProbabilityType SelectionProbType
	//For Rank selection SP
	SP float64
}

type GenerationSelectionConfig struct {
	*SelectionConfig
}

func Evolve(config EvolutionConfig) (PhenotypeI, int) {
	//1. Initialise
	population := initPopulation(config.Create, config.PopulationSize)
	//2. evaluate
	winner := evaluate(population)
	if winner != nil {
		return winner, 0
	}
	for i := 0; i < config.NumberIterationMax; i++ {
		//1. Select parents
		parents := selectParents(config.ParentsSelectionConfig, population)

		//2.Recombine TODO
		var children []GenotypeI
		for i := 0; i < len(parents)-1; i = i + 2 {
			children = append(children, parents[i].Recombine(parents[i+1])...)
		}
		//3. Mutate
		for _, child := range children {
			child.Mutate()
		}
		//4. evaluate
		winner := evaluate(children)
		if winner != nil {
			return winner, i
		}
		//5. Select next gen
		population = selectNextGen(config.SurvivorSelectionConfig, append(population, children...))
	}
	return nil, config.NumberIterationMax
}

func selectParents(config SelectionConfig, population []GenotypeI) []GenotypeI {
	var parents []GenotypeI

	switch config.ProbabilityType {
	case FPS:
		parents = selectFPS(population, config.SelectionMethod)
	case RANK:
		parents = selectRank(population, config.SP, config.SelectionMethod)
	}
	return parents
}

func selectNextGen(config GenerationSelectionConfig, population []GenotypeI) []GenotypeI {
	var nextGen []GenotypeI

	switch config.ProbabilityType {
	case FPS:
		nextGen = selectFPS(population, config.SelectionMethod)
	case RANK:
		nextGen = selectRank(population, config.SP, config.SelectionMethod)
	case BEST:
		nextGen = selectBest(population, config.SP, config.SelectionMethod)
	}
	return nextGen
}

func initPopulation(createRandom func() GenotypeI, populationSize int) []GenotypeI {
	is := make([]GenotypeI, populationSize)
	for i := 0; i < populationSize; i++ {
		is[i] = createRandom()
	}
	return is
}

func evaluate(population []GenotypeI) PhenotypeI {
	phenotypeResponse := make(chan PhenotypeI, len(population))
	var wg sync.WaitGroup
	wg.Add(len(population))
	for _, child := range population {
		go func(child GenotypeI) {
			defer wg.Done()
			phenotype := child.GetPhenotype()
			if phenotype.Good() {
				phenotypeResponse <- phenotype
			}
			return
		}(child)
	}
	wg.Wait()
	close(phenotypeResponse)
	for msg := range phenotypeResponse {
		return msg
	}
	return nil
}
