package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	rexp, err := regexp.Compile(`mul\(\d+,\d+\)|do(n't)?`)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	total := 0
	shouldDo := true
	for scanner.Scan() {
		line := scanner.Text()
		matches := rexp.FindAllString(line, -1)
		for _, match := range matches {
			fmt.Printf("%s\n", match)
			match, found := strings.CutPrefix(match, `mul(`)
			if !found {
				shouldDo = match == "do"
				continue
				// panic("Somehow found a match that wasn't an actual match")
			}
			if !shouldDo {
				continue
			}
			match, found = strings.CutSuffix(match, `)`)
			if !found {
				panic("Somehow missing the trailing paranthesis")
			}
			first, second, found := strings.Cut(match, ",")
			if !found {
				panic("Somehow missing the comma")
			}
			left, err := strconv.Atoi(first)
			if err != nil {
				panic("Unable to parse first integer argument")
			}
			right, err := strconv.Atoi(second)
			if err != nil {
				panic("Unable to parse first integer argument")
			}
			total += left * right
		}
	}

	fmt.Printf("Total: %d\n", total)

}
