package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func findNeighboursOnRow(warehouse [][]byte, row int, column int) int {
	neighboughrRollCount := 0
	leftColumn := column - 1
	rightColumn := column + 1
	if leftColumn >= 0 {
		if warehouse[row][leftColumn] == '@' {
			neighboughrRollCount++
		}
	}
	if warehouse[row][column] == '@' {
		neighboughrRollCount++
	}
	if rightColumn < len(warehouse[row]) {
		if warehouse[row][rightColumn] == '@' {
			neighboughrRollCount++
		}
	}
	return neighboughrRollCount
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

	warehouse := [][]byte{}
	for scanner.Scan() {
		line := append([]byte(nil), scanner.Bytes()...)
		warehouse = append(warehouse, line)
	}

	// accessibleRollCoordinates := [][2]uint{} // for debug

	// less than 4 @
	wasRollRemoved := true
	removedCount := 0
	// accessibleRollCount := 0
	for wasRollRemoved {
		wasRollRemoved = false
		for rowIndex, row := range warehouse {
			for columnIndex, element := range row {
				if element != '@' {
					continue
				}
				neighboughrRollCount := 0
				topRow := rowIndex - 1
				bottomRow := rowIndex + 1
				if topRow >= 0 {
					neighboughrRollCount += findNeighboursOnRow(warehouse, topRow, columnIndex)
				}
				neighboughrRollCount += findNeighboursOnRow(warehouse, rowIndex, columnIndex)
				if bottomRow < len(warehouse) {
					neighboughrRollCount += findNeighboursOnRow(warehouse, bottomRow, columnIndex)
				}
				if rowIndex == len(warehouse)-1 {
				}
				if neighboughrRollCount <= 4 {
					wasRollRemoved = true
					removedCount++
					warehouse[rowIndex][columnIndex] = '.'
					// accessibleRollCount++
					// accessibleRollCoordinates = append(accessibleRollCoordinates, [2]uint{uint(rowIndex), uint(columnIndex)})
				}
			}
		}
	}

	// for _, coords := range accessibleRollCoordinates {
	// 	warehouse[coords[0]][coords[1]] = 'X'
	// }
	for _, row := range warehouse {
		fmt.Println(string(row))
	}
	// fmt.Printf("Warehouse length - %d\n", len(warehouse))
	fmt.Println(removedCount)
}
