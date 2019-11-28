package queen

import "testing"

func TestPhenotype_CalcFitness(t *testing.T) {
	t.Run("diag", func(t *testing.T) {
		pTest := phenotype{
			//board: []int{7,5,4,3,1,6,0,2},
			//board:[]int{6,4,2,7,5,3,1,0},
			board: []int{2, 7, 3, 6, 0, 5, 1, 4},
		}
		pTest.Print()
		//pTest := NewRandQueenGenotype().GetPhenotype()
		//pTest.Print()
		fitness := pTest.CalcFitness()
		println(fitness)
	})

}
