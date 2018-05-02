package main

import (
	"log"
	"runtime/debug"

	"github.com/chneau/RK-EDA/pkg/eda"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

func gcstats() {
	gcstats := debug.GCStats{}
	debug.ReadGCStats(&gcstats)
	log.Println("gcstats.PauseTotal", gcstats.PauseTotal)
}

// Digits ...
type Digits struct{}

// Evaluate ...
func (d Digits) Evaluate(permuration []int) (float64, error) {
	sum := 0
	for i := 0; i < len(permuration); i++ {
		if permuration[i] == i {
			continue
		}
		sum = sum + 1
	}
	return float64(sum), nil
}

func main() {
	problem := Digits{}
	r := eda.Default(problem, 20)
	best := r.Run()
	log.Println("best", *best.Fitness, best.RK.Permutation())
	gcstats()
}
