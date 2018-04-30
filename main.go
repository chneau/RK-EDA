package main

import (
	"log"

	"github.com/chneau/RK-EDA/pkg/rk"
)

func main() {
	rk := rk.RK{0, 0, 0}
	log.Println("before", rk)
	perm := rk.Permutation()
	log.Println("rk", rk)
	log.Println("perm", perm)

}
