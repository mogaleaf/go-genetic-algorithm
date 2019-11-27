package queen

import (
	"go-evol/evolution"
	"go-evol/helper"
)

func (p1 *Genotype) Recombine(p2 evolution.GenotypeI) []evolution.GenotypeI {
	swapK := helper.GenerateUintNumber(sizeChessBoard)
	if swapK == 0 {
		return []evolution.GenotypeI{
			&Genotype{
				board: p1.board,
			},
			&Genotype{
				board: p2.(*Genotype).board,
			},
		}
	}
	k1 := Genotype{
		board: make([]int, sizeChessBoard),
	}
	k2 := Genotype{
		board: make([]int, sizeChessBoard),
	}
	for j := uint64(0); j < swapK; j++ {
		k1.board[j] = p1.board[j]
		k2.board[j] = p2.(*Genotype).board[j]
	}
	fillWithMissing(k1.board, p2.(*Genotype).board, swapK)
	fillWithMissing(k2.board, p1.board, swapK)
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

func fillWithMissing(toFill []int, with []int, swapK uint64) {
	j := 0
	i := swapK
	for i < sizeChessBoard && j < sizeChessBoard {
		if !contains(toFill, with[j]) {
			toFill[i] = with[j]
			i++
		}
		j++
	}
}
