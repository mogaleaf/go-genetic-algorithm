package evolution

type byFitness []GenotypeI

func (a byFitness) Len() int      { return len(a) }
func (a byFitness) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byFitness) Less(i, j int) bool {
	return a[i].GetPhenotype().CalcFitness() < a[j].GetPhenotype().CalcFitness()
}
