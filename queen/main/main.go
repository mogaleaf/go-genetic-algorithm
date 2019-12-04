package main

import (
	"fmt"
	"go-evol/evolution"
	selection2 "go-evol/evolution/selection"
	"go-evol/queen"
)

func main() {
	println("genetic algo")

	q := queen.QueensChessBoard{
		SizeChessBoard: 80,
	}
	populationSize := 500
	selection := selection2.RouletteWheelSelection{
		Selection: &selection2.Selection{
			Mu: populationSize,
		},
	}
	selectionSurvivor := selection2.BestSelection{
		Selection: &selection2.Selection{
			Mu: populationSize,
		},
	}
	parentsSelection := evolution.SelectionConfig{
		SelectionMethod: &selection,
		ProbabilityType: selection2.RANK,
		SP:              1.5,
	}
	survivorSelection := evolution.SelectionConfig{
		SelectionMethod: &selectionSurvivor,
		ProbabilityType: selection2.BEST,
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
