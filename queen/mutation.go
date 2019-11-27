package queen

import "go-evol/helper"

func (g *Genotype) Mutate() {
	// 80% chance to mutate
	shouldMutate := helper.GenerateUintNumber(100) < 80
	if shouldMutate {
		//swap mutation
		number1 := helper.GenerateUintNumber(8)
		number2 := helper.GenerateUintNumber(8)
		temp := g.board[number1]
		g.board[number1] = g.board[number2]
		g.board[number2] = temp
	}
}
