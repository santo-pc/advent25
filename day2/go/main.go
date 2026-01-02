package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	result := 0

	var ranges []string

	for scanner.Scan() {
		line := scanner.Text()

		ranges = append(ranges, strings.Split(line, ",")...)
	}

	for _, r := range ranges {
		split := strings.Split(r, "-")
		start, err := strconv.Atoi(strings.TrimSpace(split[0]))
		if err != nil {
			fmt.Println("Could not convert number", err)
		}
		end, err := strconv.Atoi(strings.TrimSpace(split[1]))

		if err != nil {
			fmt.Println("Could not convert string", err)
		}

		for i := start; i <= end; i++ {
			if isInvalid(i) {
				result += i
			}
		}

	}

	fmt.Printf("\nSolution: %v", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func isInvalid(num int) bool {
	count := 0
	num_s := strconv.Itoa(num)
	n := 1
	length := len(num_s)
	mid := length / 2
	// we split the string up to two pats when n = half
	for n <= mid {
		count = 0
		start := n
		subs := num_s[0:n]
		match := true

		// we split by n and compare that all chunks are the same
		for start+n <= length {
			slice := num_s[start : start+n]

			if subs != slice {
				match = false
				break
			}
			count++
			start += n
		}

		// Check we don't have a remaining chunk smaller than n
		gap := start - length
		if gap == 0 && match {
			return count >= 1
		}

		// reset
		count = 0
		n++

	}

	return false

}
