package onemax

import (
	"fmt"
	"go-evol/evolution"
	"go-evol/evolution/binary"
	"go-evol/helper"
)

type OneMaxProblem struct {
	L int
}

//Internal queen phenotype
type phenotype struct {
	value []uint8
	l     int
}

func GetPhenotype(g evolution.GenotypeI) evolution.PhenotypeI {
	return &phenotype{
		value: g.(*binary.Genotype).Value,
		l:     g.(*binary.Genotype).L,
	}
}

func (q *OneMaxProblem) NewRandQueenGenotype() evolution.GenotypeI {
	newGen := binary.Genotype{
		GetPhenotypeInternal: GetPhenotype,
		Value:                initRandomGeneString(q.L),
		L:                    q.L,
		RecombinationType:    binary.ONE_POINT_CROSS_OVER,
		RecombinationRate:    0.7,
		MutationRate:         0.8,
		MutationType:         binary.FLIP,
	}
	return &newGen
}

func initRandomGeneString(L int) []uint8 {
	res := make([]uint8, L)
	for i, _ := range res {
		res[i] = uint8(helper.GenerateUintNumber(1))
	}
	return res
}

func (p *phenotype) Print() {
	for _, b := range p.value {
		print(fmt.Sprintf("%d", b))
	}
	println()
}
