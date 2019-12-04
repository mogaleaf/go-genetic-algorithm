package main

import (
	"fmt"
	"go-evol/evolution"
	"go-evol/evolution/genes"
	selection2 "go-evol/evolution/selection"
	"go-evol/onemax"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	println("genetic algo")

	q := onemax.OneMaxProblem{
		L: 25,
	}
	populationSize := 100
	selection := selection2.TournamentSelection{
		Selection: &selection2.Selection{
			Mu: populationSize,
		},
		K: 2,
	}
	selection2 := selection2.ReplaceSelection{
		Selection: &selection2.Selection{
			Mu: populationSize,
		},
	}
	parentsSelection := evolution.SelectionConfig{
		SelectionMethod: &selection,
	}
	survivorSelection := evolution.SelectionConfig{
		SelectionMethod: &selection2,
	}

	recorder := NewMyRecorder()
	evolve := evolution.NewEvolve(
		q.NewRandOneMaxGenotype,
		evolution.WithNumberIterationMax(100),
		evolution.WithPopulationSize(populationSize),
		evolution.WithParentsSelectionConfig(parentsSelection),
		evolution.WithSurvivorSelectionConfig(survivorSelection),
		evolution.WithNumberRecorder(recorder))

	resp, it := evolve.Evolve()

	if resp != nil {
		resp.Print()
		println(fmt.Sprintf("found solution in %d", it))
	} else {
		println("Not found")
	}
	recorder.Plot.Save(10*vg.Centimeter, 5*vg.Centimeter, "evolution.png")

}

func NewMyRecorder() *MyRecorder {
	r := MyRecorder{}
	r.Plot, _ = plot.New()
	r.Plot.Title.Text = "population evolution"
	r.Plot.Y.Label.Text = "mean fitness"
	r.Plot.X.Label.Text = "iter"
	r.Plot.Add(plotter.NewGrid())
	return &r
}

type MyRecorder struct {
	Plot *plot.Plot
}

func (r *MyRecorder) SelectedParents(gs []genes.GenotypeI, iter int)  {}
func (r *MyRecorder) CreatedOffspring(gs []genes.GenotypeI, iter int) {}

func (r *MyRecorder) MutatedOffspring(gs []genes.GenotypeI, iter int) {}
func (r *MyRecorder) NextGeneration(gs []genes.GenotypeI, iter int) {
	mean := 0.0
	for _, g := range gs {
		mean += g.GetPhenotype().CalcFitness()
	}
	mean /= float64(len(gs))
	pts := make(plotter.XYs, 1)
	pts[0] = plotter.XY{
		X: float64(iter),
		Y: mean,
	}
	_, points, _ := plotter.NewLinePoints(pts)

	r.Plot.Add(points)
}
