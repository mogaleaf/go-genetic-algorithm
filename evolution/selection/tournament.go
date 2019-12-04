package selection

import (
	"go-evol/evolution/genes"
	"go-evol/helper"
	"sort"
)

type TournamentSelection struct {
	*Selection
	K int
}

func (s *TournamentSelection) SelectPopulation(population []genes.GenotypeI) []genes.GenotypeI {
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

func (s *TournamentSelection) SelectOffSpring(population []genes.GenotypeI, children []genes.GenotypeI) []genes.GenotypeI {
	return s.SelectPopulation(append(population, children...))
}
