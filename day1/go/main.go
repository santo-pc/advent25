package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	current := 50
	zeros := 0

	for scanner.Scan() {
		line := scanner.Text()
		dir := line[0]
		increment, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		newIndex, crossings := Handle(dir, current, increment)
		current = newIndex
		zeros += crossings

	}
	fmt.Printf("\nSolution: %v", zeros)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
func Handle(dir byte, current int, increment int) (int, int) {
	crossings := 0
	sign := 1
	switch dir {
	case 'L':
		sign = -1
		crossings = go_left(current, sign*increment)
	case 'R':
		crossings = go_right(current, increment)
	default:
		fmt.Println("Line does not start with L or R")
		return 0, 0
	}

	newCurrent := modulo(current+(increment*sign), 100)

	fmt.Printf("\n%c\t%v\t#%v->\t%v", dir, increment, newCurrent, crossings)
	return newCurrent, crossings
}

func go_right(current, increment int) int {
	return (current + increment) / 100
}

func go_left(current, increment int) int {
	sum := increment + current
	var c float64 = (float64(sum)) / 100.00

	newCurrent := modulo(sum, 100)
	crossings := Abs(int(math.Floor(c)))

	if current == 0 && increment != 0 {
		crossings--
	}

	if increment != 0 && newCurrent == 0 {
		return crossings + 1
	}

	return crossings

}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func modulo(di, dv int) int {
	return (di%dv + dv) % dv
}
