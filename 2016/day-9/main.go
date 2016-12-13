package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		str = strings.TrimSpace(str)

		output := "" // decompress(str)
		outputLen := decompressLen(str)
		//fmt.Printf("Before: %q After: %q Length: %d\n", str, output, len(output))
		fmt.Printf("Before: %q Length: %d ?= %d\n", str, len(output), outputLen)
	}
}

func decompress(str string) string {
	output := ""
	for i := 0; i < len(str); i++ {
		if str[i] != '(' {
			output += string(str[i])
			continue
		}
		// Read data until "x" to get count
		i++
		markerCountStr := ""
		for ; str[i] != 'x'; i++ {
			markerCountStr += string(str[i])
		}

		// Read data until ")" to get repeat
		i++
		markerRepeatStr := ""
		for ; str[i] != ')'; i++ {
			markerRepeatStr += string(str[i])
		}

		count, _ := strconv.Atoi(markerCountStr)
		repeat, _ := strconv.Atoi(markerRepeatStr)
		i++

		//fmt.Printf("(%dx%d) Next: %q\n", count, repeat, str[i])

		// Take the subsequent X characters
		seq := str[i : i+count]
		output += strings.Repeat(decompress(seq), repeat)
		i += count - 1
	}
	return output
}

func decompressLen(str string) int {
	output := 0
	for i := 0; i < len(str); i++ {
		if str[i] != '(' {
			output++
			continue
		}
		// Read data until "x" to get count
		i++
		markerCountStr := ""
		for ; str[i] != 'x'; i++ {
			markerCountStr += string(str[i])
		}

		// Read data until ")" to get repeat
		i++
		markerRepeatStr := ""
		for ; str[i] != ')'; i++ {
			markerRepeatStr += string(str[i])
		}

		count, _ := strconv.Atoi(markerCountStr)
		repeat, _ := strconv.Atoi(markerRepeatStr)
		i++

		//fmt.Printf("(%dx%d) Next: %q\n", count, repeat, str[i])

		// Take the subsequent X characters
		output += decompressLen(str[i:i+count]) * repeat
		i += count - 1
	}
	return output
}
