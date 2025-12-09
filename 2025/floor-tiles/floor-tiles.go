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

type Tile struct {
	Position Coord
	Index    int
}

type HorizontalWall struct {
	Y, fromX, toX int
}

type VerticalWall struct {
	X, fromY, toY int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// if they touch == still bad
// horizontal wall Y = 2
// vertical wall from Y:2 to Y:4
// then if horizontalWall.Y >= vertical.FromY && horizontalWall.Y <= vertical.ToY
// then if horizontalWall.FromX < vertical.X && horizontalWall.ToX > vertical.X

// verticalWall X = 2, FromY:3 ToY:7
// horizontal wall from X:2 to X:8, Y:5
// if verticalWall.X >= horizontal.FromX && verticalWall.X <= horizontal.ToX
// then if verticalWall.FromY < horizontal.Y && verticalWall.ToY > horizontal.Y

func wallsCrossing(vWall *VerticalWall, hWall *HorizontalWall) bool {
	return vWall.X > hWall.fromX && vWall.X < hWall.toX && hWall.Y > vWall.fromY && hWall.Y < vWall.toY
}

func horizontalCorssing(hWall *HorizontalWall, vWall *VerticalWall) bool {
	return hWall.Y >= vWall.fromY && hWall.Y <= vWall.toY && hWall.fromX < vWall.X && hWall.toX > vWall.X
}

func verticalWallCorssing(vWall *VerticalWall, hWall *HorizontalWall) bool {
	return vWall.X >= hWall.fromX && vWall.X <= hWall.toX && vWall.fromY < hWall.Y && vWall.toY > hWall.Y
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

	tiles := []Tile{}
	horizontalWalls := []HorizontalWall{}
	verticalWalls := []VerticalWall{}
	MaxX := 0
	MaxY := 0
	redTileMap := map[Coord]int{}

	currentTileIndex := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ",")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		if x+1 > MaxX {
			MaxX = x + 1
		}
		if y+1 > MaxY {
			MaxY = y + 1
		}
		tiles = append(tiles, Tile{Position: Coord{X: x, Y: y}, Index: currentTileIndex})
		redTileMap[Coord{X: x, Y: y}] = currentTileIndex
		currentTileIndex++
	}

	for index := range tiles {
		pointA := tiles[index]
		pointB := tiles[(index+1)%len(tiles)]
		if pointA.Position.X == pointB.Position.X {
			start := min(pointA.Position.Y, pointB.Position.Y)
			end := max(pointA.Position.Y, pointB.Position.Y)
			verticalWalls = append(verticalWalls, VerticalWall{X: pointA.Position.X, fromY: start, toY: end})
		} else {
			start := min(pointA.Position.X, pointB.Position.X)
			end := max(pointA.Position.X, pointB.Position.X)
			horizontalWalls = append(horizontalWalls, HorizontalWall{Y: pointA.Position.Y, fromX: start, toX: end})
		}
	}

	biggest := 0

	for outerIndex := 0; outerIndex < len(tiles)-1; outerIndex++ {
		for innerIndex := outerIndex + 1; innerIndex < len(tiles); innerIndex++ {
			width := abs(tiles[outerIndex].Position.X-tiles[innerIndex].Position.X) + 1
			height := abs(tiles[outerIndex].Position.Y-tiles[innerIndex].Position.Y) + 1
			plot := width * height
			if plot <= biggest {
				continue
			}

			fromX := min(tiles[outerIndex].Position.X, tiles[innerIndex].Position.X)
			toX := max(tiles[outerIndex].Position.X, tiles[innerIndex].Position.X)

			fromY := min(tiles[outerIndex].Position.Y, tiles[innerIndex].Position.Y)
			toY := max(tiles[outerIndex].Position.Y, tiles[innerIndex].Position.Y)

			bottomWall := HorizontalWall{Y: fromY, fromX: fromX, toX: toX}
			topWall := HorizontalWall{Y: toY, fromX: fromX, toX: toX}

			// need to find if any walls are crossing other walls
			horizontalWallsCrossing := false
			for verticalIndex := range verticalWalls {
				if horizontalCorssing(&bottomWall, &verticalWalls[verticalIndex]) ||
					horizontalCorssing(&topWall, &verticalWalls[verticalIndex]) {
					horizontalWallsCrossing = true
					break
				}
			}
			if horizontalWallsCrossing {
				continue
			}

			rightWall := VerticalWall{X: fromX, fromY: fromY, toY: toY}
			leftWall := VerticalWall{X: toX, fromY: fromY, toY: toY}

			verticalWallsCrossing := false
			for horizontalIndex := range horizontalWalls {
				if verticalWallCorssing(&rightWall, &horizontalWalls[horizontalIndex]) ||
					verticalWallCorssing(&leftWall, &horizontalWalls[horizontalIndex]) {
					verticalWallsCrossing = true
					break
				}
			}
			if verticalWallsCrossing {
				continue
			}

			// fmt.Printf("Wall corners: %+v; %+v\n", tiles[outerIndex].Position, tiles[innerIndex].Position)
			biggest = plot
		}
	}
	fmt.Println(biggest)
}
