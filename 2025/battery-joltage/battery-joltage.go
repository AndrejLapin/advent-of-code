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

	joltageSum := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		const digitCount = 12
		digitArray := [digitCount]byte{}
		highestIndex := 0
		for digit := 1; digit <= digitCount; digit++ {
			highestCurrent := highestIndex
			// fmt.Printf("Checking for digit %d\n", digit)
			// fmt.Printf("Starting from %d\n", highestCurrent+1)
			for index := highestCurrent + 1; index+(digitCount-digit) < len(line); index++ {
				if line[index] > line[highestCurrent] {
					highestCurrent = index
				}
			}
			// fmt.Printf("Highest current found %d\n", highestCurrent)
			digitArray[digit-1] = line[highestCurrent]
			highestIndex = highestCurrent + 1
		}
		foundJoltage, err := strconv.Atoi(string(digitArray[:]))
		if err != nil {
			panic(err)
		}
		// fmt.Println(found_joltage)
		joltageSum += foundJoltage
	}
	fmt.Println(joltageSum)
}
