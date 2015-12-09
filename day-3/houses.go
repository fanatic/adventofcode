package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	x := map[bool]int{false: 0, true: 0}
	y := map[bool]int{false: 0, true: 0}
	homesSeen := map[string]bool{"0,0": true}
	homesSeenMoreThanOnce := 0
	deliverer := false
	reader := bufio.NewReader(os.Stdin)
	for {
		r, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		switch r {
		case '^':
			y[deliverer] += 1
		case 'v':
			y[deliverer] -= 1
		case '>':
			x[deliverer] += 1
		case '<':
			x[deliverer] -= 1
		}
		fmt.Printf("%d,%d", x[deliverer], y[deliverer])
		if _, exists := homesSeen[fmt.Sprintf("%d,%d", x[deliverer], y[deliverer])]; exists {
			homesSeenMoreThanOnce += 1
			fmt.Printf(" *\n")
		} else {
			homesSeen[fmt.Sprintf("%d,%d", x[deliverer], y[deliverer])] = true
			fmt.Printf("\n")
		}
		deliverer = !deliverer
	}
	fmt.Println("Homes seen: ", len(homesSeen))
	fmt.Println("Homes seen more than once: ", homesSeenMoreThanOnce)
}
