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
	var result uint64 = 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		n := len(split)
		var nums = make([]int, n)
		for i, value := range split {
			v, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}

			nums[i] = v

		}

		result += magic(nums)
	}

	fmt.Printf("\nSolution is %v", result)
}

func magic(nums []int) uint64 {
	var r [12]int
	l := len(nums)
	current := 0
	start := 0
	end := l

	for current <= 11 {
		rem := 12 - current
		if rem <= l-start {
			// find max from start to end
			for j := start; j < end; j++ {
				if nums[j] > r[current] && l-j >= rem {
					r[current] = nums[j]
					start = j + 1
				}
			}
		}

		current++
	}

	var result uint64 = 0
	for n := range r {
		result = result*10 + uint64(r[n])
	}
	return result

}
