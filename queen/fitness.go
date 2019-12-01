package queen

func (p *phenotype) CalcFitness() float64 {
	if p.knownFitness {
		return p.fitness
	}
	count := 0
	for i := 0; i < len(p.board); i++ {
		for j := i + 1; j < len(p.board); j++ {
			if (j-i)*(j-i) == (p.board[j]-p.board[i])*(p.board[j]-p.board[i]) {
				count++
			}
		}
	}
	p.knownFitness = true
	p.fitness = float64(-count)
	return p.fitness
}

//No diagonal checking
func (p *phenotype) Good() bool {
	return p.CalcFitness() == 0
}
