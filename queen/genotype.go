package queen

import (
	"go-evol/evolution"
	"go-evol/evolution/permutation"
	"math/rand"
)

type QueensChessBoard struct {
	SizeChessBoard int
}

//Internal queen phenotype
type phenotype struct {
	board        []int
	knownFitness bool
	fitness      float64
}

func GetPhenotype(g evolution.GenotypeI) evolution.PhenotypeI {
	return &phenotype{
		board: g.(*permutation.Genotype).Permutation,
	}
}

func (q *QueensChessBoard) NewRandQueenGenotype() evolution.GenotypeI {
	newGen := permutation.Genotype{
		GetPhenotypeInternal: GetPhenotype,
		Permutation:          rand.Perm(q.SizeChessBoard),
		PermutationSize:      q.SizeChessBoard,
		RecombinationType:    permutation.CUT_CROSS_FILL,
		RecombinationRate:    0.9,
		MutationRate:         0.8,
		MutationType:         permutation.SWAP,
	}
	return &newGen
}

func (p *phenotype) Print() {
	for i := 0; i < len(p.board); i++ {
		for j := 0; j < len(p.board); j++ {
			if j == p.board[i] {
				print(" 1 ")
			} else {
				print(" 0 ")
			}
		}
		println()
	}
}
