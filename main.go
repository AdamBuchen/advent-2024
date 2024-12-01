package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var pattern *regexp.Regexp = regexp.MustCompile(`^(\d+)(\s+)(\d+)`)

func main() {
	// Open input.txt, for each line, parse out two ints
	file, fileErr := os.Open("input.txt")
	if fileErr != nil {
		fmt.Println(fileErr)
		return
	}

	contents, readErr := io.ReadAll(file)
	if readErr != nil {
		fmt.Println(readErr)
		return
	}

	contentsStr := string(contents)
	splitByRow := strings.Split(contentsStr, "\n")

	leftSlice := make([]int, 0)
	rightSlice := make([]int, 0)

	var totalDiff int64 = 0

	//totalDifference := 0
	for _, row := range splitByRow {
		matches := pattern.FindStringSubmatch(row)
		if matches == nil {
			panic("No matches found")
		}

		leftStr := strings.TrimSpace(matches[1])
		rightStr := strings.TrimSpace(matches[3])

		leftInt, leftErr := strconv.Atoi(leftStr)
		if leftErr != nil {
			panic(leftErr)
		}

		rightInt, rightErr := strconv.Atoi(rightStr)
		if rightErr != nil {
			panic(rightErr)
		}

		leftSlice = append(leftSlice, leftInt)
		rightSlice = append(rightSlice, rightInt)
	}

	// sort leftSlice and rightSlice
	sort.Ints(leftSlice)
	sort.Ints(rightSlice)

	for i, left := range leftSlice {
		right := rightSlice[i]
		diff := math.Abs(float64(left - right))
		diffInt64 := int64(diff)
		totalDiff += diffInt64
	}

	fmt.Println(totalDiff)
}
