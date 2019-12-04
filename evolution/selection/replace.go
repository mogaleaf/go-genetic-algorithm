package selection

import (
	"go-evol/evolution/genes"
)

type ReplaceSelection struct {
	*Selection
}

func (s *ReplaceSelection) SelectPopulation(population []genes.GenotypeI) []genes.GenotypeI {
	return population
}

func (s *ReplaceSelection) SelectOffSpring(population []genes.GenotypeI, children []genes.GenotypeI) []genes.GenotypeI {
	return children
}
