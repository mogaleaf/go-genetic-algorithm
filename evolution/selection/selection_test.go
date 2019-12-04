package selection

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_selectParentFPS(t *testing.T) {
	m := []GenotypeI{
		newMockGenotypeI(withFuncGetPhenotype(func() PhenotypeI {
			return newMockPhenotypeI(withFuncCalcFitness(func() float64 {
				return 1.0
			}))
		})),
		newMockGenotypeI(withFuncGetPhenotype(func() PhenotypeI {
			return newMockPhenotypeI(withFuncCalcFitness(func() float64 {
				return 2.0
			}))
		})),
		newMockGenotypeI(withFuncGetPhenotype(func() PhenotypeI {
			return newMockPhenotypeI(withFuncCalcFitness(func() float64 {
				return 3.0
			}))
		})),
	}

	selectFPS(m, newMockSelectionI(withFuncselectPopulation(func(population []GenotypeI, a []float64) []GenotypeI {
		for i, ai := range a {
			switch i {
			case 0:
				assert.EqualValues(t, 1.0/6.0, ai)
			case 1:
				assert.EqualValues(t, 0.5, ai)
			case 2:
				assert.EqualValues(t, 1, ai)
			}
		}
		return nil
	})))
}

func Test_selectParentRank(t *testing.T) {
	m := []GenotypeI{
		newMockGenotypeI(withFuncGetPhenotype(func() PhenotypeI {
			return newMockPhenotypeI(withFuncCalcFitness(func() float64 {
				return 1.0
			}))
		})),
		newMockGenotypeI(withFuncGetPhenotype(func() PhenotypeI {
			return newMockPhenotypeI(withFuncCalcFitness(func() float64 {
				return 2.0
			}))
		})),
		newMockGenotypeI(withFuncGetPhenotype(func() PhenotypeI {
			return newMockPhenotypeI(withFuncCalcFitness(func() float64 {
				return 3.0
			}))
		})),
	}

	selectRank(m, 1.5, newMockSelectionI(withFuncgetMu(func() int {
		return 3
	}),
		withFuncselectPopulation(func(population []GenotypeI, a []float64) []GenotypeI {
			for i, ai := range a {
				switch i {
				case 2:
					assert.EqualValues(t, 1, ai)
				}
				println(fmt.Sprintf("%0.01f", ai))
			}
			return nil
		})))
}

func Test_selectPopulationRoulette(t *testing.T) {
	p1 := newMockGenotypeI(withFuncGetPhenotype(func() PhenotypeI {
		return newMockPhenotypeI(withFuncCalcFitness(func() float64 {
			return 1.0
		}))
	}))
	p2 := newMockGenotypeI(withFuncGetPhenotype(func() PhenotypeI {
		return newMockPhenotypeI(withFuncCalcFitness(func() float64 {
			return 2.0
		}))
	}))
	p3 := newMockGenotypeI(withFuncGetPhenotype(func() PhenotypeI {
		return newMockPhenotypeI(withFuncCalcFitness(func() float64 {
			return 3.0
		}))
	}))
	m := []GenotypeI{
		p1,
		p2,
		p3,
	}

	s := RouletteWheelSelection{
		Selection: &Selection{
			Mu: 3,
		},
	}

	a := []float64{0.2, 0.5, 1}
	population := s.selectPopulation(m, a)
	for _, p := range population {
		println(fmt.Sprintf("p1 %t", p == p1))
		println(fmt.Sprintf("p2 %t", p == p2))
		println(fmt.Sprintf("p3 %t", p == p3))
	}

}

func Test_selectPopulationSus(t *testing.T) {
	p1 := newMockGenotypeI(withFuncGetPhenotype(func() PhenotypeI {
		return newMockPhenotypeI(withFuncCalcFitness(func() float64 {
			return 1.0
		}))
	}))
	p2 := newMockGenotypeI(withFuncGetPhenotype(func() PhenotypeI {
		return newMockPhenotypeI(withFuncCalcFitness(func() float64 {
			return 2.0
		}))
	}))
	p3 := newMockGenotypeI(withFuncGetPhenotype(func() PhenotypeI {
		return newMockPhenotypeI(withFuncCalcFitness(func() float64 {
			return 3.0
		}))
	}))
	m := []GenotypeI{
		p1,
		p2,
		p3,
	}

	s := SusSelection{
		Selection: &Selection{
			Mu: 3,
		},
	}

	a := []float64{0.2, 0.5, 1}
	population := s.selectPopulation(m, a)
	for _, p := range population {
		println(fmt.Sprintf("p1 %t", p == p1))
		println(fmt.Sprintf("p2 %t", p == p2))
		println(fmt.Sprintf("p3 %t", p == p3))
	}

}

type mockGenotypeIOptions struct {
	funcGetPhenotype func() PhenotypeI
	funcMutate       func()
	funcRecombine    func(GenotypeI) []GenotypeI
	funcPrint        func()
}

