package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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
	safeCount := 0

	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Outer loop - line:%d - %s\n", lineNum, line)
		report := strings.Split(line, " ")
		if len(report) < 2 {
			log.Fatalf("Error processing file on line %d. Too few reports to process", lineNum)
			panic("Error loading file.")
		}

		for skip_level := range len(report) {
			levels := []string{}
			levels = append(levels, report[:skip_level]...)
			levels = append(levels, report[skip_level+1:]...)
			fmt.Printf("middle loop - report:%v - skip level:%d\n", levels, skip_level)
			lastJump := 0.
			valid := true
			for i := 1; i < len(levels); i++ {
				previousReport, _ := strconv.Atoi(levels[i-1])
				currentReport, _ := strconv.Atoi(levels[i])

				currentJump := float64(previousReport - currentReport)
				if currentJump == 0 || math.Abs(currentJump) > 3 || lastJump*currentJump < 0 {
					valid = false
				}
				lastJump = currentJump
			}

			fmt.Printf("Level Result %d - %d - %v - %v\n", lineNum, skip_level, levels, valid)
			if valid {
				safeCount++
				break
			}
		}
		lineNum++
	}

	fmt.Printf("The final total is %d\n", safeCount)
}
