package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Device struct {
	Id              string
	SequenceId      int
	Inputs, Outputs []*Device
}

type ConnectingSearchState struct {
	Node       *Device
	Increasing bool
}

type ConnectionCount struct {
	FromLeft, FromRight int
}

func searchForward(firstSearchState Device, target string) int {
	pathCount := 1
	searchPaths := []Device{}
	searchPaths = append(searchPaths, firstSearchState)
	for len(searchPaths) > 0 {
		currentSearchPath := searchPaths[0]
		if currentSearchPath.Id == target {
			searchPaths = searchPaths[1:]
			continue
		}

		pathCount += len(currentSearchPath.Outputs) - 1

		newSearchPaths := []Device{}
		for _, ouput := range currentSearchPath.Outputs {
			newSearchPaths = append(newSearchPaths, *ouput)
		}
		searchPaths = append(searchPaths, newSearchPaths...)
		searchPaths = searchPaths[1:]
	}
	return pathCount
}

func searchBackward(firstSearchState Device, target string) int {
	pathCount := 1
	searchPaths := []Device{}
	searchPaths = append(searchPaths, firstSearchState)
	for len(searchPaths) > 0 {
		currentSearchPath := searchPaths[0]
		if currentSearchPath.Id == target {
			searchPaths = searchPaths[1:]
			continue
		}

		pathCount += len(currentSearchPath.Inputs) - 1
		newSearchPaths := []Device{}
		for _, input := range currentSearchPath.Inputs {
			newSearchPaths = append(newSearchPaths, *input)
		}
		searchPaths = append(searchPaths, newSearchPaths...)
		searchPaths = searchPaths[1:]
	}

	return pathCount
}

func findNode(firstNode, targetNode Device) int {
	pathCount := 0
	searchStates := []Device{firstNode}
	for len(searchStates) > 0 {
		currentState := searchStates[0]
		skipNode := false
		if currentState.Id == targetNode.Id {
			pathCount++
			skipNode = true

		} else if currentState.SequenceId > targetNode.SequenceId {
			skipNode = true
		}

		if skipNode {
			searchStates = searchStates[1:]
			continue
		}

		newSearchStates := []Device{}
		for _, output := range currentState.Outputs {
			newSearchStates = append(newSearchStates, *output)
		}
		searchStates = append(searchStates, newSearchStates...)
		searchStates = searchStates[1:]
	}
	return pathCount
}

