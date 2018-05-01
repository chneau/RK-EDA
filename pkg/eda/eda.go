package eda

import (
	"math"

	"github.com/bradfitz/slice"
	"github.com/chneau/RK-EDA/pkg/rk"
)

// Problem interface ...
type Problem interface {
	Evaluate(permuration []int) (float64, error)
}

// Sol ...
type Sol struct {
	Fitness *float64
	RK      rk.RK
}

// EDA of the RKEDA
type EDA struct {
	Problem        Problem
	PermSize       int
	PopSize        int
	TruncationSize int
	Elitism        int
	MaxEv          int
	Variance       float64
	CurrentEv      int
	BestSol        *Sol
	Elits          []Sol
}

// Run starts the RKEDA
func (e *EDA) Run() float64 {
	population := []Sol{}
	for len(population) < e.PopSize {
		s := &Sol{
			RK: rk.Random(e.PermSize),
		}
		fit := e.evaluate(s)
		s.Fitness = &fit
		population = append(population, *s)
		sortPop(population)
	}
	for e.CurrentEv < e.MaxEv {
		// elitism
		for i := 0; i <= e.Elitism; i++ {
			e.Elits = append(e.Elits, population[i])
		}
		sortPop(e.Elits)
		e.Elits = e.Elits[:e.Elitism]

		// trunc
		trunc := []Sol{}
		for i := 0; i < e.TruncationSize; i++ {
			trunc = append(trunc, population[i])
		}
		truncRK := []rk.RK{}
		for _, v := range trunc {
			truncRK = append(truncRK, v.RK)
		}
		baseRK := rk.Mean(truncRK)

		population = []Sol{}
		// generating next pop
		for i := 0; i < e.TruncationSize; i++ {
			s := &Sol{
				RK: baseRK.VarianceMutate(0.1),
			}
			fit := e.evaluate(s)
			s.Fitness = &fit
			population = append(population)
		}
	}
	return 0
}

func sortPop(pop []Sol) {
	slice.Sort(pop, func(i, j int) bool {
		return *pop[i].Fitness > *pop[j].Fitness
	})
}

func (e *EDA) evaluate(sol *Sol) float64 {
	e.CurrentEv = e.CurrentEv + 1
	f, err := e.Problem.Evaluate(sol.RK.Permutation())
	sol.Fitness = &f
	if e.BestSol == nil {
		e.BestSol = sol
	} else if f > *e.BestSol.Fitness {
		e.BestSol = sol
	}
	if err != nil {
		panic(err)
	}
	return f
}

// Reset the state of the EDA
func (e *EDA) Reset() {
	e.CurrentEv = 0
}

// Default return a defined version of RKEDA
func Default(problem Problem, permSize int) *EDA {
	ps := 50
	return &EDA{
		Problem:        problem,
		PermSize:       permSize,
		PopSize:        ps,
		TruncationSize: int(0.1 * float64(ps)),
		Elitism:        1,
		MaxEv:          1000 * permSize * permSize,
		Variance:       1 / (math.Pi * math.Log10(float64(permSize))),
	}
}
