package onemax

func (p *phenotype) CalcFitness() float64 {
	count := 0.0
	for i := 0; i < len(p.value); i++ {
		count += float64(p.value[i])
	}
	return count
}

//No diagonal checking
func (p *phenotype) Good() bool {
	return p.CalcFitness() == float64(p.l)
}
