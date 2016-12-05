package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
)

func main() {
	numCharCode := 0
	numCharMemory := 0
	numCharEncoded := 0

	replacements, err := regexp.Compile(`(\\")|(\\\\)|(\\x[0-9a-f][0-9a-f])`)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		str = str[0 : len(str)-1]
		fmt.Printf("String: %s\n", str)
		numCharCode += len(str)

		parsed := str[1 : len(str)-1] // Strip ""
		parsed = replacements.ReplaceAllString(parsed, "?")

		fmt.Printf("Parsed: %q\n", parsed)
		numCharMemory += len(parsed)

		encoded := fmt.Sprintf("%q", str)
		fmt.Printf("Encoded: %s\n", encoded)
		numCharEncoded += len(encoded)
	}
	fmt.Printf("Part 1 Answer: %d\n", numCharCode-numCharMemory)
	fmt.Printf("Part 2 Answer: %d\n", numCharEncoded-numCharCode)
}
