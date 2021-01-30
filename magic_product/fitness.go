package magic_product

func (p *phenotype) CalcFitness() float64 {
	if p.knownFitness {
		return p.fitness
	}
	count := 0
	for i := 0; i < 3; i++ {
		product := p.board[i*3] * p.board[i*3+1] * p.board[i*3+2]

		if product == 1 {
			count++
		}
	}
	for i := 0; i < 3; i++ {
		product := p.board[i] * p.board[i+3] * p.board[i+6]
		if product == 1 {
			count++
		}
	}
	productdiag1 := p.board[0] * p.board[4] * p.board[8]
	if productdiag1 == 1 {
		count++
	}
	productdiag2 := p.board[2] * p.board[4] * p.board[6]
	if productdiag2 == 1 {
		count++
	}
	p.knownFitness = true
	p.fitness = float64(count)
	return p.fitness
}

//No diagonal checking
func (p *phenotype) Good() bool {
	return p.CalcFitness() == 8
}
