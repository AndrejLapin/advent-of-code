package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type RotatingDial struct {
	DialValue      int
	ClicksPastZero int
}

func (d *RotatingDial) Add(toAdd int) {
	fmt.Printf("R%d\n", toAdd)
	hundreds := toAdd / 100
	// fmt.Printf("Hundreds - %d\n", hundreds)
	clicksPastZero := hundreds
	toAdd -= hundreds * 100
	d.DialValue += toAdd
	if d.DialValue > 99 {
		d.DialValue -= 100
		clicksPastZero++
	}
	d.ClicksPastZero += clicksPastZero
	fmt.Printf("Dial value %d\n", d.DialValue)
	fmt.Printf("Clicks past zero %d\n", clicksPastZero)
}

func (d *RotatingDial) Sub(toSub int) {
	fmt.Printf("L%d\n", toSub)
	hundreds := toSub / 100
	// fmt.Printf("Hundreds - %d\n", hundreds)
	startedAsZero := d.DialValue == 0
	clicksPastZero := hundreds
	toSub -= hundreds * 100
	d.DialValue -= toSub
	if d.DialValue < 0 {
		d.DialValue += 100
		if !startedAsZero {
			clicksPastZero++
		}
	} else if d.DialValue == 0 {
		clicksPastZero++
	}
	d.ClicksPastZero += clicksPastZero
	fmt.Printf("Dial value %d\n", d.DialValue)
	fmt.Printf("Clicks past zero %d\n", clicksPastZero)
}

// func (d *RotatingDial) Add(toAdd int) {
// 	fmt.Printf("R%d\n", toAdd)
// 	skipClick := d.DialValue == 0
// 	d.DialValue += toAdd
// 	for d.DialValue > 99 {
// 		d.DialValue -= 100
// 		if skipClick {
// 			skipClick = false
// 			continue
// 		}
// 		d.ClicksPastZero++
// 		fmt.Println("Clicked past zero")
// 	}
// 	fmt.Printf("Dial value - %d\n", d.DialValue)
// }

// func (d *RotatingDial) Sub(toSub int) {
// 	fmt.Printf("L%d\n", toSub)
// 	skipClick := d.DialValue == 0
// 	d.DialValue -= toSub
// 	for d.DialValue < 0 {
// 		d.DialValue += 100
// 		if skipClick {
// 			skipClick = false
// 			continue
// 		}
// 		d.ClicksPastZero++
// 		fmt.Println("Clicked past zero")
// 	}
// 	if d.DialValue == 0 {
// 		d.ClicksPastZero++
// 		fmt.Println("Clicked past zero")
// 	}
// 	fmt.Printf("Dial value - %d\n", d.DialValue)
// }

func main() {
	// read the file
	// file, err := os.Open("test-input-1.txt")
	file, err := os.Open("input-1.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dial := RotatingDial{DialValue: 50, ClicksPastZero: 0}
	// dial.Add(1000)
	// fmt.Println(dial.ClicksPastZero)

	zeroCount := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		first := line[0]
		number, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		switch first {
		case 'R':
			dial.Add(number)
		case 'L':
			dial.Sub(number)
		}
		if dial.DialValue == 0 {
			zeroCount++
		}
	}
	// fmt.Println(zeroCount + dial.ClicksPastZero)
	fmt.Printf("Dial position - %d\n", dial.DialValue)
	fmt.Printf("Ended on zero - %d\n", zeroCount)
	fmt.Printf("Clicks past zero - %d\n", dial.ClicksPastZero)
}
