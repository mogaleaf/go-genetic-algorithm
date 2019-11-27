package main

import (
	"fmt"
	"go-evol/evolution"
	"go-evol/queen"
)

func main() {
	println("genetic algo")

	evolve, it := evolution.Evolve(queen.NewRandQueenGenotype, 5000, 500)
	if evolve != nil {
		evolve.Print()
		println(fmt.Sprintf("found solution in %d", it))
	} else {
		println("Not found")
	}

}
