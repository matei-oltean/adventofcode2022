package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Solve04(input string) {
	curdir, _ := os.Getwd()
	file, _ := os.Open(filepath.Join(curdir, "inputs", input))
	defer file.Close()
	solve04(file)
}

func solve04(file *os.File) {
	dup := 0
	inter := 0
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		elves := strings.Split(txt, ",")
		e1 := strings.Split(elves[0], "-")
		e2 := strings.Split(elves[1], "-")
		e11, _ := strconv.Atoi(e1[0])
		e12, _ := strconv.Atoi(e1[1])
		e21, _ := strconv.Atoi(e2[0])
		e22, _ := strconv.Atoi(e2[1])
		if (e11 <= e21 && e12 >= e22) || (e21 <= e11 && e22 >= e12) {
			dup += 1
		}
		if (e12 >= e21 && e11 <= e21) || (e22 >= e11 && e21 <= e11) {
			inter += 1
		}
	}
	fmt.Println(dup)
	fmt.Println(inter)
}
