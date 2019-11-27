package queen

import "testing"

func TestGenotype_Recombine(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		p1 := NewRandQueenGenotype()
		p2 := NewRandQueenGenotype()
		p1.Print()
		p2.Print()
		k := p1.Recombine(p2)
		for _, ki := range k {
			ki.Print()
		}
	})
}
