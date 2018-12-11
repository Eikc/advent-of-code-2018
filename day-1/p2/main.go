package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("open file error: %v", err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	var numbers []int
	for sc.Scan() {
		line := sc.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("could not parse int: %v", err)
		}
		numbers = append(numbers, number)
	}

	var sum int
	seenBefore := map[int]bool{0: true}
	for {
		for _, number := range numbers {
			sum += number
			if _, ok := seenBefore[sum]; ok {
				fmt.Println(sum)
				return
			}
			seenBefore[sum] = true
		}
	}
}
