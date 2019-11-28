package queen

import "go-evol/helper"

func (g *genotype) Mutate() {
	// 80% chance to mutate
	shouldMutate := helper.GenerateUintNumber(100) < 80
	if shouldMutate {
		//swap mutation
		number1 := helper.GenerateUintNumber(len(g.board))
		number2 := helper.GenerateUintNumber(len(g.board))
		temp := g.board[number1]
		g.board[number1] = g.board[number2]
		g.board[number2] = temp
	}
}
