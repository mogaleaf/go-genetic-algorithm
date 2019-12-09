package triangular

func (p *phenotype) CalcFitness() float64 {
	if p.knownFitness {
		return p.fitness
	}
	row := make([]int, len(p.combi)/2)
	col := make([]int, len(p.combi)/2)
	rowN := 0
	colN := 0
	count := 0
	mat := make([][]int, len(p.combi)/2)
	for i := 0; i < len(p.combi)/2; i++ {
		mat[i] = make([]int, len(p.combi)/2)
	}
	for i := 0; i < len(p.combi); i++ {

		if p.combi[i] < len(p.combi)/2 {
			row[rowN] = p.combi[i]
			rowN++
		} else {
			col[colN] = len(p.combi) - 1 - p.combi[i]
			colN++
		}
	}
	for i := 0; i < len(row); i++ {
		for j := 0; j < len(col); j++ {
			mat[i][j] = p.TriangularProblem.Values[row[i]][col[j]]
		}
	}

	for i := 0; i < len(row); i++ {
		for j := 0; j < i; j++ {
			if mat[i][j] == 0 {
				count++
			}
		}
	}
	p.knownFitness = true
	p.fitness = float64(count)
	return p.fitness
}

//No diagonal checking
func (p *phenotype) Good() bool {
	count := 0
	for i := 1; i < len(p.combi)/2; i++ {
		count += len(p.combi)/2 - i
	}
	return int(p.CalcFitness()) == count
}
