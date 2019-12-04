package selection

import (
	"go-evol/evolution/genes"
	"go-evol/helper"
	"sort"
)

type AlgoType int

const (
	ROULETTE AlgoType = iota
	SUS
)

type ProbabilitySelection struct {
	*Selection
	ProbabilityType SelectionProbType
	AlgoType        AlgoType
	SP              float64
}

func (r *ProbabilitySelection) SelectPopulation(population []genes.GenotypeI) []genes.GenotypeI {
	a := []float64{}
	switch r.ProbabilityType {
	case FPS:
		fps := r.SelectFPS(population)
		a = append(a, fps...)
	case RANK:
		rank := r.SelectRank(population, r.SP)
		a = append(a, rank...)
	}
	switch r.AlgoType {
	case SUS:
		return r.sus(population, a)
	case ROULETTE:
		return r.roulette(population, a)
	}
	return nil
}

func (r *ProbabilitySelection) SelectFPS(population []genes.GenotypeI) []float64 {
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
	return a
}

func (r *ProbabilitySelection) SelectRank(population []genes.GenotypeI, s float64) []float64 {
	parentsPr := make([]float64, len(population))
	a := make([]float64, len(population))
	sort.Sort(helper.ByFitness(population))
	accumulateProb := 0.0
	for i := 0; i < len(population); i++ {
		parentsPr[i] = (2-s)/float64(r.getMu()) + 2*(float64(i))*(s-1)/(float64(r.getMu())*(float64(r.getMu())-1))
		accumulateProb += parentsPr[i]
		a[i] = accumulateProb
	}
	return a
}

func (r *ProbabilitySelection) roulette(population []genes.GenotypeI, a []float64) []genes.GenotypeI {
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

func (s *ProbabilitySelection) sus(population []genes.GenotypeI, a []float64) []genes.GenotypeI {
	currentMember := 0
	i := 0
	matingPool := make([]genes.GenotypeI, s.Mu)
	r := helper.GenerateFloatNumberRange(1 / float64(s.Mu))
	for currentMember <= s.Mu && i < len(a) {
		for r <= a[i] && i < len(a) {
			matingPool[currentMember] = population[i]
			r = r + (1 / float64(s.Mu))
			currentMember++
		}
		i++
	}
	return matingPool
}

func (s *ProbabilitySelection) SelectOffSpring(population []genes.GenotypeI, children []genes.GenotypeI) []genes.GenotypeI {
	return s.SelectPopulation(append(population, children...))
}
