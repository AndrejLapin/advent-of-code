package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func invalid_basic(number int) bool {
	stringified := strconv.Itoa(number)
	if len(stringified)%2 != 0 {
		return false
	}
	half_length := len(stringified) / 2
	first_half := stringified[0:half_length]
	second_half := stringified[half_length:]
	return first_half == second_half
}

func invalid_complex(number int) bool {
	// fmt.Printf("Trying number %d\n", number)
	stringified := strconv.Itoa(number)
	if len(stringified) < 2 {
		return false
	}
	for part_length := len(stringified) / 2; part_length > 0; part_length-- {
		// fmt.Printf("Trying to part into lengths %d\n", part_length)
		if len(stringified)%part_length != 0 {
			// fmt.Println("Skipping")
			continue
		}
		part_amount := len(stringified) / part_length
		// fmt.Printf("Trying to part into %d\n", part_amount)
		parts := []string{}
		for part_index := 0; part_index < part_amount; part_index++ {
			part_start := part_index * part_length
			part := stringified[part_start : part_start+part_length]
			// fmt.Printf("Constructed part %s\n", part)
			parts = append(parts, part)
		}
		// fmt.Printf("Number %d, was parted into:\n", number)
		// fmt.Println(parts)
		is_invalid := true
		for i := 1; i < len(parts); i++ {
			if parts[0] != parts[i] {
				is_invalid = false
				break
			}
		}
		if is_invalid {
			// fmt.Println("Is invalid")
			return true
		}
	}
	return false
}

func main() {
	// invalid_complex(2121212118)

	var invalid_id_sum uint64 = 0
	file, err := os.Open("input-2.txt")
	// file, err := os.Open("test-input-2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	read := true
	for read {
		range_string, err := reader.ReadString(',')
		if err == io.EOF {
			read = false
		} else if err != nil {
			panic(err)
		} else {
			range_string = range_string[:len(range_string)-1]
		}

		// fmt.Println(range_string)

		range_seperator_index := strings.Index(range_string, "-")
		range_start, err := strconv.Atoi(range_string[0:range_seperator_index])
		if err != nil {
			panic(err)
		}
		range_end, err := strconv.Atoi(range_string[range_seperator_index+1:])
		if err != nil {
			panic(err)
		}
		// fmt.Println(range_start)
		// fmt.Println(range_end)
		for number := range_start; number <= range_end; number++ {
			if invalid_complex(number) {
				// fmt.Println(number)
				invalid_id_sum += uint64(number)
			}

			// fmt.Printf("First half %s, second half %s\n", first_half, second_half)
		}
	}
	fmt.Println(invalid_id_sum)
}
