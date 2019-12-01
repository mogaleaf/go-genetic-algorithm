package main

import (
	"fmt"
	"go-evol/evolution"
	"go-evol/queen"
)

func main() {
	println("genetic algo")

	q := queen.QueensChessBoard{
		SizeChessBoard: 10,
	}

	selection := evolution.RouletteWheelSelection{
		Selection: &evolution.Selection{
			Mu: 500,
		},
	}
	config := evolution.EvolutionConfig{
		ParentsSelectionConfig: evolution.SelectionConfig{
			SelectionMethod: &selection,
			ProbabilityType: evolution.RANK,
			SP:              1.5,
		},
		NumberIterationMax: 50000,
		PopulationSize:     500,
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
