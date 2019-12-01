package binary

import "go-evol/helper"

func (g *Genotype) Mutate() {
	shouldMutate := helper.GenerateFloatNumber() < g.MutationRate
	if !shouldMutate {
		return
	}
	g.phenotype = nil
	switch g.MutationType {
	case FLIP:
		g.mutateFlip()
	}

}

func (g *Genotype) mutateFlip() {
	for i, _ := range g.Value {
		shouldMutate := helper.GenerateFloatNumber() < (1 / float64(g.L))
		if !shouldMutate {
			continue
		}
		if g.Value[i] == 1 {
			g.Value[i] = 0
		} else {
			g.Value[i] = 1
		}
	}
	g.Print()
}
