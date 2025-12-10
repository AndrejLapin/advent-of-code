package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// nicer error output would be better
func Between(data, start, end string) (string, bool) {
	startPos := strings.Index(data, start)
	if startPos == -1 {
		return "", false
	}
	result := data[startPos+1:]
	endPos := strings.Index(result, end)
	if endPos == -1 {
		return "", false
	}
	return result[:endPos], true
}

func Reset[T any](slice []T) {
	for i := range slice {
		var zero T
		slice[i] = zero
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

type SearchState struct {
	Indicator     []bool
	Score         int
	BestNextScore int
	Depth         int
}

func main() {
	fileName := flag.String("input", "input.txt", "input file")
	flag.Parse()

	file, err := os.Open(*fileName)
	// file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	indicatorTargets := [][]bool{}
	allButtons := [][][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		indicators, ok := Between(line, "[", "]")
		if !ok {
			panic("Could not find start or end character")
		}
		indicatorsTarget := make([]bool, len(indicators))
		for index, char := range indicators {
			if char == '#' {
				indicatorsTarget[index] = true
			}
		}
		indicatorTargets = append(indicatorTargets, indicatorsTarget)

		buttonString, ok := Between(line, "]", "{")
		if !ok {
			panic("Could not find start or end character for buttons")
		}
		buttonsParts := strings.Fields(buttonString)

		buttonArray := make([][]int, len(buttonsParts))

		for buttonIndex, buttonsPart := range buttonsParts {
			buttonSubstring, ok := Between(buttonsPart, "(", ")")
			if !ok {
				panic("Could not find start or end character for button substring")
			}
			indicatorsStr := strings.Split(buttonSubstring, ",")
			buttonIndeciesSlice := make([]int, len(indicatorsStr))

			for indicatorIndex, indicatorIndexStr := range indicatorsStr {
				buttonIndexNum, _ := strconv.Atoi(indicatorIndexStr)
				buttonIndeciesSlice[indicatorIndex] = buttonIndexNum
				// buttonArray[index] = append(buttonArray[index], buttonIndexNum)
			}
			buttonArray[buttonIndex] = buttonIndeciesSlice
		}
		allButtons = append(allButtons, buttonArray)
	}

	fewestMoveSum := 0
	for machineIndex := 0; machineIndex < len(indicatorTargets); machineIndex++ {
		targetScore := 0
		for _, indicatorState := range indicatorTargets[machineIndex] {
			if indicatorState {
				targetScore++
			}
		}

		// indicatorOfTheState := indicatorTargets[machineIndex]
		currentButtons := allButtons[machineIndex]
		indicatorsLength := len(indicatorTargets[machineIndex])

		firstSearchState := SearchState{Indicator: indicatorTargets[machineIndex]}
		searchStates := []SearchState{}
		searchStates = append(searchStates, firstSearchState)

		indicatorPerButton := make([]bool, indicatorsLength)
		indicatorPerSecondBest := make([]bool, indicatorsLength)

		// now lets calc each button
		bestMoveFound := false
		for !bestMoveFound {
			if len(searchStates) < 1 {
				panic("We some how ran out of search states!")
			}
			currentState := searchStates[0]
			//create and array of numbers from the target
			// to easier compare arrays, I don't think I need? Because we can just index?

			newSearchStates := []SearchState{}
			for buttonIndex, button := range currentButtons {
				copy(indicatorPerButton, currentState.Indicator) // probably copied from the state
				buttonPressScore := currentState.Score
				targetFound := false
				for _, indicatorIndex := range button {
					if indicatorPerButton[indicatorIndex] {
						targetFound = true
						buttonPressScore++
					} else {
						buttonPressScore--
					}
					indicatorPerButton[indicatorIndex] = !indicatorPerButton[indicatorIndex]
				}

				if buttonPressScore == targetScore {
					for _, indicator := range indicatorPerButton {
						if indicator {
							// fmt.Println("Messed up with this")
							fmt.Printf("Bout to panic but depth is %d\n", currentState.Depth+1)
							panic("All indicators should be false!")
						}
					}
					bestMoveFound = true
					fmt.Printf("For machine %d, found best move at depth - %d\n", machineIndex, currentState.Depth+1)
					fewestMoveSum += currentState.Depth + 1
					break
				}

				if !targetFound {
					// no reason to press this button if no targets were found
					continue
				}

				// calc best next
				bestNextScore := math.MinInt
				for otherButtonIndex, otherButton := range currentButtons {
					innerButtonScore := 0
					if buttonIndex == otherButtonIndex {
						continue
					}
					copy(indicatorPerSecondBest, indicatorPerButton)
					for _, indicatorIndex := range otherButton {
						if indicatorPerSecondBest[indicatorIndex] {
							innerButtonScore++
						} else {
							innerButtonScore--
						}
					}
					bestNextScore = Max(innerButtonScore, bestNextScore)
				}

				indicatorInstance := make([]bool, indicatorsLength)
				copy(indicatorInstance, indicatorPerButton)
				newSearchState := SearchState{Indicator: indicatorInstance,
					Score: buttonPressScore, BestNextScore: bestNextScore, Depth: currentState.Depth + 1}
				newSearchStates = append(newSearchStates, newSearchState)
			}

			sort.Slice(newSearchStates, func(i, j int) bool {
				return newSearchStates[i].Score+newSearchStates[i].BestNextScore >
					newSearchStates[j].Score+newSearchStates[j].BestNextScore
			})

			searchStates = append(searchStates, newSearchStates...)
			searchStates = searchStates[1:]
		}
		// for searchStateIndex := range searchStates {
		// 	fmt.Printf("%+v, total score %d\n", searchStates[searchStateIndex], searchStates[searchStateIndex].Score+searchStates[searchStateIndex].BestNextScore)
		// }
		// fmt.Println("=============================================")
	}

	fmt.Println(fewestMoveSum)

	// for machineIndex := 0; machineIndex < len(indicatorTargets); machineIndex++ {
	// 	fmt.Printf("Indicator targets - %v; Buttons - %v\n", indicatorTargets[machineIndex], allButtons[machineIndex])
	// }
}
