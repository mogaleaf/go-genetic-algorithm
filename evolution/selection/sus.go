package selection

import (
	"go-evol/evolution/genes"
	"go-evol/helper"
)

type SusSelection struct {
	*Selection
}

func (s *SusSelection) selectPopulation(population []genes.GenotypeI, a []float64) []genes.GenotypeI {
	currentMember := 0
	i := 0
	matingPool := make([]genes.GenotypeI, s.Mu)
	r := helper.GenerateFloatNumberRange(1 / float64(s.Mu))
	for currentMember < s.Mu {
		for r <= a[i] {
			matingPool[currentMember] = population[i]
			r = r + (1 / float64(s.Mu))
			currentMember++
		}
		i++
	}
	return matingPool
}

