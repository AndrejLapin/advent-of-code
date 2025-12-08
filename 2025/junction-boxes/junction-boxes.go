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

type Connection struct {
	Distance float64
	PointA   *JunctionBox
	PointB   *JunctionBox
}

type JunctionBox struct {
	X, Y, Z       int
	ParentCircuit *Circuit
}

type Circuit struct {
	JunctionBoxes []*JunctionBox
}

func (j *JunctionBox) InitCircuit(circuit *Circuit) {
	*circuit = Circuit{JunctionBoxes: []*JunctionBox{j}}
	j.ParentCircuit = circuit
}

func main() {
	fileName := flag.String("input", "input.txt", "input file")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	junctionBoxes := []JunctionBox{}

	for scanner.Scan() {
		line := scanner.Text()
		coords := strings.Split(line, ",")

		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])
		junctionBox := JunctionBox{X: x, Y: y, Z: z}
		junctionBoxes = append(junctionBoxes, junctionBox)
	}

	circuits := make([]Circuit, len(junctionBoxes))

	for currentBox, _ := range junctionBoxes {
		junctionBoxes[currentBox].InitCircuit(&circuits[currentBox])
	}

	connections := make([]Connection, len(junctionBoxes)*(len(junctionBoxes)-1)/2)

	newPathIndex := 0
	for outerBox := 0; outerBox < len(junctionBoxes)-1; outerBox++ {
		for innerBox := outerBox + 1; innerBox < len(junctionBoxes); innerBox++ {
			xPow := math.Pow(float64(junctionBoxes[innerBox].X-junctionBoxes[outerBox].X), 2)
			yPow := math.Pow(float64(junctionBoxes[innerBox].Y-junctionBoxes[outerBox].Y), 2)
			zPow := math.Pow(float64(junctionBoxes[innerBox].Z-junctionBoxes[outerBox].Z), 2)
			distance := math.Sqrt(xPow + yPow + zPow)
			connections[newPathIndex] =
				Connection{Distance: distance, PointA: &junctionBoxes[outerBox], PointB: &junctionBoxes[innerBox]}
			newPathIndex++
		}
	}
	fmt.Printf("\n")

	sort.Slice(connections, func(i, j int) bool {
		return connections[i].Distance < connections[j].Distance
	})

	connectionIndex := 0
	connectedAmount := 0
	for true {
		// fmt.Printf("================== %d ===================\n", connectedAmount)
		pointA := connections[connectionIndex].PointA
		pointB := connections[connectionIndex].PointB
		if pointA.ParentCircuit == pointB.ParentCircuit {
			connectedAmount++
			connectionIndex++
			continue
		}
		// otherwise merge
		indexOfCurrentCircuit := -1
		for circuitIndex, _ := range circuits {
			if &circuits[circuitIndex] == pointA.ParentCircuit {
				indexOfCurrentCircuit = circuitIndex
			}
		}

		if indexOfCurrentCircuit == -1 {
			panic("Something has gone terribly wrong with trying to find current circuit!")
		}

		pointB.ParentCircuit.JunctionBoxes = append(pointB.ParentCircuit.JunctionBoxes,
			pointA.ParentCircuit.JunctionBoxes...)

		if len(pointB.ParentCircuit.JunctionBoxes) == len(junctionBoxes) {
			fmt.Printf("Point A: %+v; Point B: %+v\n", *pointA, *pointB)
			fmt.Printf("X coords multiplied - %d\n", pointA.X*pointB.X)
			break
		}

		for _, junctionBoxRef := range pointA.ParentCircuit.JunctionBoxes {
			junctionBoxRef.ParentCircuit = pointB.ParentCircuit
		}

		circuits[indexOfCurrentCircuit] = Circuit{}
		connectedAmount++
		connectionIndex++
	}
}
