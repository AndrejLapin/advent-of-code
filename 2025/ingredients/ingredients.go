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

	ranges := [][2]int{}

	scanningRanges := true
	for scanner.Scan() {
		if scanningRanges {
			line := scanner.Text()
			if len(line) == 0 {
				scanningRanges = false
				break
			}
			parts := strings.Split(line, "-")
			rangeStart, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(err)
			}
			rangeEnd, err := strconv.Atoi(parts[1])
			if err != nil {
				panic(err)
			}

			previousExpandedRange := -1
			currentRange := [2]int{rangeStart, rangeEnd}

			emplaceRange := true
			findCanExpand := true
			loopCount := 10
			for findCanExpand && loopCount > 0 {
				loopCount--
				rangeExpanded := false
				findCanExpand = false
				currentExpandedRange := -1

				for index, singleRange := range ranges {
					if index == previousExpandedRange {
						continue
					}
					if currentRange[0] >= singleRange[0] && currentRange[0] <= singleRange[1] &&
						currentRange[1] > singleRange[1] {
						ranges[index][1] = currentRange[1]
						currentExpandedRange = index
						rangeExpanded = true
						break
					} else if currentRange[1] >= singleRange[0] && currentRange[1] <= singleRange[1] &&
						currentRange[0] < singleRange[0] {
						ranges[index][0] = currentRange[0]
						currentExpandedRange = index
						rangeExpanded = true
						break
					} else if currentRange[0] <= singleRange[0] && currentRange[1] >= singleRange[1] {
						ranges[index][0] = currentRange[0]
						ranges[index][1] = currentRange[1]
						currentExpandedRange = index
						rangeExpanded = true
						break
					} else if currentRange[0] >= singleRange[0] && currentRange[1] <= singleRange[1] {
						emplaceRange = false
						break
					}
				}

				if rangeExpanded {
					emplaceRange = false
					findCanExpand = true
					currentRange[0] = ranges[currentExpandedRange][0]
					currentRange[1] = ranges[currentExpandedRange][1]
					if previousExpandedRange != currentExpandedRange && previousExpandedRange != -1 {
						ranges = append(ranges[:previousExpandedRange], ranges[previousExpandedRange+1:]...)
						if previousExpandedRange < currentExpandedRange {
							currentExpandedRange--
						}
					}
				}
				previousExpandedRange = currentExpandedRange
			}
			if emplaceRange {
				ranges = append(ranges, [2]int{rangeStart, rangeEnd})
			}
		}
	}
	freshIngredientTotal := 0
	for _, singleRange := range ranges {
		freshIngredientTotal += singleRange[1] - singleRange[0] + 1
	}
	fmt.Println(freshIngredientTotal)
}
