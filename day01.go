package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func Solve01() {
	curdir, _ := os.Getwd()
	file, _ := os.Open(filepath.Join(curdir, "inputs", "01"))
	defer file.Close()
	cals := countCal(file)
	sort.Ints(cals)
	fmt.Println(cals[len(cals)-1])
	fmt.Println(cals[len(cals)-1] + cals[len(cals)-2] + cals[len(cals)-3])
}

func countCal(file *os.File) (cals []int) {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	acc := 0
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		if txt == "" {
			cals = append(cals, acc)
			acc = 0
		} else {
			cal, _ := strconv.Atoi(txt)
			acc += cal
		}
	}
	return
}
