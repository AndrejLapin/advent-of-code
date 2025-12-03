package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, err := os.Open("input-3.txt")
	// file, err := os.Open("test-input-3.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	joltage_sum := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		digitArray := []byte{}
		digit_count := 12
		highestIndex := 0
		for digit := 1; digit <= digit_count; digit++ {
			highestCurrent := highestIndex
			// fmt.Printf("Checking for digit %d\n", digit)
			// fmt.Printf("Starting from %d\n", highestCurrent+1)
			for index := highestCurrent + 1; index+(digit_count-digit) < len(line); index++ {
				if line[index] > line[highestCurrent] {
					highestCurrent = index
				}
			}
			// fmt.Printf("Highest current found %d\n", highestCurrent)
			digitArray = append(digitArray, line[highestCurrent])
			highestIndex = highestCurrent + 1
		}
		found_joltage, err := strconv.Atoi(string(digitArray))
		if err != nil {
			panic(err)
		}
		// fmt.Println(found_joltage)
		joltage_sum += found_joltage
	}
	fmt.Println(joltage_sum)
}
