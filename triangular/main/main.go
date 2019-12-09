package main

import (
	"fmt"
	"go-evol/evolution"
	"go-evol/evolution/genes"
	selection2 "go-evol/evolution/selection"
	"go-evol/triangular"
)

func main() {
	println("genetic algo")

	q := triangular.NewTriangularProblem(
		[][]int{
			{
				2, 3, 0, 0,
			},
			{
				20, 30, 10, 0,
			},
			{
				35, 0, 0, 0,
			},
			{
				35, 10, 45, 46,
			},
		},
	)

	populationSize := 10
	selection := selection2.ProbabilitySelection{
		Selection: &selection2.Selection{
			Mu: populationSize,
		},
		AlgoType:        selection2.ROULETTE,
		ProbabilityType: selection2.RANK,
		SP:              1.5,
	}
	selectionSurvivor := selection2.TournamentSelection{
		Selection: &selection2.Selection{
			Mu: populationSize,
		},
		K: 2,
	}
	parentsSelection := evolution.SelectionConfig{
		SelectionMethod: &selection,
	}
	survivorSelection := evolution.SelectionConfig{
		SelectionMethod: &selectionSurvivor,
	}

	evolve := evolution.NewEvolve(
		q.NewRandTriangularGenotype,
		evolution.WithNumberIterationMax(500),
		evolution.WithPopulationSize(populationSize),
		evolution.WithParentsSelectionConfig(parentsSelection),
		evolution.WithNumberRecorder(&MyRecorder{}),
		evolution.WithSurvivorSelectionConfig(survivorSelection))

	resp, it := evolve.Evolve()

	if resp != nil {
		resp.Print()
		println(fmt.Sprintf("found solution in %d", it))
	} else {
		println("Not found")
	}

}

type MyRecorder struct {
}

func (r *MyRecorder) SelectedParents(gs []genes.GenotypeI, iter int)  {}
func (r *MyRecorder) CreatedOffspring(gs []genes.GenotypeI, iter int) {}

func (r *MyRecorder) MutatedOffspring(gs []genes.GenotypeI, iter int) {}
func (r *MyRecorder) NextGeneration(gs []genes.GenotypeI, iter int) {
	println("New generation")
	for i := 0; i < len(gs); i++ {
		gs[i].GetPhenotype().Print()
	}
}
