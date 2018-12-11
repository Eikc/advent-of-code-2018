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

	var sum int
	for sc.Scan() {
		line := sc.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("could not parse int: %v", err)
		}
		sum += number
	}

	fmt.Println(sum)
}