func searchToMiddle(firstIncreasing, firstDecreasing *Device, targetIncreasing, targetDecreasing Device) int {
	const overlookAmount = 0
	connectionPoints := map[*Device]ConnectionCount{}
	searchStates := []ConnectingSearchState{
		{Node: firstIncreasing, Increasing: true},
		{Node: firstDecreasing, Increasing: false}}
	for len(searchStates) > 0 {
		currentSearchState := searchStates[0]
		_, ok := connectionPoints[currentSearchState.Node]
		connectNode := ok
		if !connectNode && len(searchStates) > 1 {
			for i := range searchStates[1:] {
				otherState := &searchStates[i+1]
				if otherState.Increasing != currentSearchState.Increasing && currentSearchState.Node == otherState.Node {
					// fmt.Printf("Creating new connection node for node %s\n", currentSearchState.Node.Id)
					connectNode = true
					newConnectionCount := ConnectionCount{}
					if currentSearchState.Increasing {
						newConnectionCount.FromLeft = 1
					} else {
						newConnectionCount.FromRight = 1
					}
					connectionPoints[currentSearchState.Node] = newConnectionCount
					break
				}
			}
		}
		if connectNode {
			// pathCount++
			connectionCount := connectionPoints[currentSearchState.Node]
			if currentSearchState.Increasing {
				connectionCount.FromLeft++
			} else {
				connectionCount.FromRight++
			}
			connectionPoints[currentSearchState.Node] = connectionCount
			searchStates = searchStates[1:]
			continue
		}
		endPath := false
		if currentSearchState.Increasing {
			endPath = currentSearchState.Node.SequenceId > targetIncreasing.SequenceId+overlookAmount
		} else {
			endPath = currentSearchState.Node.SequenceId < targetDecreasing.SequenceId+overlookAmount
		}
		if endPath {
			searchStates = searchStates[1:]
			continue
		}

		newSearchStates := []ConnectingSearchState{}
		if currentSearchState.Increasing {
			for _, output := range currentSearchState.Node.Outputs {
				newSearchStates = append(newSearchStates, ConnectingSearchState{Node: output, Increasing: true})
			}
		} else {
			for _, input := range currentSearchState.Node.Inputs {
				newSearchStates = append(newSearchStates, ConnectingSearchState{Node: input, Increasing: false})
			}
		}
		searchStates = append(searchStates, newSearchStates...)
		searchStates = searchStates[1:]
	}
	pathCount := 0
	for node, connectionCount := range connectionPoints {
		fmt.Printf("Connections under node %s, from left - %d, from right - %d\n",
			node.Id, connectionCount.FromLeft, connectionCount.FromRight)
		pathCount += connectionCount.FromLeft * connectionCount.FromRight
	}
	return pathCount
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

	devices := []Device{}
	deviceMap := map[string]*Device{}

	for scanner.Scan() {
		line := scanner.Text()
		nodeIds := strings.Fields(line)
		deviceId := nodeIds[0]

		deviceId = deviceId[:len(deviceId)-1]
		devices = append(devices, Device{Id: deviceId, SequenceId: -1})
	}
	devices = append(devices, Device{Id: "out"})

	for i := range devices {
		deviceRef := &devices[i]
		deviceMap[deviceRef.Id] = deviceRef
	}

	file.Seek(0, io.SeekStart)
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nodeIds := strings.Fields(line)
		deviceId := nodeIds[0]
		deviceId = deviceId[:len(deviceId)-1]
		outputIds := nodeIds[1:]
		deviceRef := deviceMap[deviceId]
		deviceRef.Outputs = make([]*Device, len(outputIds))
		for index, outputId := range outputIds {
			outputDevice := deviceMap[outputId]
			outputDevice.Inputs = append(outputDevice.Inputs, deviceRef)
			deviceRef.Outputs[index] = deviceMap[outputId]
		}
	}

	const start = "svr"
	const end = "out"
	const fft = "fft"
	const dac = "dac"

	{
		startDevice := deviceMap[start]
		startDevice.SequenceId = 0
		searchPaths := []*Device{}
		searchPaths = append(searchPaths, startDevice)
		for len(searchPaths) > 0 {
			currentSearchPath := searchPaths[0]
			// if currentSearchPath.SequenceId != -1 {
			// 	searchPaths = searchPaths[1:]
			// 	continue
			// }

			// currentSearchPath.SequenceId += 1

			newSearchPaths := []*Device{}
			for _, ouput := range currentSearchPath.Outputs {
				if ouput.SequenceId == -1 {
					ouput.SequenceId = currentSearchPath.SequenceId + 1
					newSearchPaths = append(newSearchPaths, ouput)
				}
			}
			searchPaths = append(searchPaths, newSearchPaths...)
			searchPaths = searchPaths[1:]
		}

	}

	fromFftToStart := searchBackward(*deviceMap[fft], start)
	fmt.Printf("From fft to srv %d\n", fromFftToStart)

	// ftomFftToDac := findNode(*deviceMap[fft], *deviceMap[dac])

	ftomFftToDac := searchToMiddle(deviceMap[fft], deviceMap[dac], *deviceMap[dac], *deviceMap[fft])
	fmt.Printf("From fft to dac %d\n", ftomFftToDac)

	// fromDacToFft := searchBackward(*deviceMap[dac], fft)
	// fmt.Printf("From dac to fft %d\n", fromDacToFft)

	fromDacToOut := searchForward(*deviceMap[dac], end)
	fmt.Printf("From dac to out %d\n", fromDacToOut)

	result := fromFftToStart * ftomFftToDac * fromDacToOut

	fmt.Printf("Result is %d\n", result)

	// for i := range devices {
	// 	device := &devices[i]
	// 	fmt.Printf("%s: sequence - %d\n", device.Id, device.SequenceId)
	// 	// for _, output := range device.Outputs {
	// 	// 	fmt.Printf("%s ", output.Id)
	// 	// }
	// 	// fmt.Printf("\n")
	// }
}
