package main

import (
	"fmt"
	"go-evol/evolution"
	"go-evol/onemax"
)

func main() {
	println("genetic algo")

	q := onemax.OneMaxProblem{
		L: 25,
	}
	populationSize := 100
	selection := evolution.RouletteWheelSelection{
		Selection: &evolution.Selection{
			Mu: populationSize,
		},
	}
	parentsSelection := evolution.SelectionConfig{
		SelectionMethod: &selection,
		ProbabilityType: evolution.FPS,
	}
	survivorSelection := evolution.SelectionConfig{
		ProbabilityType: evolution.REPLACE,
	}

	evolve := evolution.NewEvolve(
		q.NewRandQueenGenotype,
		evolution.WithNumberIterationMax(100),
		evolution.WithPopulationSize(populationSize),
		evolution.WithParentsSelectionConfig(parentsSelection),
		evolution.WithSurvivorSelectionConfig(survivorSelection))

	resp, it := evolve.Evolve()

	if resp != nil {
		resp.Print()
		println(fmt.Sprintf("found solution in %d", it))
	} else {
		println("Not found")
	}

}
