package permutation

import (
	"fmt"
	"go-evol/evolution/genes"
)

type MutationType int
type RecombinationType int

const (
	SWAP MutationType = iota

	CUT_CROSS_FILL RecombinationType = iota
)

type Genotype struct {
	Permutation          []int
	PermutationSize      int
	MutationRate         float64
	RecombinationRate    float64
	MutationType         MutationType
	RecombinationType    RecombinationType
	GetPhenotypeInternal func(genes.GenotypeI) genes.PhenotypeI
	phenotype            genes.PhenotypeI
}

type GenotypeConfig struct {
	MutationRate    float64
	PermutationSize int
	MutationType    MutationType
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
	for j := 0; j < len(g.Permutation); j++ {
		print(fmt.Sprintf("%d", g.Permutation[j]))
	}
	println()
}
