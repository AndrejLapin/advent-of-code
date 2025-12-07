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

	// easy with char positions
	// with S we save the initial index
	// then we add more indecies
	scanner.Scan()
	firstLine := scanner.Text()
	beamLine := make([]int, len(firstLine))
	nextBeamLine := make([]int, len(firstLine))
	beamLine[strings.Index(firstLine, "S")] = 1

	for scanner.Scan() {
		line := scanner.Text()

		reducedString := line
		splitterPos := strings.Index(reducedString, "^")
		actualSplitterPos := splitterPos
		// fmt.Println(beamLine)
		// fmt.Println(line)
		copy(nextBeamLine, beamLine)

		for splitterPos != -1 {
			reducedString = reducedString[splitterPos+1:]
			// fmt.Println(reducedString)
			if splitterPos != -1 && beamLine[actualSplitterPos] != 0 {
				nextBeamLine[actualSplitterPos-1] += beamLine[actualSplitterPos]
				nextBeamLine[actualSplitterPos] = 0
				nextBeamLine[actualSplitterPos+1] += beamLine[actualSplitterPos]
			}
			splitterPos = strings.Index(reducedString, "^")
			// fmt.Printf("Splitter pos: %d\n", splitterPos)
			actualSplitterPos += splitterPos + 1
			// fmt.Printf("Actual splitter pos: %d\n", actualSplitterPos)
		}
		copy(beamLine, nextBeamLine)
	}
	timeLineCount := 0
	for _, beamCount := range beamLine {
		timeLineCount += beamCount
	}
	fmt.Println(timeLineCount)
}
