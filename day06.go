package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func Solve06(input string) {
	curdir, _ := os.Getwd()
	file, _ := os.Open(filepath.Join(curdir, "inputs", input))
	defer file.Close()
	solve06(file, 14)
}

func solve06(file *os.File, unique int) {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		counts := make([]int, 27)
		i := 0
		for ; i < unique; i++ {
			counts[Priority(rune(txt[i]))] += 1
		}
		for ; i < len(txt); i++ {
			stop := false
			for _, c := range counts {
				if c > 1 {
					counts[Priority(rune(txt[i]))] += 1
					counts[Priority(rune(txt[i-unique]))] -= 1
					stop = true
					break
				}
			}
			if !stop {
				fmt.Println(i)
				break
			}
		}
	}
}
