package eda

import (
	"log"
	"math"
	"sync"

	"github.com/bradfitz/slice"
	"github.com/chneau/RK-EDA/pkg/rk"
	"github.com/chneau/limiter"
)

// Problem interface ...
type Problem interface {
	Evaluate(permuration []int) (float64, error)
}

// Cooler interface ...
type Cooler interface {
	NewTemperature(improvements int) float64
}

// Sol ...
type Sol struct {
	RK      rk.RK
	Fitness *float64
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
	Elits          []*Sol
	population     []*Sol
	mutex          sync.Mutex
}

// Run starts the RKEDA
func (e *EDA) Run() Sol {
	for len(e.population) < e.PopSize {
		s := &Sol{
			RK: rk.Random(e.PermSize),
		}
		e.evaluate(s)
		e.population = append(e.population, s)
	}
	for e.CurrentEv < e.MaxEv {
		newPopulation := []*Sol{}
		// elitism
		e.elitism()

		// trunc
		trunc, baseRK := e.trunc()
		// fill with elits
		for _, sol := range e.Elits {
			s := &Sol{RK: sol.RK}
			newPopulation = append(newPopulation, s)
		}
		// fill with trunced
		for i := 0; len(newPopulation) < e.PopSize; i++ {
			s := &Sol{RK: trunc[i%e.TruncationSize].RK}
			newPopulation = append(newPopulation, s)

		}
		limit := limiter.New(50)
		for i := range newPopulation {
			sol := newPopulation[i]
			limit.Execute(func() {
				sol.RK = sol.RK.VarianceMutate(baseRK, e.Variance)
				e.evaluate(sol)
			})
		}
		limit.Wait()
		e.population = newPopulation
	}
	return *e.BestSol
}

func (e *EDA) trunc() ([]*Sol, rk.RK) {
	trunc := []*Sol{}
	for i := 0; i < e.TruncationSize; i++ {
		trunc = append(trunc, e.population[i])
	}
	truncRK := []rk.RK{}
	for _, v := range trunc {
		truncRK = append(truncRK, v.RK)
	}
	baseRK := rk.Mean(truncRK)
	return trunc, baseRK
}

func (e *EDA) elitism() {
	sortPop(e.population)
	for i := 0; i <= e.Elitism; i++ {
		e.Elits = append(e.Elits, e.population[i])
	}
	sortPop(e.Elits)
	e.Elits = e.Elits[:e.Elitism]
}

func sortPop(pop []*Sol) {
	slice.Sort(pop, func(i, j int) bool {
		return *pop[i].Fitness < *pop[j].Fitness
	})
}

func (e *EDA) evaluate(sol *Sol) {
	e.mutex.Lock()
	e.CurrentEv = e.CurrentEv + 1
	currentEv := e.CurrentEv
	e.mutex.Unlock()
	if currentEv%100000 == 0 {
		log.Println(currentEv, "/", e.MaxEv, "=>", *e.BestSol.Fitness)
	}
	f, err := e.Problem.Evaluate(sol.RK.Permutation())
	sol.Fitness = &f
	if e.BestSol == nil {
		e.BestSol = sol
	} else if f < *e.BestSol.Fitness {
		e.BestSol = sol
	}
	if err != nil {
		panic(err)
	}
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
