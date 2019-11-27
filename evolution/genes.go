package evolution

type Phenotype struct {
}

type PhenotypeI interface {
	CalcFitness() int
	Good() bool
	Print()
}

type GenotypeI interface {
	GetPhenotype() PhenotypeI
	Mutate()
	Recombine(GenotypeI) []GenotypeI
	Print()
}

type CreateRandomGeneFunc func() GenotypeI
