package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	X, Y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

	tiles := []Coord{}

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		tiles = append(tiles, Coord{X: x, Y: y})
	}
	biggest := 0

	for outerIndex := 0; outerIndex < len(tiles)-1; outerIndex++ {
		for innerIndex := outerIndex + 1; innerIndex < len(tiles); innerIndex++ {
			width := abs(tiles[outerIndex].X-tiles[innerIndex].X) + 1
			height := abs(tiles[outerIndex].Y-tiles[innerIndex].Y) + 1
			plot := width * height
			if plot > biggest {
				biggest = plot
			}
		}
	}
	fmt.Println(biggest)
}
