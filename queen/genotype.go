package queen

import (
	"fmt"
	"go-evol/evolution"
	"math/rand"
	"time"
)

type QueensChessBoard struct {
	SizeChessBoard int
}

type phenotype struct {
	board []int
}

type genotype struct {
	board []int
}

func (g *genotype) GetPhenotype() evolution.PhenotypeI {
	phen := phenotype{
		board: g.board,
	}
	return &phen
}

func (q *QueensChessBoard) NewRandQueenGenotype() evolution.GenotypeI {
	newGen := genotype{}
	newGen.board = make([]int, q.SizeChessBoard)
	newGen.board = make([]int, q.SizeChessBoard)
	for i := 0; i < q.SizeChessBoard; i++ {
		newGen.board[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(newGen.board), func(i, j int) { newGen.board[i], newGen.board[j] = newGen.board[j], newGen.board[i] })
	return &newGen
}

func (g *genotype) Print() {
	for j := 0; j < len(g.board); j++ {
		print(fmt.Sprintf("%d", g.board[j]))
	}
	println()
}

func (p *phenotype) Print() {
	for i := 0; i < len(p.board); i++ {
		for j := 0; j < len(p.board); j++ {
			if j == p.board[i] {
				print(" 1 ")
			} else {
				print(" 0 ")
			}
		}
		println()
	}

}
