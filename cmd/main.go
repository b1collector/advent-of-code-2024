package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	x    = 'X'
	m    = 'M'
	a    = 'A'
	s    = 'S'
	size = 140
)

func main() {
	filePath := "input.txt"
	file, err := os.Open(filePath)
	if err != nil {
		panic("Error opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	searchArea := [][]rune{}
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != size {
			log.Fatalf("Error on line %d. Line has length %d", lineNumber, len(line))
			panic("Unrecoverable error")
		}
		searchArea = append(searchArea, []rune(line))
	}
	if len(searchArea) != size {
		log.Fatalf("Error on line %d. Line has length %d", lineNumber, len(searchArea))
		panic("Unrecoverable error")
	}

	count := 0
	for i := 1; i < size-1; i++ {
		for j := 1; j < size-1; j++ {
			// skip all the characters that do not match the first character in the
			// search string
			if searchArea[i][j] != a {
				continue
			}

			if isXMatch(searchArea, i, j) {
				count++
			}

		}
	}
	fmt.Printf("Count: %d\n", count)
}

func isXMatch(searchArea [][]rune, i int, j int) bool {
	return searchArea[i][j] == a &&
		((searchArea[i-1][j-1] == m && searchArea[i+1][j+1] == s) || (searchArea[i-1][j-1] == s && searchArea[i+1][j+1] == m)) &&
		((searchArea[i-1][j+1] == m && searchArea[i+1][j-1] == s) || (searchArea[i-1][j+1] == s && searchArea[i+1][j-1] == m))
}
