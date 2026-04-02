package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "strconv"
	"strings"
)

func main() {
	solution1()
	// solution2()
}

func solution2() {

	// test
}

func solution1() {
	f, err := os.Open("input-sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		fmt.Println(split)
	}

}
