package selection

import (
	"go-evol/evolution/genes"
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
	SelectPopulation(population []genes.GenotypeI) []genes.GenotypeI
	SelectOffSpring(population []genes.GenotypeI, children []genes.GenotypeI) []genes.GenotypeI
	getMu() int
}

type Selection struct {
	Mu int
}

func (r *Selection) getMu() int {
	return r.Mu
}
