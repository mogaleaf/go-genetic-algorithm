package main

import (
	"fmt"
	"go-evol/evolution"
	"go-evol/queen"
)

func main() {
	println("genetic algo")

	q := queen.QueensChessBoard{
		SizeChessBoard: 30,
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
	config := evolution.EvolutionConfig{
		ParentsSelectionConfig: evolution.SelectionConfig{
			SelectionMethod: &selection,
			ProbabilityType: evolution.RANK,
			SP:              1.5,
		},
		SurvivorSelectionConfig: evolution.GenerationSelectionConfig{
			SelectionConfig: &evolution.SelectionConfig{
				SelectionMethod: &selectionSurvivor,
				ProbabilityType: evolution.BEST,
			},
		},
		NumberIterationMax: 50000,
		PopulationSize:     populationSize,
		Create:             q.NewRandQueenGenotype,
	}

	evolve, it := evolution.Evolve(config)

	if evolve != nil {
		evolve.Print()
		println(fmt.Sprintf("found solution in %d", it))
	} else {
		println("Not found")
	}

}
