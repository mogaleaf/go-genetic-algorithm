package selection

import (
	"go-evol/evolution/genes"
	"go-evol/helper"
	"sort"
)

type SelectionProbType int

const (
	FPS SelectionProbType = iota
	RANK
	TOURNAMENT
	BEST
	REPLACE
)

type SelectionI interface {
	selectPopulation(population []genes.GenotypeI, a []float64) []genes.GenotypeI
	getMu() int
}

type Selection struct {
	Mu int
}


func (r *Selection) getMu() int {
	return r.Mu
}


func SelectFPS(population []genes.GenotypeI, selection SelectionI) []genes.GenotypeI {
	parentsPr := make([]float64, len(population))
	a := make([]float64, len(population))
	sum := 0.0
	for i := 0; i < len(population); i++ {
		sum += population[i].GetPhenotype().CalcFitness()
	}
	accumulateProb := 0.0
	for i := 0; i < len(population); i++ {
		parentsPr[i] = population[i].GetPhenotype().CalcFitness() / sum
		accumulateProb += parentsPr[i]
		a[i] = accumulateProb
	}
	return selection.selectPopulation(population, a)
}

func SelectRank(population []genes.GenotypeI, s float64, selection SelectionI) []genes.GenotypeI {
	parentsPr := make([]float64, len(population))
	a := make([]float64, len(population))
	sort.Sort(helper.ByFitness(population))
	accumulateProb := 0.0
	for i := 0; i < len(population); i++ {
		parentsPr[i] = (2-s)/float64(selection.getMu()) + 2*(float64(i))*(s-1)/(float64(selection.getMu())*(float64(selection.getMu())-1))
		accumulateProb += parentsPr[i]
		a[i] = accumulateProb
	}
	return selection.selectPopulation(population, a)
}
