package magic_product

import (
	"fmt"
	"go-evol/evolution/genes"
	"go-evol/evolution/genes/permutation"
	"math/rand"
)

type MagicNumberBoard struct {
}

//Internal queen phenotype
type phenotype struct {
	board        []float64
	knownFitness bool
	fitness      float64
}

func GetPhenotype(g genes.GenotypeI) genes.PhenotypeI {
	return &phenotype{
		board: g.(*permutation.Genotype).Permutation,
	}
}

func (q *MagicNumberBoard) NewRandMagicProductGenotype() genes.GenotypeI {
	newGen := permutation.Genotype{
		GetPhenotypeInternal: GetPhenotype,
		Permutation:          Init(),
		PermutationSize:      3,
		RecombinationType:    permutation.CUT_CROSS_FILL,
		RecombinationRate:    0.9,
		MutationRate:         0.8,
		MutationType:         permutation.SWAP,
	}
	return &newGen
}

func (p *phenotype) Print() {
	for i := 0; i < 3; i++ {
		fmt.Print(fmt.Sprintf("%0.2f  %0.2f %0.2f", p.board[i*3], p.board[i*3+1], p.board[i*3+2]))
		println()
	}
}

func Init() []float64 {
	m := []float64{1.0, 2.0, 3.0, 6.0, 1.0 / 2.0, 1.0 / 3.0, 1.0 / 6.0, 2.0 / 3.0, 3.0 / 2.0}

	for i := 0; i < 9; i++ {
		j := rand.Intn(i + 1)
		tmp := m[i]
		m[i] = m[j]
		m[j] = tmp
	}
	return m
}
