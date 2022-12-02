package main

import (
	"os"
	"strconv"
)

var solvers map[int]func()

func init() {
	solvers = map[int]func(){
		1: Solve01,
		2: Solve02,
	}
}

func main() {
	pb, _ := strconv.Atoi(os.Args[len(os.Args)-1])
	if solver, ok := solvers[pb]; ok {
		solver()
	}
}
