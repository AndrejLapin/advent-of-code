package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	fileName := flag.String("input", "input.txt", "input file")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	currentFloor := 0
	characterPosition := 0
	for {
		characterPosition++
		readByte, err := reader.ReadByte()
		if err != nil {
			break
		}
		switch readByte {
		case '(':
			currentFloor++
		case ')':
			currentFloor--
		}
		if currentFloor == -1 {
			break
		}
	}
	fmt.Println(characterPosition)
}
