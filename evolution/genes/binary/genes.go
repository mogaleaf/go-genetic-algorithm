package binary

import (
	"fmt"
	"go-evol/evolution/genes"
)

type MutationType int
type RecombinationType int

const (
	FLIP MutationType = iota

	ONE_POINT_CROSS_OVER RecombinationType = iota
)

type Genotype struct {
	Value                []uint8
	L                    int
	MutationRate         float64
	RecombinationRate    float64
	MutationType         MutationType
	RecombinationType    RecombinationType
	GetPhenotypeInternal func(genes.GenotypeI) genes.PhenotypeI
	phenotype            genes.PhenotypeI
}

func (g *Genotype) GetPhenotype() genes.PhenotypeI {
	if g.phenotype != nil {
		return g.phenotype
	}
	internal := g.GetPhenotypeInternal(g)
	g.phenotype = internal
	return internal
}

func (g *Genotype) Print() {
	for _, b := range g.Value {
		print(fmt.Sprintf("%d", b))
	}
	println()
}
