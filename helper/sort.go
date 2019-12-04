package helper

import (
	"go-evol/evolution/genes"
)

type ByFitness []genes.GenotypeI

func (a ByFitness) Len() int      { return len(a) }
func (a ByFitness) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByFitness) Less(i, j int) bool {
	return a[i].GetPhenotype().CalcFitness() < a[j].GetPhenotype().CalcFitness()
}
