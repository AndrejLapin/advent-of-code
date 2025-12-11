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
	Id      string
	Outputs []*Device
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
			deviceRef.Outputs[index] = deviceMap[outputId]
		}
	}

	const start = "you"
	const end = "out"

	pathCount := 1

	searchPaths := []Device{}
	searchPaths = append(searchPaths, *deviceMap["you"])
	for len(searchPaths) > 0 {
		currentSearchPath := searchPaths[0]
		if currentSearchPath.Id == end {
			searchPaths = searchPaths[1:]
			continue
		}

		pathCount += len(currentSearchPath.Outputs) - 1

		newSearchPaths := []Device{}
		for _, output := range currentSearchPath.Outputs {
			newSearchPaths = append(newSearchPaths, *output)
		}
		searchPaths = append(searchPaths, newSearchPaths...)
		searchPaths = searchPaths[1:]
	}

	fmt.Println(pathCount)

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
