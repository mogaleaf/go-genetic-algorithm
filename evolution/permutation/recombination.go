package permutation

import (
	"go-evol/evolution"
	"go-evol/helper"
)

func (g *Genotype) Recombine(p2 evolution.GenotypeI) []evolution.GenotypeI {
	shouldRecombine := helper.GenerateFloatNumber() < g.RecombinationRate
	if !shouldRecombine {
		return []evolution.GenotypeI{
			g,
			p2,
		}
	}
	switch g.RecombinationType {
	case CUT_CROSS_FILL:
		return g.RecombineCutAndCrossFill(p2)
	}
	return nil
}

func (p1 *Genotype) RecombineCutAndCrossFill(p2 evolution.GenotypeI) []evolution.GenotypeI {
	swapK := helper.GenerateUintNumber(len(p1.Permutation))
	if swapK == 0 {
		return []evolution.GenotypeI{
			p1,
			p2,
		}
	}
	k1 := buildNewFromParent(p1)
	k2 := buildNewFromParent(p1)
	for j := uint64(0); j < swapK; j++ {
		k1.Permutation[j] = p1.Permutation[j]
		k2.Permutation[j] = p2.(*Genotype).Permutation[j]
	}
	fillWithMissing(k1.Permutation, p2.(*Genotype).Permutation, int(swapK))
	fillWithMissing(k2.Permutation, p1.Permutation, int(swapK))
	return []evolution.GenotypeI{k1, k2}
}

func buildNewFromParent(p1 *Genotype) *Genotype {
	return &Genotype{
		Permutation:          make([]int, len(p1.Permutation)),
		PermutationSize:      p1.PermutationSize,
		MutationType:         p1.MutationType,
		GetPhenotypeInternal: p1.GetPhenotypeInternal,
		MutationRate:         p1.MutationRate,
		RecombinationRate:    p1.RecombinationRate,
		RecombinationType:    p1.RecombinationType,
	}
}

func contains(slice []int, value int) bool {
	for j := 0; j < len(slice); j++ {
		if slice[j] == value {
			return true
		}
	}
	return false
}

func fillWithMissing(toFill []int, with []int, swapK int) {
	j := 0
	i := swapK
	for i < len(with) && j < len(with) {
		if !contains(toFill, with[j]) {
			toFill[i] = with[j]
			i++
		}
		j++
	}
}
