package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// keypad := [][]int{
	// 	[]int{1, 2, 3},
	// 	[]int{4, 5, 6},
	// 	[]int{7, 8, 9},
	// }
	keypad := [][]int{
		[]int{-1, -1, 1, -1, -1},
		[]int{-1, 2, 3, 4, -1},
		[]int{5, 6, 7, 8, 9},
		[]int{-1, 10, 11, 12, -1},
		[]int{-1, -1, 13, -1, -1},
	}
	//x, y := 1, 1 // start at "5"
	x, y := 2, 0 // start at "5"

	for {
		str, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		for _, c := range str {
			if c == 'U' {
				x--
				if x == -1 || keypad[x][y] == -1 {
					x++
				}
			} else if c == 'R' {
				y++
				if y == len(keypad) || keypad[x][y] == -1 {
					y--
				}
			} else if c == 'D' {
				x++
				if x == len(keypad) || keypad[x][y] == -1 {
					x--
				}
			} else if c == 'L' {
				y--
				if y == -1 || keypad[x][y] == -1 {
					y++
				}
			} else if c == '\n' {
				continue
			}
			fmt.Printf("Moving %q to %s\n", c, print(keypad[x][y]))
		}
		fmt.Printf("Button is %s\n", print(keypad[x][y]))
	}
}

func print(i int) string {
	if i == 10 {
		return "A"
	} else if i == 11 {
		return "B"
	} else if i == 12 {
		return "C"
	} else if i == 13 {
		return "D"
	}
	return strconv.Itoa(i)
}
