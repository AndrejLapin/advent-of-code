package main

import (
	"crypto/md5"
	"encoding/binary"
	"flag"
	"fmt"
	"strconv"
)

func main() {
	secretKey := flag.String("input", "ckczppom", "secret key")
	flag.Parse()
	var zerosCheck uint32 = 0xFFFFFF00
	hashStart := zerosCheck
	var currentNumber uint64 = 0
	for zerosCheck&hashStart != 0 {
		currentNumber++
		hash := md5.Sum([]byte(*secretKey + strconv.FormatUint(currentNumber, 10)))
		hashStart = binary.BigEndian.Uint32(hash[:4])
	}
	fmt.Println(currentNumber)
}
