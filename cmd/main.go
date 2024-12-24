package main

// The order is bad IFF the rules contains a value for a pair in the order and the order is reversed.

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	proceeding string
	following  string
}

func reorderLine(lineNumber int, rules []rule, line []string) (bool, []string) {
	for i, proceeding := range line {
		for j, following := range line[i:] {
			for _, rule := range rules {
				if rule.proceeding == following && rule.following == proceeding {
					line[i] = following
					line[i+j] = proceeding
					return false, line
				}
			}
		}
	}
	return true, line
}

func checkLine(lineNumber int, rules []rule, line string) int {
	strs := strings.Split(line, ",")
	isValid, strs := reorderLine(lineNumber, rules, strs)
	if isValid {
		return 0
	}
	for !isValid {
		isValid, strs = reorderLine(lineNumber, rules, strs)
	}
	middleIndex := len(strs) / 2
	middle := strs[middleIndex]
	x, err := strconv.ParseInt(middle, 10, 64)
	if err != nil {
		log.Fatalf("Unable to parse center character %s on line %d", err, lineNumber)
	}
	return int(x)
}

func main() {
	filepath := "input.txt"
	// filepath := "test_input.txt"
	file, err := os.Open(filepath)
	if err != nil {
		log.Panic("Unable to open file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0
	rules := []rule{}
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		rulearray := strings.Split(line, "|")
		rule := rule{
			proceeding: rulearray[0],
			following:  rulearray[1],
		}
		rules = append(rules, rule)
	}

	for scanner.Scan() {
		line := scanner.Text()
		value := checkLine(lineNumber, rules, line)
		sum += value
		lineNumber++
	}
	log.Printf("Final sum: %d", sum)
}
