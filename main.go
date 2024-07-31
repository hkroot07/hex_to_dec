package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {
	var input string
	fmt.Scanln(&input)
	input = strings.ToLower(input)
	if input == "stop" {
		os.Exit(2)
	}
	i := new(big.Int)
	i.SetString(input, 16)
}

func processHex(hexStr string) string {
	return strings.TrimPrefix(hexStr, "0x")
}
