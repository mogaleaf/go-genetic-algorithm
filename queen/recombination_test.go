package queen

import "testing"

func TestGenotype_Recombine(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		q := QueensChessBoard{
			SizeChessBoard: 8,
		}
		p1 := q.NewRandQueenGenotype()
		p2 := q.NewRandQueenGenotype()
		p1.Print()
		p2.Print()
		k := p1.Recombine(p2)
		for _, ki := range k {
			ki.Print()
		}
	})
}
