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

func (s *BestSelection) selectPopulation(population []genes.GenotypeI, a []float64) []genes.GenotypeI {
	sort.Sort(helper.ByFitness(population))
	newPopulation := population[len(population)-s.Mu:]
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(newPopulation), func(i, j int) { newPopulation[i], newPopulation[j] = newPopulation[j], newPopulation[i] })
	return newPopulation
}


func SelectBest(population []genes.GenotypeI, s float64, selection SelectionI) []genes.GenotypeI {
	return selection.selectPopulation(population, nil)
}
