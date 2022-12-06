package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"unicode"
	"unicode/utf8"
)

func Priority(r rune) int32 {
	if unicode.IsLower(r) {
		a, _ := utf8.DecodeRuneInString("a")
		return r - a + 1
	}
	A, _ := utf8.DecodeRuneInString("A")
	return r - A + 27
}

func Solve03(input string) {
	curdir, _ := os.Getwd()
	file, _ := os.Open(filepath.Join(curdir, "inputs", input))
	defer file.Close()
	countPriorities(file)
}

func countPriorities(file *os.File) {
	var empty struct{}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var prio1 int32 = 0
	var prio2 int32 = 0
	intersects := make(map[int32]struct{})
	i := 0
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		counts := make(map[rune]struct{})
		intersects2 := make(map[int32]struct{})
		for index, char := range txt {
			if index < len(txt)/2 {
				counts[char] = empty
			} else {
				if _, ok := counts[char]; ok {
					prio1 += Priority(char)
					delete(counts, char)
				}
			}
			prio := Priority(char)
			if i == 0 {
				intersects[prio] = empty
			} else {
				if _, ok := intersects[prio]; ok {
					if i == 1 {
						intersects2[prio] = empty
					} else {
						prio2 += prio
						delete(intersects, prio)
					}
				}
			}
		}
		if i == 1 {
			intersects = intersects2
		} else if i == 2 {
			intersects = make(map[int32]struct{})
		}
		i = (i + 1) % 3
	}
	fmt.Println(prio1)
	fmt.Println(prio2)
}
