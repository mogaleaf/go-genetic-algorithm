package binary

import (
	"testing"
)

func Test_Genes(t *testing.T) {
	t.Run("print", func(t *testing.T) {

		g := Genotype{
			L:            8,
			Value:        []uint8{0, 0, 1, 0, 1, 1, 1, 0},
			MutationType: FLIP,
		}

		g.Print()
		g.mutateFlip()
		g.Print()
	})

}
