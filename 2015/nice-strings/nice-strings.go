package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
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

	const minVowelCount int = 3
	const targetVowels string = "aeiou"
	bannedSubstrings := [...]string{"ab", "cd", "pq", "xy"}

	niceStrings := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		// fmt.Printf("Checking string %s\n", line)
		hasBannedSubstring := false
		for _, substring := range bannedSubstrings {
			if strings.Contains(line, substring) {
				hasBannedSubstring = true
				break
			}
		}
		if hasBannedSubstring {
			continue
		}
		// fmt.Printf("Did not contain banned substrings\n")

		vowelCount := 0
		duplicatesInARowFound := false
		var previous rune = 0
		for _, char := range line {
			if !duplicatesInARowFound && char == previous {
				duplicatesInARowFound = true
			}
			if strings.ContainsRune(targetVowels, char) {
				vowelCount++
			}
			previous = char
		}

		if vowelCount >= minVowelCount && duplicatesInARowFound {
			niceStrings++
			// fmt.Println(line)
		}
	}
	fmt.Println(niceStrings)
}
