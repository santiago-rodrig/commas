package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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
	pointIndex := strings.LastIndex(s, ".")
	signIndex := -1

	if pointIndex == -1 {
		pointIndex = len(s)
	}

	if !unicode.IsDigit(rune(s[0])) {
		b.WriteByte(s[0])
		signIndex = 0
	}

	strLen := len(s[signIndex+1 : pointIndex])

	for i, r := range s[signIndex+1 : pointIndex] {
		posFromRight = strLen - i - 1

		if posFromRight%3 == 0 && posFromRight != 0 {
			b.WriteRune(r)
			b.WriteRune(',')
		} else {
			b.WriteRune(r)
		}
	}

	if pointIndex != len(s) {
		b.WriteString(s[pointIndex:])
	}

	return b.String()
}
