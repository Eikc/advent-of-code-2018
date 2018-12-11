package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("part one")
	partOne()

	fmt.Println("part two")
	partTwo()
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	var twos, threes int
	for sc.Scan() {
		line := sc.Text()
		charCount := map[rune]int{}
		for _, c := range line {
			charCount[c]++
		}

		var hasTwo, hasThree bool
		for _, count := range charCount {
			if count == 2 {
				hasTwo = true
			}
			if count == 3 {
				hasThree = true
			}
		}

		if hasTwo {
			twos++
		}
		if hasThree {
			threes++
		}
	}

	result := twos * threes
	fmt.Println(result)
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	var words []string
	for sc.Scan() {
		words = append(words, strings.TrimSpace(sc.Text()))
	}

	for i, a := range words {
		for _, b := range words[i+1:] {
			common, ok := compare(a, b)
			if ok {
				fmt.Println(common)
				return
			}
		}
	}
}

func compare(a, b string) (string, bool) {
	var missedAt int
	var misses int

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			misses++
			missedAt = i
		}
	}

	if misses == 1 {
		return a[:missedAt] + a[missedAt+1:], true
	}

	return "", false
}
