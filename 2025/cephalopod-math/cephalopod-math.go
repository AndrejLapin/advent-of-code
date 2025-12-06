package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileName := flag.String("input", "input.txt", "input file")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	numbers := [][]int{}
	operations := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if line[0] == '*' || line[0] == '+' {
			operations = parts
		} else {
			// []int{}
			numberRow := make([]int, len(parts))
			for index, part := range parts {
				number, _ := strconv.Atoi(part)
				numberRow[index] = number
			}
			numbers = append(numbers, numberRow)
		}
	}

	// fmt.Println(operations)
	// fmt.Println(len(operations))
	total := 0
	for columnIndex, operator := range operations {
		operationResult := 0
		if operator == "*" {
			operationResult = 1
		}
		for _, row := range numbers {
			if operator == "*" {
				operationResult *= row[columnIndex]
			} else {
				operationResult += row[columnIndex]
			}
		}
		total += operationResult
	}
	fmt.Println(total)
}
