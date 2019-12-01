package evolution

type PhenotypeI interface {
	CalcFitness() float64
	Good() bool
	Print()
}

type GenotypeI interface {
	GetPhenotype() PhenotypeI
	Mutate()
	Recombine(GenotypeI) []GenotypeI
}

//Create a New random Genotype to init the algo
type CreateRandomGeneFunc func() GenotypeI
