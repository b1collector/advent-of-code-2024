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
	for i := range size {
		for j := range size {
			// skip all the characters that do not match the first character in the
			// search string
			if searchArea[i][j] != x {
				continue
			}

			if isNorthMatch(searchArea, i, j) {
				count++
			}

			if isSouthMatch(searchArea, i, j) {
				count++
			}

			if isWestMatch(searchArea, i, j) {
				count++
			}

			if isEastMatch(searchArea, i, j) {
				count++
			}
			if isNorthEastMatch(searchArea, i, j) {
				count++
			}
			if isSouthEastMatch(searchArea, i, j) {
				count++
			}
			if isSouthWestMatch(searchArea, i, j) {
				count++
			}
			if isNorthWestMatch(searchArea, i, j) {
				count++
			}
		}
	}
	fmt.Printf("Count: %d\n", count)
}

func isNorthMatch(searchArea [][]rune, i int, j int) bool {
	if i-3 < 0 {
		return false
	}
	return searchArea[i][j] == x &&
		searchArea[i-1][j] == m &&
		searchArea[i-2][j] == a &&
		searchArea[i-3][j] == s
}

func isSouthMatch(searchArea [][]rune, i int, j int) bool {
	if i+3 >= size {
		return false
	}
	return searchArea[i][j] == x &&
		searchArea[i+1][j] == m &&
		searchArea[i+2][j] == a &&
		searchArea[i+3][j] == s
}

func isWestMatch(searchArea [][]rune, i int, j int) bool {
	if j-3 < 0 {
		return false
	}
	return searchArea[i][j] == x &&
		searchArea[i][j-1] == m &&
		searchArea[i][j-2] == a &&
		searchArea[i][j-3] == s
}

func isEastMatch(searchArea [][]rune, i int, j int) bool {
	if j+3 >= size {
		return false
	}
	return searchArea[i][j] == x &&
		searchArea[i][j+1] == m &&
		searchArea[i][j+2] == a &&
		searchArea[i][j+3] == s
}

func isNorthEastMatch(searchArea [][]rune, i int, j int) bool {
	if j+3 >= size || i-3 < 0 {
		return false
	}
	return searchArea[i][j] == x &&
		searchArea[i-1][j+1] == m &&
		searchArea[i-2][j+2] == a &&
		searchArea[i-3][j+3] == s
}

func isSouthEastMatch(searchArea [][]rune, i int, j int) bool {
	if j+3 >= size || i+3 >= size {
		return false
	}
	return searchArea[i][j] == x &&
		searchArea[i+1][j+1] == m &&
		searchArea[i+2][j+2] == a &&
		searchArea[i+3][j+3] == s
}

func isNorthWestMatch(searchArea [][]rune, i int, j int) bool {
	if j-3 < 0 || i-3 < 0 {
		return false
	}
	return searchArea[i][j] == x &&
		searchArea[i-1][j-1] == m &&
		searchArea[i-2][j-2] == a &&
		searchArea[i-3][j-3] == s
}

func isSouthWestMatch(searchArea [][]rune, i int, j int) bool {
	if j-3 < 0 || i+3 >= size {
		return false
	}
	return searchArea[i][j] == x &&
		searchArea[i+1][j-1] == m &&
		searchArea[i+2][j-2] == a &&
		searchArea[i+3][j-3] == s
}
