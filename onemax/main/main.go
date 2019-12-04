package main

import (
	"fmt"
	"go-evol/evolution"
	selection2 "go-evol/evolution/selection"
	"go-evol/onemax"
)

func main() {
	println("genetic algo")

	q := onemax.OneMaxProblem{
		L: 75,
	}
	populationSize := 100
	selection := selection2.TournamentSelection{
		Selection: &selection2.Selection{
			Mu: populationSize,
		},
		K:2,
	}
	parentsSelection := evolution.SelectionConfig{
		SelectionMethod: &selection,
		ProbabilityType: selection2.TOURNAMENT,
	}
	survivorSelection := evolution.SelectionConfig{
		ProbabilityType: selection2.REPLACE,
	}

	evolve := evolution.NewEvolve(
		q.NewRandOneMaxGenotype,
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
