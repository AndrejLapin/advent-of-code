package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/aclements/go-z3/z3"
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

	joltageCounters := [][]int64{}
	allButtons := [][][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		countersStr, ok := Between(line, "{", "}")
		if !ok {
			panic("Could not find start or end character")
		}
		counterParts := strings.Split(countersStr, ",")
		indicatorsTarget := make([]int64, len(counterParts))
		for index, counterStr := range counterParts {
			counterNumber, _ := strconv.ParseInt(counterStr, 10, 64)
			indicatorsTarget[index] = counterNumber
		}
		joltageCounters = append(joltageCounters, indicatorsTarget)

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

	var fewestMoveSum int64 = 0
	for machineIndex := 0; machineIndex < len(joltageCounters); machineIndex++ {

		currentButtons := allButtons[machineIndex]
		machineStartingCounter := joltageCounters[machineIndex]

		z3Ctx := z3.NewContext(nil)
		iSort := z3Ctx.IntSort()
		solver := z3.NewSolver(z3Ctx)

		buttonPresses := make([]z3.Int, len(currentButtons))
		for i := range buttonPresses {
			button := z3Ctx.Const(fmt.Sprintf("b%d", i), iSort).(z3.Int)
			buttonPresses[i] = button
			solver.Assert(button.GE(z3Ctx.FromInt(0, iSort).(z3.Int)))
		}

		for i, value := range machineStartingCounter {
			voltageSum := z3Ctx.FromInt(0, iSort).(z3.Int)
			for j, button := range currentButtons {
				if slices.Contains(button, i) {
					voltageSum = voltageSum.Add(buttonPresses[j])
				}
			}
			solver.Assert(voltageSum.Eq(z3Ctx.FromInt(value, iSort).(z3.Int)))
		}

		var n int64 = 0
		for sat, _ := solver.Check(); sat; sat, _ = solver.Check() {
			model := solver.Model()
			n = 0
			pressesSum := z3Ctx.FromInt(0, iSort).(z3.Int)
			for buttonVar := range buttonPresses {
				pressesSum = pressesSum.Add(buttonPresses[buttonVar])
				actualVal, _, _ := model.Eval(buttonPresses[buttonVar], true).(z3.Int).AsInt64()
				n += actualVal
			}
			// fmt.Printf("Solver found solution for machine %d in %d moves\n", machineIndex, n)
			solver.Assert(pressesSum.LT(z3Ctx.FromInt(n, iSort).(z3.Int)))
		}
		// fmt.Printf("Final result for the machine %d is %d\n", machineIndex, n)
		fewestMoveSum += n
	}

	fmt.Println(fewestMoveSum)

	// for machineIndex := 0; machineIndex < len(indicatorTargets); machineIndex++ {
	// 	fmt.Printf("Indicator targets - %v; Buttons - %v\n", indicatorTargets[machineIndex], allButtons[machineIndex])
	// }
}
