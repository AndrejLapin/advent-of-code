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

	hitSplitterCount := 0

	// easy with char positions
	// with S we save the initial index
	// then we add more indecies
	scanner.Scan()
	firstLine := scanner.Text()
	beamLine := make([]bool, len(firstLine))
	nextBeamLine := make([]bool, len(firstLine))
	beamLine[strings.Index(firstLine, "S")] = true

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
			if splitterPos != -1 && beamLine[actualSplitterPos] {
				nextBeamLine[actualSplitterPos-1] = true
				nextBeamLine[actualSplitterPos] = false
				nextBeamLine[actualSplitterPos+1] = true
				hitSplitterCount++
			}
			splitterPos = strings.Index(reducedString, "^")
			// fmt.Printf("Splitter pos: %d\n", splitterPos)
			actualSplitterPos += splitterPos + 1
			// fmt.Printf("Actual splitter pos: %d\n", actualSplitterPos)
		}
		copy(beamLine, nextBeamLine)
	}
	fmt.Println(hitSplitterCount)
}
