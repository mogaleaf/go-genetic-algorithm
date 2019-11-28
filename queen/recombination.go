package queen

import (
	"go-evol/evolution"
	"go-evol/helper"
)

func (p1 *genotype) Recombine(p2 evolution.GenotypeI) []evolution.GenotypeI {
	swapK := helper.GenerateUintNumber(len(p1.board))
	if swapK == 0 {
		return []evolution.GenotypeI{
			&genotype{
				board: p1.board,
			},
			&genotype{
				board: p2.(*genotype).board,
			},
		}
	}
	k1 := genotype{
		board: make([]int, len(p1.board)),
	}
	k2 := genotype{
		board: make([]int, len(p1.board)),
	}
	for j := uint64(0); j < swapK; j++ {
		k1.board[j] = p1.board[j]
		k2.board[j] = p2.(*genotype).board[j]
	}
	fillWithMissing(k1.board, p2.(*genotype).board, int(swapK))
	fillWithMissing(k2.board, p1.board, int(swapK))
	return []evolution.GenotypeI{&k1, &k2}
}

func contains(slice []int, value int) bool {
	for j := 0; j < len(slice); j++ {
		if slice[j] == value {
			return true
		}
	}
	return false
}

func fillWithMissing(toFill []int, with []int, swapK int) {
	j := 0
	i := swapK
	for i < len(with) && j < len(with) {
		if !contains(toFill, with[j]) {
			toFill[i] = with[j]
			i++
		}
		j++
	}
}
