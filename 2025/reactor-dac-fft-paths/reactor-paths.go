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
	IterationCount  int
	Inputs, Outputs []*Device
}

type SearchState struct {
	Node                   *Device
	dacVisited, fftVisited bool
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
	foundAfterXIterations := -1
	const maxNewIterations = 3
	pathCount := 1
	searchPaths := []Device{}
	searchPaths = append(searchPaths, firstSearchState)
	for len(searchPaths) > 0 {
		currentSearchPath := searchPaths[0]
		if currentSearchPath.Id == target {
			if foundAfterXIterations == -1 {
				foundAfterXIterations = currentSearchPath.IterationCount
			}
			searchPaths = searchPaths[1:]
			continue
		}

		currentSearchPath.IterationCount += 1

		if foundAfterXIterations != -1 && currentSearchPath.IterationCount >
			foundAfterXIterations+maxNewIterations {
			searchPaths = searchPaths[1:]
			pathCount -= 1
			continue
		}

		pathCount += len(currentSearchPath.Inputs) - 1
		newSearchPaths := []Device{}
		for _, input := range currentSearchPath.Inputs {
			newSearchPaths = append(newSearchPaths, Device{
				Id:             input.Id,
				IterationCount: currentSearchPath.IterationCount,
				Inputs:         input.Inputs, Outputs: input.Outputs,
			})
		}
		searchPaths = append(searchPaths, newSearchPaths...)
		searchPaths = searchPaths[1:]
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
		devices = append(devices, Device{Id: deviceId})
		// deviceRef := deviceMap[deviceId]
		// if deviceRef == nil {
		// 	devices = append(devices, Device{Id: deviceId})
		// 	deviceRef = &devices[len(devices)-1]
		// 	deviceMap[deviceId] = deviceRef
		// }
		// deviceRef.Outputs = make([]*Device, len(outputIds))

		// outputIds := nodeIds[1:]
		// for outputIndex := range outputIds {
		// 	outputId := &outputIds[outputIndex]
		// 	outputRef := deviceMap[*outputId]
		// 	if outputRef == nil {
		// 		devices = append(devices, Device{Id: *outputId})
		// 		outputRef := &devices[len(devices)-1]
		// 		deviceMap[*outputId] = outputRef
		// 	}
		// 	deviceRef.Outputs[outputIndex] = outputRef
		// }
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

	// dac is after fft
	// start to fft - starts too soon or something?
	// dac to out
	// probably from dact back to fft

	// 3 nodes that let you pass to fft
	// 3 nodes that let you pass to dac

	// need 4 things

	// paths from dac to out
	// paths fft to dac (in reverse from dac)
	// paths from srv to fft (in reverse from fft)

	fromFftToStart := searchBackward(*deviceMap[fft], start)
	fmt.Printf("From fft to srv %d\n", fromFftToStart)

	fromDacToFft := searchBackward(*deviceMap[dac], fft)
	fmt.Printf("From dac to fft %d\n", fromDacToFft)

	// these have to navigate towards eachother, the forward head should find dac
	// the backward head should find fft
	// but how do we know when to halt them?
	// how do we make them not go past eachother?
	// there must be multiple way to go from fft to dac

	// in the current backward search dac finds all paths to fft
	// but then after some time it just goes past it and continues searching for useless paths
	// let's explore the pattern maybe?

	fromDacToOut := searchForward(*deviceMap[dac], end)
	fmt.Printf("From dac to out %d\n", fromDacToOut)

	result := fromFftToStart * fromDacToFft * fromDacToOut

	fmt.Printf("Result is %d\n", result)

	// for i := range devices {
	// 	device := &devices[i]
	// 	fmt.Printf("%s: ", device.Id)
	// 	for _, output := range device.Outputs {
	// 		fmt.Printf("%s ", output.Id)
	// 	}
	// 	fmt.Printf("\n")
	// }
	// scanner.Reset()
}
