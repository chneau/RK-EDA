package main

import (
	"log"
	"runtime/debug"

	"github.com/chneau/RK-EDA/pkg/rk"
)

func gcstats() {
	gcstats := debug.GCStats{}
	debug.ReadGCStats(&gcstats)
	log.Println("gcstats.PauseTotal", gcstats.PauseTotal)
}

func main() {
	rk := rk.RK{0, 0, 0}
	log.Println("before", rk)
	perm := rk.Permutation()
	log.Println("rk", rk)
	log.Println("perm", perm)

	gcstats()
}
