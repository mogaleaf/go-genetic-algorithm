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
	SurvivorSelectionConfig SelectionConfig
}

type SelectionConfig struct {
	// SelectionMethod method
	SelectionMethod SelectionI
	//Probability type (FPS,RANK)
	ProbabilityType SelectionProbType
	//For Rank selection SP
	SP float64
}

type Evolve struct {
	// Randomly create the genotype
	Create CreateRandomGeneFunc
	// Max iteration before stopping
	NumberIterationMax int
	// Size of the population
	PopulationSize int
	//SelectionMethod of the parents
	ParentsSelectionConfig SelectionConfig
	//Next generation Selection
	SurvivorSelectionConfig SelectionConfig
}

func NewEvolve(create CreateRandomGeneFunc, options ...EvolveOpts) *Evolve {
	e := &Evolve{
		Create:             create,
		NumberIterationMax: 100,
		PopulationSize:     100,
		SurvivorSelectionConfig: SelectionConfig{
			SelectionMethod: &BestSelection{
				Selection: &Selection{
					Mu: 100,
				},
			},
			ProbabilityType: BEST,
		},
		ParentsSelectionConfig: SelectionConfig{
			SelectionMethod: &SusSelection{
				Selection: &Selection{
					Mu: 100,
				},
			},
			SP:              1.5,
			ProbabilityType: RANK,
		},
	}
	for _, option := range options {
		option(e)
	}
	return e
}

type EvolveOpts func(*Evolve)

func WithPopulationSize(populationSize int) EvolveOpts {
	return func(e *Evolve) {
		e.PopulationSize = populationSize
	}
}

func WithNumberIterationMax(iteration int) EvolveOpts {
	return func(e *Evolve) {
		e.NumberIterationMax = iteration
	}
}

func WithParentsSelectionConfig(selection SelectionConfig) EvolveOpts {
	return func(e *Evolve) {
		e.ParentsSelectionConfig = selection
	}
}

func WithSurvivorSelectionConfig(selection SelectionConfig) EvolveOpts {
	return func(e *Evolve) {
		e.SurvivorSelectionConfig = selection
	}
}

func (e *Evolve) Evolve() (PhenotypeI, int) {
	//1. Initialise
	population := initPopulation(e.Create, e.PopulationSize)
	//2. evaluate
	winner := evaluate(population)
	if winner != nil {
		return winner, 0
	}
	for i := 0; i < e.NumberIterationMax; i++ {
		//1. Select parents
		parents := selectParents(e.ParentsSelectionConfig, population)

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
		population = selectNextGen(e.SurvivorSelectionConfig, population, children)
	}
	return nil, e.NumberIterationMax
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

func selectNextGen(config SelectionConfig, population []GenotypeI, children []GenotypeI) []GenotypeI {
	var nextGen []GenotypeI

	switch config.ProbabilityType {
	case FPS:
		nextGen = selectFPS(append(population, children...), config.SelectionMethod)
	case RANK:
		nextGen = selectRank(append(population, children...), config.SP, config.SelectionMethod)
	case BEST:
		nextGen = selectBest(append(population, children...), config.SP, config.SelectionMethod)
	case REPLACE:
		nextGen = children
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
