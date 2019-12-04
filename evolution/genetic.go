package evolution

import (
	"go-evol/evolution/genes"
	"go-evol/evolution/selection"
	"sync"
)

type EvolutionConfig struct {
	// Randomly create the genotype
	Create genes.CreateRandomGeneFunc
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
	SelectionMethod selection.SelectionI
	//Probability type (FPS,RANK)
	ProbabilityType selection.SelectionProbType
	//For Rank selection SP
	SP float64
}

type Evolve struct {
	// Randomly create the genotype
	Create genes.CreateRandomGeneFunc
	// Max iteration before stopping
	NumberIterationMax int
	// Size of the population
	PopulationSize int
	//SelectionMethod of the parents
	ParentsSelectionConfig SelectionConfig
	//Next generation Selection
	SurvivorSelectionConfig SelectionConfig
}

func NewEvolve(create genes.CreateRandomGeneFunc, options ...EvolveOpts) *Evolve {
	e := &Evolve{
		Create:             create,
		NumberIterationMax: 100,
		PopulationSize:     100,
		SurvivorSelectionConfig: SelectionConfig{
			SelectionMethod: &selection.BestSelection{
				Selection: &selection.Selection{
					Mu: 100,
				},
			},
			ProbabilityType: selection.BEST,
		},
		ParentsSelectionConfig: SelectionConfig{
			SelectionMethod: &selection.SusSelection{
				Selection: &selection.Selection{
					Mu: 100,
				},
			},
			SP:              1.5,
			ProbabilityType: selection.RANK,
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

func (e *Evolve) Evolve() (genes.PhenotypeI, int) {
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
		var children []genes.GenotypeI
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

func selectParents(config SelectionConfig, population []genes.GenotypeI) []genes.GenotypeI {
	var parents []genes.GenotypeI

	switch config.ProbabilityType {
	case selection.FPS:
		parents = selection.SelectFPS(population, config.SelectionMethod)
	case selection.RANK:
		parents = selection.SelectRank(population, config.SP, config.SelectionMethod)
	case selection.TOURNAMENT:
		parents = selection.SelectTournament(population, config.SP, config.SelectionMethod)
	}
	return parents
}

func selectNextGen(config SelectionConfig, population []genes.GenotypeI, children []genes.GenotypeI) []genes.GenotypeI {
	var nextGen []genes.GenotypeI

	switch config.ProbabilityType {
	case selection.FPS:
		nextGen = selection.SelectFPS(append(population, children...), config.SelectionMethod)
	case selection.RANK:
		nextGen = selection.SelectRank(append(population, children...), config.SP, config.SelectionMethod)
	case selection.BEST:
		nextGen = selection.SelectBest(append(population, children...), config.SP, config.SelectionMethod)
	case selection.REPLACE:
		nextGen = children
	}
	return nextGen
}

func initPopulation(createRandom func() genes.GenotypeI, populationSize int) []genes.GenotypeI {
	is := make([]genes.GenotypeI, populationSize)
	for i := 0; i < populationSize; i++ {
		is[i] = createRandom()
	}
	return is
}

func evaluate(population []genes.GenotypeI) genes.PhenotypeI {
	phenotypeResponse := make(chan genes.PhenotypeI, len(population))
	var wg sync.WaitGroup
	wg.Add(len(population))
	for _, child := range population {
		go func(child genes.GenotypeI) {
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