var defaultMockGenotypeIOptions = mockGenotypeIOptions{
	funcGetPhenotype: func() PhenotypeI {
		return nil
	},
	funcMutate: func() {
	},
	funcRecombine: func(GenotypeI) []GenotypeI {
		return nil
	},
	funcPrint: func() {
	},
}

type mockGenotypeIOption func(*mockGenotypeIOptions)

func withFuncGetPhenotype(f func() PhenotypeI) mockGenotypeIOption {
	return func(o *mockGenotypeIOptions) {
		o.funcGetPhenotype = f
	}
}
func withFuncMutate(f func()) mockGenotypeIOption {
	return func(o *mockGenotypeIOptions) {
		o.funcMutate = f
	}
}
func withFuncRecombine(f func(GenotypeI) []GenotypeI) mockGenotypeIOption {
	return func(o *mockGenotypeIOptions) {
		o.funcRecombine = f
	}
}
func withFuncPrint(f func()) mockGenotypeIOption {
	return func(o *mockGenotypeIOptions) {
		o.funcPrint = f
	}
}

func (s *mockGenotypeI) GetPhenotype() PhenotypeI {
	return s.options.funcGetPhenotype()
}
func (s *mockGenotypeI) Mutate() {
	s.options.funcMutate()
}
func (s *mockGenotypeI) Recombine(g GenotypeI) []GenotypeI {
	return s.options.funcRecombine(g)
}
func (s *mockGenotypeI) Print() {
	s.options.funcPrint()
}

type mockGenotypeI struct {
	options mockGenotypeIOptions
}

func newMockGenotypeI(opt ...mockGenotypeIOption) GenotypeI {
	opts := defaultMockGenotypeIOptions
	for _, o := range opt {
		o(&opts)
	}
	return &mockGenotypeI{
		options: opts,
	}
}

type mockSelectionIOptions struct {
	funcselectPopulation func(population []GenotypeI, a []float64) []GenotypeI
	funcgetMu            func() int
}

var defaultMockSelectionIOptions = mockSelectionIOptions{
	funcselectPopulation: func(population []GenotypeI, a []float64) []GenotypeI {
		return nil
	},
	funcgetMu: func() int {
		return 0
	},
}

type mockSelectionIOption func(*mockSelectionIOptions)

func withFuncselectPopulation(f func(population []GenotypeI, a []float64) []GenotypeI) mockSelectionIOption {
	return func(o *mockSelectionIOptions) {
		o.funcselectPopulation = f
	}
}
func withFuncgetMu(f func() int) mockSelectionIOption {
	return func(o *mockSelectionIOptions) {
		o.funcgetMu = f
	}
}

func (s *mockSelectionI) selectPopulation(population []GenotypeI, a []float64) []GenotypeI {
	return s.options.funcselectPopulation(population, a)
}
func (s *mockSelectionI) getMu() int {
	return s.options.funcgetMu()
}

type mockSelectionI struct {
	options mockSelectionIOptions
}

func newMockSelectionI(opt ...mockSelectionIOption) SelectionI {
	opts := defaultMockSelectionIOptions
	for _, o := range opt {
		o(&opts)
	}
	return &mockSelectionI{
		options: opts,
	}
}

type mockPhenotypeIOptions struct {
	funcCalcFitness func() float64
	funcGood        func() bool
	funcPrint       func()
}

var defaultMockPhenotypeIOptions = mockPhenotypeIOptions{
	funcCalcFitness: func() float64 {
		return 0.0
	},
	funcGood: func() bool {
		return false
	},
	funcPrint: func() {
	},
}

type mockPhenotypeIOption func(*mockPhenotypeIOptions)

func withFuncCalcFitness(f func() float64) mockPhenotypeIOption {
	return func(o *mockPhenotypeIOptions) {
		o.funcCalcFitness = f
	}
}
func withFuncGood(f func() bool) mockPhenotypeIOption {
	return func(o *mockPhenotypeIOptions) {
		o.funcGood = f
	}
}
func withFuncPrintP(f func()) mockPhenotypeIOption {
	return func(o *mockPhenotypeIOptions) {
		o.funcPrint = f
	}
}

func (s *mockPhenotypeI) CalcFitness() float64 {
	return s.options.funcCalcFitness()
}
func (s *mockPhenotypeI) Good() bool {
	return s.options.funcGood()
}
func (s *mockPhenotypeI) Print() {
	s.options.funcPrint()
}

type mockPhenotypeI struct {
	options mockPhenotypeIOptions
}

func newMockPhenotypeI(opt ...mockPhenotypeIOption) PhenotypeI {
	opts := defaultMockPhenotypeIOptions
	for _, o := range opt {
		o(&opts)
	}
	return &mockPhenotypeI{
		options: opts,
	}
}
