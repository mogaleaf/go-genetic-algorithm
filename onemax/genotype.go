package onemax

import (
	"fmt"
	"go-evol/evolution/genes"
	"go-evol/evolution/genes/binary"
	"go-evol/helper"
)

type OneMaxProblem struct {
	L int
}

type phenotype struct {
	value  []uint8
	l      int
	fit    float64
	setFit bool
}

func GetPhenotype(g genes.GenotypeI) genes.PhenotypeI {
	return &phenotype{
		value: g.(*binary.Genotype).Value,
		l:     g.(*binary.Genotype).L,
	}
}

func (q *OneMaxProblem) NewRandOneMaxGenotype() genes.GenotypeI {
	newGen := binary.Genotype{
		GetPhenotypeInternal: GetPhenotype,
		Value:                initRandomGeneString(q.L),
		L:                    q.L,
		RecombinationType:    binary.ONE_POINT_CROSS_OVER,
		RecombinationRate:    0.7,
		MutationRate:         1 / float64(q.L),
		MutationType:         binary.FLIP,
	}
	return &newGen
}

func initRandomGeneString(L int) []uint8 {
	res := make([]uint8, L)
	for i, _ := range res {
		res[i] = uint8(helper.GenerateUintNumber(2))
	}
	return res
}

func (p *phenotype) Print() {
	for _, b := range p.value {
		print(fmt.Sprintf("%d", b))
	}
	println()
}
