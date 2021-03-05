package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		printHelpAndExit()
	}

	num := os.Args[1]

	if _, err := strconv.ParseFloat(num, 64); err != nil {
		printHelpAndExit()
	}

	fmt.Println(commas(num))
}

func printHelpAndExit() {
	fmt.Println("Wrong directive")
	fmt.Println("Usage: commas 1297982")
	fmt.Println("Note: the argument must be a number")
	os.Exit(1)
}

func commas(s string) string {
	var b bytes.Buffer
	var posFromRight int
	n := len(s)

	for i, r := range s {
		posFromRight = n - i - 1

		if posFromRight%3 == 0 && posFromRight != 0 {
			b.WriteRune(r)
			b.WriteRune(',')
		} else {
			b.WriteRune(r)
		}
	}

	return b.String()
}
