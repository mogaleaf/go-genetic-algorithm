package selection

import (
	"go-evol/evolution/genes"
	"go-evol/helper"
	"math/rand"
	"sort"
	"time"
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

type RouletteWheelSelection struct {
	*Selection
}

type SusSelection struct {
	*Selection
}

type BestSelection struct {
	*Selection
}

type TournamentSelection struct {
	*Selection
	K int
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

func SelectTournament(population []genes.GenotypeI, s float64, selection SelectionI) []genes.GenotypeI {
	return selection.selectPopulation(population, nil)
}

func SelectBest(population []genes.GenotypeI, s float64, selection SelectionI) []genes.GenotypeI {
	return selection.selectPopulation(population, nil)
}

func (r *Selection) getMu() int {
	return r.Mu
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

func (s *BestSelection) selectPopulation(population []genes.GenotypeI, a []float64) []genes.GenotypeI {
	sort.Sort(helper.ByFitness(population))
	newPopulation := population[len(population)-s.Mu:]
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(newPopulation), func(i, j int) { newPopulation[i], newPopulation[j] = newPopulation[j], newPopulation[i] })
	return newPopulation
}

func (s *TournamentSelection) selectPopulation(population []genes.GenotypeI, a []float64) []genes.GenotypeI {
	var newPopulation []genes.GenotypeI
	for i := 0; i < s.Mu; i++ {
		var currParent []genes.GenotypeI
		for j := 0; j < s.K; j++ {
			randK := helper.GenerateUintNumber(len(population))
			currParent = append(currParent, population[randK])
		}
		sort.Sort(helper.ByFitness(currParent))
		newPopulation = append(newPopulation, currParent[len(currParent)-1])
	}
	return newPopulation
}
