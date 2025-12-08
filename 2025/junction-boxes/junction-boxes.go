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
	connectionCount := flag.Int("count", 1000, "connection count")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	junctionBoxes := []JunctionBox{}
	// circuits := []Circuit{}

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
		// fmt.Println(junctionBox)
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

	// for index, _ := range connections {
	// 	fmt.Printf("From: %+v; To: %+v; Distance - %d\n", *connections[index].PointA,
	// 		*connections[index].PointB, connections[index].Distance)
	// }

	connectionIndex := 0
	connectedAmount := 0
	for connectedAmount < *connectionCount && connectionIndex < len(connections) {
		// fmt.Printf("================== %d ===================\n", connectedAmount)
		pointA := connections[connectionIndex].PointA
		pointB := connections[connectionIndex].PointB
		// fmt.Printf("Trying to connect %p and %p\n", pointA, pointB)
		// fmt.Printf("A coords: %+v; B coords: %+v\n", pointA, pointB)
		if pointA.ParentCircuit == pointB.ParentCircuit {
			// fmt.Printf("Nodes already connected!\n")
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
			// fmt.Println("Something has gone terribly wrong with trying to find current circuit!")
			panic("Something has gone terribly wrong with trying to find current circuit!")
		}

		// fmt.Printf("PointA parent circuit boxes - %v\n", pointA.ParentCircuit.JunctionBoxes)
		// fmt.Printf("PointB parent circuit boxes - %v\n", pointB.ParentCircuit.JunctionBoxes)

		pointB.ParentCircuit.JunctionBoxes = append(pointB.ParentCircuit.JunctionBoxes,
			pointA.ParentCircuit.JunctionBoxes...)
		// fmt.Printf("Junction boxes after appending - %v\n", pointB.ParentCircuit.JunctionBoxes)

		for _, junctionBoxRef := range pointA.ParentCircuit.JunctionBoxes {
			junctionBoxRef.ParentCircuit = pointB.ParentCircuit
			// fmt.Printf("Assigning new parent to %p\n", junctionBoxRef)
			// fmt.Printf("New parent boxes - %v\n", junctionBoxRef.ParentCircuit)
		}

		circuits[indexOfCurrentCircuit] = Circuit{}
		connectedAmount++
		connectionIndex++
	}

	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i].JunctionBoxes) > len(circuits[j].JunctionBoxes)
	})

	// for index, _ := range circuits {
	// 	fmt.Println(circuits[index])
	// }

	result := 1

	fmt.Printf("Biggest numbers: ")
	for index := 0; index < 3; index++ {
		number := len(circuits[index].JunctionBoxes)
		fmt.Printf(" %d;", number)
		result *= number
	}
	fmt.Printf("\nMultiplication result - %d\n", result)
}
