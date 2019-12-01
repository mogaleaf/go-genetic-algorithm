package binary

import (
	"go-evol/evolution"
	"go-evol/helper"
)

func copyGenotype(p1 *Genotype) *Genotype {
	return &Genotype{
		L:                    p1.L,
		Value:                p1.Value,
		MutationType:         p1.MutationType,
		GetPhenotypeInternal: p1.GetPhenotypeInternal,
		MutationRate:         p1.MutationRate,
		RecombinationRate:    p1.RecombinationRate,
		RecombinationType:    p1.RecombinationType,
	}
}

func (g *Genotype) Recombine(p2 evolution.GenotypeI) []evolution.GenotypeI {
	shouldRecombine := helper.GenerateFloatNumber() < g.RecombinationRate
	if !shouldRecombine {
		return []evolution.GenotypeI{
			copyGenotype(g),
			copyGenotype(p2.(*Genotype)),
		}
	}
	switch g.RecombinationType {
	case ONE_POINT_CROSS_OVER:
		return g.RecombineCrossFill(p2)
	}
	return nil
}

func (p1 *Genotype) RecombineCrossFill(p2 evolution.GenotypeI) []evolution.GenotypeI {
	swapK := helper.GenerateUintNumber(p1.L)
	if swapK == 0 {
		return []evolution.GenotypeI{
			copyGenotype(p1),
			copyGenotype(p2.(*Genotype)),
		}
	}
	k1 := buildNewFromParent(p1)
	k2 := buildNewFromParent(p1)
	for j := uint64(0); j < swapK; j++ {
		k1.Value[j] = p1.Value[j]
		k2.Value[j] = p2.(*Genotype).Value[j]
	}
	for j := int(swapK); j < p1.L; j++ {
		k2.Value[j] = p1.Value[j]
		k1.Value[j] = p2.(*Genotype).Value[j]
	}
	return []evolution.GenotypeI{k1, k2}
}

func buildNewFromParent(p1 *Genotype) *Genotype {
	return &Genotype{
		L:                    p1.L,
		Value:                make([]uint8, p1.L),
		MutationType:         p1.MutationType,
		GetPhenotypeInternal: p1.GetPhenotypeInternal,
		MutationRate:         p1.MutationRate,
		RecombinationRate:    p1.RecombinationRate,
		RecombinationType:    p1.RecombinationType,
	}
}
