package recorder

import "go-evol/evolution/genes"

type Recoder interface {
	SelectedParents(gs []genes.GenotypeI, iter int)
	CreatedOffspring(gs []genes.GenotypeI, iter int)
	MutatedOffspring(gs []genes.GenotypeI, iter int)
	NextGeneration(gs []genes.GenotypeI, iter int)
}
