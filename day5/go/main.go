package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start  uint64
	end    uint64
	ignore bool
}

func main() {
	solution1()
	solution2()

}

type ByStart []Range

func (a ByStart) Len() int           { return len(a) }
func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByStart) Less(i, j int) bool { return a[i].start < a[j].start }

func solution2() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	ranges := []Range{}

	for scanner.Scan() {
		line := scanner.Text()
		println(line)
		if line == "" {
			break
		}

		split := strings.Split(line, "-")
		s, err := strconv.ParseUint(split[0], 10, 64)

		if err != nil {
			log.Fatal(err)
		}

		e, err := strconv.ParseUint(split[1], 10, 64)

		if err != nil {
			log.Fatal(err)
		}

		r := Range{
			start: s,
			end:   e,
		}

		if s > e {
			fmt.Printf("\n%v", r)
			log.Panic("s is bigger than e")
		}

		ranges = append(ranges, r)

	}

	ranges = merge(ranges)

	var count uint64 = 0
	for _, r := range ranges {
		if !r.ignore {
			count += (r.end - r.start) + 1
		}

	}

	fmt.Printf("Solution is: %v", count)

}

func merge(ranges []Range) []Range {

	sort.Sort(ByStart(ranges))
	merged := []Range{ranges[0]}

	for _, r := range ranges {
		last := &merged[len(merged)-1]

		if r.start <= last.end+1 {
			last.end = max(last.end, r.end)
		} else {
			merged = append(merged, r)
		}

	}

	return merged
}

func overlap(a, b Range) bool {
	return a.end >= b.start && b.end >= a.start
}

func solution1() {

	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	readingRanges := true
	ranges := []Range{}
	nums := []uint64{}

	for scanner.Scan() {
		line := scanner.Text()
		println(line)

		if line == "" {
			readingRanges = false
			continue
		}

		split := strings.Split(line, "-")

		if readingRanges {
			s, err := strconv.ParseUint(split[0], 10, 64)

			if err != nil {
				log.Fatal(err)
			}

			e, err := strconv.ParseUint(split[1], 10, 64)

			if err != nil {
				log.Fatal(err)
			}

			ranges = append(ranges, Range{
				start: s,
				end:   e,
			})

		} else {
			n, err := strconv.ParseUint(split[0], 10, 64)

			if err != nil {
				log.Fatal(err)
			}

			nums = append(nums, n)
		}

	}

	var count uint64 = 0

	for _, n := range nums {
		for _, r := range ranges {
			if n >= r.start && n <= r.end {
				count++
				break
			}
		}

	}

	fmt.Printf("Solution is: %v", count)

}
