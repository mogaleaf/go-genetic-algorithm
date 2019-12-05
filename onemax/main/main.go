package main

import (
	"fmt"
	"go-evol/evolution"
	"go-evol/evolution/genes"
	selection2 "go-evol/evolution/selection"
	"go-evol/helper"
	"go-evol/onemax"

	"golang.org/x/image/colornames"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	//Tournament selection Mean: 62.500 , sd: 6.317
	// FPS using Rank SUS Mean: 94.300 , sd: 12.582
	numberOfRun := 3
	runs, count := iterateAlgo(numberOfRun)
	mean := float64(count) / float64(numberOfRun)
	deviationInt := helper.StandardDeviationInt(runs, count)
	println(fmt.Sprintf("Mean: %0.3f , sd: %0.3f", mean, deviationInt))
}

func iterateAlgo(numberOfRun int) ([]int, int) {
	count := 0
	time := make([]int, numberOfRun)
	for i := 0; i < numberOfRun; i++ {
		time[i] = launch(i)
		count += time[i]
	}
	return time, count
}

func launch(iterNumber int) int {
	println("genetic algo")

	q := onemax.OneMaxProblem{
		L: 75,
	}
	populationSize := 100
	parentSelection := selection2.ProbabilitySelection{
		Selection: &selection2.Selection{
			Mu: populationSize,
		},
		ProbabilityType: selection2.FPS,
		AlgoType:        selection2.SUS,
		SP:              1.5,

		//K: 2,
	}
	survivorSelection := selection2.ReplaceSelection{
		Selection: &selection2.Selection{
			Mu: populationSize,
		},
	}

	recorder := NewMyRecorder()
	evolve := evolution.NewEvolve(
		q.NewRandOneMaxGenotype,
		evolution.WithNumberIterationMax(100),
		evolution.WithPopulationSize(populationSize),
		evolution.WithParentsSelectionConfig(evolution.SelectionConfig{
			SelectionMethod: &parentSelection,
		}),
		evolution.WithSurvivorSelectionConfig(evolution.SelectionConfig{
			SelectionMethod: &survivorSelection,
		}),
		evolution.WithNumberRecorder(recorder))

	resp, it := evolve.Evolve()

	if resp != nil {
		resp.Print()
		println(fmt.Sprintf("found solution in %d", it))
	} else {
		println("Not found")
	}
	recorder.Plot.Save(10*vg.Centimeter, 5*vg.Centimeter, fmt.Sprintf("evolution%d.png", iterNumber))
	return it
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
	worst := 100.0
	best := 0.0
	for _, g := range gs {
		fitness := g.GetPhenotype().CalcFitness()
		mean += fitness
		if fitness < worst {
			worst = fitness
		}
		if fitness > best {
			best = fitness
		}
	}
	mean /= float64(len(gs))
	ptsMean := make(plotter.XYs, 1)
	ptsMean[0] = plotter.XY{
		X: float64(iter),
		Y: mean,
	}
	_, points, _ := plotter.NewLinePoints(ptsMean)
	points.Color = colornames.Blue

	ptsWorst := make(plotter.XYs, 1)
	ptsWorst[0] = plotter.XY{
		X: float64(iter),
		Y: worst,
	}
	_, pointsW, _ := plotter.NewLinePoints(ptsWorst)
	pointsW.Color = colornames.Red

	ptsBest := make(plotter.XYs, 1)
	ptsBest[0] = plotter.XY{
		X: float64(iter),
		Y: best,
	}
	_, pointsB, _ := plotter.NewLinePoints(ptsBest)
	pointsB.Color = colornames.Green

	r.Plot.Add(points)
	r.Plot.Add(pointsW)
	r.Plot.Add(pointsB)
}
