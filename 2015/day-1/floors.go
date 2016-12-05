package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	floor := 0
	reader := bufio.NewReader(os.Stdin)
	for pos := 1; ; pos += 1 {
		r, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		switch r {
		case '(':
			floor += 1
		case ')':
			floor -= 1
		}
		if floor == -1 {
			fmt.Println("First enters basement: ", pos)
		}
	}
	fmt.Println("Final floor: ", floor)
}
