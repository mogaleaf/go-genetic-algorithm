package main

import (
	"fmt"
	"go-evol/evolution"
	"go-evol/queen"
)

func main() {
	println("genetic algo")

	q := queen.QueensChessBoard{
		SizeChessBoard: 80,
	}
	populationSize := 500
	selection := evolution.RouletteWheelSelection{
		Selection: &evolution.Selection{
			Mu: populationSize,
		},
	}
	selectionSurvivor := evolution.BestSelection{
		Selection: &evolution.Selection{
			Mu: populationSize,
		},
	}
	parentsSelection := evolution.SelectionConfig{
		SelectionMethod: &selection,
		ProbabilityType: evolution.RANK,
		SP:              1.5,
	}
	survivorSelection := evolution.SelectionConfig{
		SelectionMethod: &selectionSurvivor,
		ProbabilityType: evolution.BEST,
	}

	evolve := evolution.NewEvolve(
		q.NewRandQueenGenotype,
		evolution.WithNumberIterationMax(50000),
		evolution.WithPopulationSize(populationSize),
		evolution.WithParentsSelectionConfig(parentsSelection),
		evolution.WithSurvivorSelectionConfig(survivorSelection))

	resp, it := evolve.Evolve()

	if evolve != nil {
		resp.Print()
		println(fmt.Sprintf("found solution in %d", it))
	} else {
		println("Not found")
	}

}
