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
	// solution1()
	solution2()
}

type Operation func(a, b uint64) uint64

var mat = [][]string{}

var operationsMap = map[string]Operation{
	"*": func(a, b uint64) uint64 {
		fmt.Printf("%v * %v = %v\n", a, b, a*b)
		return a * b
	},
	"+": func(a, b uint64) uint64 {
		fmt.Printf("%v + %v = %v\n", a, b, a+b)
		return a + b
	},
}

type operator struct {
	operator string
	start    int
}

func solution2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	var operators = []operator{}
	var maxL = 0

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "")
		op := split[0]

		// append operator line
		if op == "+" || op == "*" {
			fmt.Printf("\n")
			for i, w := range split {
				// text := strings.TrimSpace(w)
				if w == "*" || w == "+" {
					fmt.Printf("%v", w)
					op := operator{
						operator: w,
						start:    i,
					}

					operators = append(operators, op)
				}
			}
		} else { // append numbers
			numsLine := []string{}
			for _, w := range split {
				numsLine = append(numsLine, w)
			}

			mat = append(mat, numsLine)
			if len(numsLine) > maxL {
				maxL = len(numsLine)
			}
		}
	}

	fmt.Printf("\n%v", mat)

	var total uint64 = 0
	fmt.Printf("\n%v", operators)

	for i, x := range operators {
		start := x.start
		end := 0
		// if there is no next we get end as end of mat
		if i == len(operators)-1 {
			end = maxL - 1
			fmt.Printf("\nend  = maxL %v", end)
		} else {
			// we get the next operator index - 1 as end
			end = operators[i+1].start - 1
			fmt.Printf("\nend  = next -1 %v", end)
		}

		total += solveRange(mat, x, start, end)

	}

	fmt.Printf("\nSolution is: %v\n", total)

}

func solveRange(mat [][]string, x operator, start, end int) uint64 {
	fmt.Printf("\nRange %v-%v", start, end)
	h := len(mat)
	var total uint64 = 0
	if x.operator == "*" {
		total = 1
	}

	for col := end; col >= start; col-- {
		current := ""
		for row := range h {
			c := mat[row][col]

			if c != "" && c != " " {
				fmt.Printf("\nmat[%v][%v]= %v ", row, col, c)
				current += string(c)
				fmt.Printf("Reading  %v", current)
			}
		}

		fmt.Printf("\nCurrent is   %v", current)
		if current == "" {
			continue
		}

		n, err := strconv.ParseUint(current, 10, 64)
		if err != nil {
			log.Panicf("Error parsing %v", current)
		}

		total = operationsMap[x.operator](total, n)
		fmt.Printf("\ntotal %v\n", total)

	}
	return total

}

func solution1() {
	f, err := os.Open("input-sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var nums = [][]uint64{}
	var operators = []string{}

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		op := strings.TrimSpace(split[0])

		if op == "+" || op == "*" {
			for _, w := range split {
				text := strings.TrimSpace(w)
				if text != "" {
					operators = append(operators, text)
				}
			}
		} else {
			numsLine := []uint64{}
			for _, w := range split {
				w = strings.TrimSpace(w)
				if w != "" {
					n, err := strconv.ParseUint(w, 10, 64)
					if err != nil {
						log.Panicf("Error parsing %v", w)
					}
					numsLine = append(numsLine, n)

				}
			}

			nums = append(nums, numsLine)
			fmt.Printf("nums= %v", nums)
		}
		fmt.Printf("%v", split)
	}

	var total uint64 = 0
	h := len(nums)
	w := len(nums[0])

	for col := range w {
		sign := operators[col]
		opFunc := operationsMap[sign]
		var col_total uint64 = 0
		if sign == "*" {
			col_total = 1
		}
		for row := range h {
			n := nums[row][col]
			fmt.Println(n)
			col_total = opFunc(n, col_total)
		}

		total += col_total
		fmt.Printf("\ntotal= %v\n", total)
	}
	fmt.Printf("\nSolution is: %v\n", total)

}
