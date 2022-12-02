package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

func Solve02() {
	curdir, _ := os.Getwd()
	file, _ := os.Open(filepath.Join(curdir, "inputs", "02"))
	defer file.Close()
	countScore(file)
}

func countScore(file *os.File) {
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	scores := map[int32]int32{
		0: 3,
		1: 6,
		2: 0,
	}
	scores2 := map[int32]int32{
		0: 2,
		1: 0,
		2: 1,
	}
	var score1 int32 = 0
	var score2 int32 = 0
	a, _ := utf8.DecodeRuneInString("A")
	x, _ := utf8.DecodeRuneInString("X")
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		moves := strings.Split(txt, " ")
		op, _ := utf8.DecodeRuneInString(moves[0])
		mv, _ := utf8.DecodeRuneInString(moves[1])
		op -= a
		mv -= x
		score1 += mv + 1 + scores[(mv-op+3)%3]
		score2 += (op+scores2[mv])%3 + 1 + 3*mv
	}
	fmt.Println(score1)
	fmt.Println(score2)
}
