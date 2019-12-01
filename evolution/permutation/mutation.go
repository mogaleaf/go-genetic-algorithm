package permutation

import "go-evol/helper"

func (g *Genotype) Mutate() {
	// 80% chance to mutate
	shouldMutate := helper.GenerateFloatNumber() < g.MutationRate
	if !shouldMutate {
		return
	}
	g.phenotype = nil
	switch g.MutationType {
	case SWAP:
		g.mutateSwap()
	}

}

func (g *Genotype) mutateSwap() {
	//swap mutation
	number1 := helper.GenerateUintNumber(len(g.Permutation))
	number2 := helper.GenerateUintNumber(len(g.Permutation))
	temp := g.Permutation[number1]
	g.Permutation[number1] = g.Permutation[number2]
	g.Permutation[number2] = temp
}
