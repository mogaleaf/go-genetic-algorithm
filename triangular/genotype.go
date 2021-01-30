package triangular

import (
	"fmt"
	"go-evol/evolution/genes"
	"go-evol/evolution/genes/permutation"
	"go-evol/helper"
)

type TriangularProblem struct {
	size     int
	Values   [][]int
	maxZeros int
}

func NewTriangularProblem(values [][]int) *TriangularProblem {
	count := 0
	for i := 1; i < len(values); i++ {
		count += len(values) - i
	}
	return &TriangularProblem{
		maxZeros: count,
		size:     len(values),
		Values:   values,
	}
}

//Internal queen phenotype
type phenotype struct {
	combi             []float64
	TriangularProblem *TriangularProblem
	knownFitness      bool
	fitness           float64
}

func GetPhenotype(g genes.GenotypeI) genes.PhenotypeI {
	return &phenotype{
		combi:             g.(*permutation.Genotype).Permutation,
		TriangularProblem: g.(*permutation.Genotype).Data.(*TriangularProblem),
	}
}

func (q *TriangularProblem) NewRandTriangularGenotype() genes.GenotypeI {
	newGen := &permutation.Genotype{
		GetPhenotypeInternal: GetPhenotype,
		Permutation:          helper.Perm(q.size * 2),
		PermutationSize:      q.size * 2,
		RecombinationType:    permutation.CUT_CROSS_FILL,
		RecombinationRate:    0.9,
		MutationRate:         0.8,
		MutationType:         permutation.SWAP,
		Data:                 q,
	}
	return newGen
}

func (p *phenotype) Print() {
	row := make([]int, len(p.combi)/2)
	col := make([]int, len(p.combi)/2)
	rowN := 0
	colN := 0
	for i := 0; i < len(p.combi); i++ {
		print(fmt.Sprintf(" %0.2f ", p.combi[i]))
		if int(p.combi[i]) < len(p.combi)/2 {
			row[rowN] = int(p.combi[i])
			rowN++
		} else {
			col[colN] = len(p.combi) - 1 - int(p.combi[i])
			colN++
		}
	}
	println()
	print("ROW ORDER: [")
	for i := 0; i < len(row); i++ {
		print(fmt.Sprintf(" %d ", row[i]))
	}
	print("]")
	println()
	print("COL ORDER:[")
	for i := 0; i < len(col); i++ {
		print(fmt.Sprintf(" %d ", col[i]))
	}
	print("]")
	println()
	for i := 0; i < len(row); i++ {
		for j := 0; j < len(col); j++ {
			print(fmt.Sprintf(" %d ", p.TriangularProblem.Values[row[i]][col[j]]))
		}
		println()
	}
	println()
	println()
}
