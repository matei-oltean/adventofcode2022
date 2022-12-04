package main

import (
	"os"
	"strconv"
)

var solvers map[int]func(string)

func init() {
	solvers = map[int]func(string){
		1: Solve01,
		2: Solve02,
		3: Solve03,
		4: Solve04,
	}
}

func main() {
	pb, _ := strconv.Atoi(os.Args[1])
	if solver, ok := solvers[pb]; ok && len(os.Args) >= 3 {
		solver(os.Args[2])
	}
}
