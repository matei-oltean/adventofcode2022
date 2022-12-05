package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func Solve05(input string) {
	curdir, _ := os.Getwd()
	file, _ := os.Open(filepath.Join(curdir, "inputs", input))
	defer file.Close()
	solve05(file)
}

func solve05(file *os.File) {
	var stacks [][]string
	var stacks2 [][]string
	move := false
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		if len(stacks) == 0 {
			stacks = make([][]string, (len(txt)+1)/4)
			stacks2 = make([][]string, (len(txt)+1)/4)
		}
		if len(txt) == 0 || txt[1] == '1' {
			move = true
			continue
		}
		if !move {
			for i := 0; i < len(stacks); i++ {
				index := 4*i + 1
				if txt[index] != ' ' {
					stacks[i] = append([]string{string(txt[index])}, stacks[i]...)
					stacks2[i] = append([]string{string(txt[index])}, stacks2[i]...)
				}
			}
		} else {
			re := regexp.MustCompile(`(\d+)`)
			match := re.FindAllString(txt, -1)
			nb, _ := strconv.Atoi(match[0])
			from, _ := strconv.Atoi(match[1])
			from -= 1
			to, _ := strconv.Atoi(match[2])
			to -= 1
			for i := 0; i < nb; i++ {
				stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
				stacks[from] = stacks[from][:len(stacks[from])-1]
			}
			stacks2[to] = append(stacks2[to], stacks2[from][len(stacks2[from])-nb:]...)
			stacks2[from] = stacks2[from][:len(stacks2[from])-nb]
		}
	}
	for _, stack := range stacks {
		fmt.Print(stack[len(stack)-1])
	}
	fmt.Println()
	for _, stack := range stacks2 {
		fmt.Print(stack[len(stack)-1])
	}
}
