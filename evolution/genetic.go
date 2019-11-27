package evolution

import (
	"go-evol/helper"
	"math/rand"
	"sort"
	"time"
)

type byFitness []GenotypeI

func (a byFitness) Len() int      { return len(a) }
func (a byFitness) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byFitness) Less(i, j int) bool {
	return a[i].GetPhenotype().CalcFitness() < a[j].GetPhenotype().CalcFitness()
}

func Evolve(create CreateRandomGeneFunc, numberIterationMax int, populationSize int) (PhenotypeI, int) {
	//1. Initialise
	population := initPopulation(create, populationSize)
	//2. evaluate
	winner := evaluate(population)
	if winner != nil {
		return winner, 0
	}
	for i := 0; i < numberIterationMax; i++ {
		//1. Select parents
		p1, p2 := selectParent(population)
		//2.Recombine
		children := p1.Recombine(p2)
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
		population = selectNextGeneration(population, children, populationSize)
	}
	return nil, numberIterationMax
}

func initPopulation(createRandom func() GenotypeI, populationSize int) []GenotypeI {
	is := make([]GenotypeI, populationSize)
	for i := 0; i < populationSize; i++ {
		is[i] = createRandom()
	}
	return is
}

func selectParent(population []GenotypeI) (GenotypeI, GenotypeI) {
	parentsPossibility := make([]GenotypeI, 5)
	for i := 0; i < 5; i++ {
		number := helper.GenerateUintNumber(len(population))
		parentsPossibility[i] = population[number]
	}
	sort.Sort(byFitness(parentsPossibility))
	return parentsPossibility[4], parentsPossibility[3]
}

func selectNextGeneration(population []GenotypeI, children []GenotypeI, populationSize int) []GenotypeI {
	population = append(population, children...)
	sort.Sort(byFitness(population))
	newPopulation := population[len(population)-populationSize:]
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(newPopulation), func(i, j int) { newPopulation[i], newPopulation[j] = newPopulation[j], newPopulation[i] })
	return newPopulation
}

func evaluate(population []GenotypeI) PhenotypeI {
	for _, child := range population {
		phenotype := child.GetPhenotype()
		if phenotype.Good() {
			return phenotype
		}
	}
	return nil
}
