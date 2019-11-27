package queen

import (
	"fmt"
	"go-evol/evolution"
	"math/rand"
	"time"
)

const (
	sizeChessBoard = 15
)

type Phenotype struct {
	board []int
}

type Genotype struct {
	board []int
}

func (g *Genotype) GetPhenotype() evolution.PhenotypeI {
	phen := Phenotype{
		board: g.board,
	}
	return &phen
}

func NewRandQueenGenotype() evolution.GenotypeI {
	newGen := Genotype{}
	newGen.board = make([]int, sizeChessBoard)
	newGen.board = make([]int, sizeChessBoard)
	for i := 0; i < sizeChessBoard; i++ {
		newGen.board[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(newGen.board), func(i, j int) { newGen.board[i], newGen.board[j] = newGen.board[j], newGen.board[i] })
	return &newGen
}

func (g *Genotype) Print() {
	for j := 0; j < sizeChessBoard; j++ {
		print(fmt.Sprintf("%d", g.board[j]))
	}
	println()
}

func (p *Phenotype) Print() {
	for i := 0; i < sizeChessBoard; i++ {
		for j := 0; j < sizeChessBoard; j++ {
			if j == p.board[i] {
				print(" 1 ")
			} else {
				print(" 0 ")
			}
		}
		println()
	}

}
