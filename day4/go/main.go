package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// Directions to access each neighbour
var DIR = [8][2]int{
	{-1, 0},  // TOP
	{1, 0},   // BOTTOM
	{0, -1},  // LEFT
	{-1, -1}, // LEFT UP
	{1, -1},  // LEFT DOWN
	{0, 1},   // RIGHT
	{1, 1},   // RIGHT DOWN
	{-1, 1},  // RIGHT UP
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	var mat [][]string = make([][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		mat = append(mat, strings.Split(line, ""))
	}

	fmt.Printf("Solution is %v", solution1(mat))
	fmt.Printf("\nSolution2 is %v", solution2(&mat))

}

func inBounds(mat [][]string, y int, x int) bool {
	w := len(mat[0])
	h := len(mat)
	if (y >= 0 && y < h) && (x >= 0 && x < w) {
		return true
	}

	return false
}

func safeGet(mat [][]string, y int, x int) (string, error) {
	if inBounds(mat, y, x) {
		return mat[y][x], nil
	}

	return "", errors.New("Out of bounds")
}

func countAdjacents(mat [][]string, y, x int) int {
	count := 0
	for _, d := range DIR {
		adj_y := y + d[0]
		adj_x := x + d[1]
		adj, err := safeGet(mat, adj_y, adj_x)
		if err == nil {
			if adj == "@" || adj == "M" {
				count++
			}
		}
	}

	return count

}
func pass(mat *[][]string) int {
	count := 0
	for y, row := range *mat {
		for x, item := range row {
			if item == "@" || item == "M" {
				if countAdjacents(*mat, y, x) < 4 {
					(*mat)[y][x] = "M"
					count++

				}
			}
		}
	}

	for y, row := range *mat {
		for x, item := range row {
			if item == "M" {
				(*mat)[y][x] = "."

			}
		}
	}

	return count
}

func solution2(mat *[][]string) int {
	count := pass(mat)
	total := count

	for count != 0 {
		count = pass(mat)
		total += count
	}

	return total
}

func solution1(mat [][]string) int {
	count := 0
	for y, row := range mat {
		for x, item := range row {
			if item == "@" {
				if countAdjacents(mat, y, x) < 4 {
					count++
				} else {
				}
			}
		}
	}

	return count
}
