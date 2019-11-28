package main

import (
	"fmt"
	"go-evol/evolution"
	"go-evol/queen"
)

func main() {
	println("genetic algo")

	q := queen.QueensChessBoard{
		SizeChessBoard: 12,
	}
	evolve, it := evolution.Evolve(q.NewRandQueenGenotype, 5000, 500)
	if evolve != nil {
		evolve.Print()
		println(fmt.Sprintf("found solution in %d", it))
	} else {
		println("Not found")
	}

}
