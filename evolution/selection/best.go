package selection

import (
	"go-evol/evolution/genes"
	"go-evol/helper"
	"math/rand"
	"sort"
	"time"
)

type BestSelection struct {
	*Selection
}

func (s *BestSelection) SelectPopulation(population []genes.GenotypeI) []genes.GenotypeI {
	sort.Sort(helper.ByFitness(population))
	newPopulation := population[len(population)-s.Mu:]
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(newPopulation), func(i, j int) { newPopulation[i], newPopulation[j] = newPopulation[j], newPopulation[i] })
	return newPopulation
}

func (s *BestSelection) SelectOffSpring(population []genes.GenotypeI, children []genes.GenotypeI) []genes.GenotypeI {
	return s.SelectPopulation(append(population, children...))
}
