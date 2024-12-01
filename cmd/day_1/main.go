package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	filePath := "input.txt"

	file, err := os.Open(filePath)
	if err != nil {
		panic("Error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	leftBuffer := make([]int, 0)
	rightBuffer := make([]int, 0)

	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")
		if len(split) != 2 {
			log.Fatalf("Error processing file on line %d.", lineNum)
			panic("Error loading file.")
		}
		left, err := strconv.Atoi(split[0])
		if err != nil {
			log.Fatalf("Error processing left string on line %d.", lineNum)
			panic("Unable to convert left string to int.")
		}
		right, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatalf("Error processing right string on line %d.", lineNum)
			panic("Unable to convert right string to int.")
		}
		leftBuffer = append(leftBuffer, left)
		rightBuffer = append(rightBuffer, right)
		lineNum++
	}

	log.Printf("Buffer length is %d", len(leftBuffer))
	if len(leftBuffer) != len(rightBuffer) {
		panic("The length of the two arrays do not match.")
	}

	slices.Sort(leftBuffer)
	slices.Sort(rightBuffer)

	sum := 0
	for i := range len(leftBuffer) {
		diff := int(math.Abs(float64(rightBuffer[i] - leftBuffer[i])))
		sum += diff
		log.Printf("Line: %d. Left: %d. Right: %d. Diff: %d. Sum: %d\n", i, leftBuffer[i], rightBuffer[i], diff, sum)
	}

	fmt.Printf("The final total is %d\n", sum)
}
