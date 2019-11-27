package queen

func (p *Phenotype) CalcFitness() int {
	count := 0
	for i := (0); i < 8; i++ {
		for j := i + 1; j < 8; j++ {
			if (j-i)*(j-i) == (p.board[j]-p.board[i])*(p.board[j]-p.board[i]) {
				count++
			}
		}
	}
	return -count
}

func (p *Phenotype) Good() bool {
	return p.CalcFitness() == 0
}
