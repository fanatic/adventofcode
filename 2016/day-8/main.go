package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Screen [6][50]bool

//type Screen [3][7]bool

var rect = regexp.MustCompile(`(\d+)x(\d+)`)
var equals = regexp.MustCompile(`[xy]=(\d+)`)

func main() {
	reader := bufio.NewReader(os.Stdin)

	screen := &Screen{}

	for {
		str, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		cmd := strings.Split(strings.TrimSpace(str), " ")
		if cmd[0] == "rect" {
			// rect 3x2
			m := rect.FindStringSubmatch(cmd[1])
			x, _ := strconv.Atoi(m[1])
			y, _ := strconv.Atoi(m[2])

			for i := 0; i < x; i++ {
				for j := 0; j < y; j++ {
					screen[j][i] = true
				}
			}
		} else if cmd[0] == "rotate" && cmd[1] == "column" {
			// rotate column x=1 by 1
			m := equals.FindStringSubmatch(cmd[2])
			col, _ := strconv.Atoi(m[1])
			amount, _ := strconv.Atoi(cmd[4])

			max := len(screen)
			for i := 0; i < amount; i++ {
				savedValue := screen[max-1][col]
				for j := 0; j < max; j++ {
					pos := j - 1
					if pos == -1 {
						pos = max - 1
					}
					v := screen[j][col]
					screen[j][col] = savedValue
					savedValue = v
				}
			}

		} else if cmd[0] == "rotate" && cmd[1] == "row" {
			// rotate row y=0 by 4
			m := equals.FindStringSubmatch(cmd[2])
			row, _ := strconv.Atoi(m[1])
			amount, _ := strconv.Atoi(cmd[4])

			max := len(screen[0])
			for i := 0; i < amount; i++ {
				savedValue := screen[row][max-1]
				for j := 0; j < max; j++ {
					pos := j - 1
					if pos == -1 {
						pos = max - 1
					}
					v := screen[row][j]
					screen[row][j] = savedValue
					savedValue = v
				}
			}

		}
		fmt.Printf("Running %q\n", strings.TrimSpace(str))
		printScreen(screen)
	}

	litPixels := 0
	for _, row := range screen {
		for _, col := range row {
			if col {
				litPixels++
			}
		}
	}
	fmt.Printf("Lit pixels: %d\n", litPixels)
}

func printScreen(screen *Screen) {
	for _, row := range screen {
		for _, col := range row {
			if col {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}
