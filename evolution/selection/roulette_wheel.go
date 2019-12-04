package selection

import (
	"go-evol/evolution/genes"
	"go-evol/helper"
)

type RouletteWheelSelection struct {
	*Selection
}


func (r *RouletteWheelSelection) selectPopulation(population []genes.GenotypeI, a []float64) []genes.GenotypeI {
	matingPool := make([]genes.GenotypeI, r.Mu)
	currentMember := 0
	for currentMember < r.Mu {
		r := helper.GenerateFloatNumber()
		i := 0
		for ; a[i] < r; i++ {
		}
		matingPool[currentMember] = population[i]
		currentMember++
	}
	return matingPool
}
