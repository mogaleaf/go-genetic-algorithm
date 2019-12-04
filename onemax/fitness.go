package onemax

func (p *phenotype) CalcFitness() float64 {
	if p.setFit {
		return p.fit
	}
	count := 0.0
	for i := 0; i < len(p.value); i++ {
		count += float64(p.value[i])
	}
	p.setFit = true
	p.fit = count
	return count
}

func (p *phenotype) Good() bool {
	return p.CalcFitness() == float64(p.l)
}
