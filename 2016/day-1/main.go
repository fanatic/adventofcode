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

		// facing: 0 N, 1 E, 2 S, 3 W
		var facing int

		str = strings.TrimSpace(str)
		directions := strings.Split(str, ", ")

		locationsVisited := [][]int{}

		x, y := 0, 0
		for _, direction := range directions {
			fmt.Printf("Next: %s", direction)
			if direction[0] == 'L' {
				if facing == 0 {
					facing = 3
				} else {
					facing = (facing - 1) % 4
				}

			} else if direction[0] == 'R' {
				facing = (facing + 1) % 4
			}
			distance, _ := strconv.Atoi(direction[1:])
			fmt.Printf("  Now facing %s.  Pacing %d blocks.", dir(facing), distance)

			// Lets pace each step and record it
			for i := 0; i < distance; i++ {
				if facing == 0 {
					y += 1
				} else if facing == 1 {
					x += 1
				} else if facing == 2 {
					y -= 1
				} else if facing == 3 {
					x -= 1
				}
				fmt.Printf("  Visiting %d,%d\n", x, y)
				if exists(x, y, locationsVisited) {
					fmt.Printf("We've visited this twice!!!\n")
					fmt.Printf("Blocks: %d (x:%d y:%d)\n", abs(x)+abs(y), x, y)

				}
				locationsVisited = append(locationsVisited, []int{x, y})
			}

		}
		fmt.Printf("End Blocks: %d (x:%d y:%d)\n", abs(x)+abs(y), x, y)
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func dir(i int) string {
	if i == 0 {
		return "North"
	} else if i == 1 {
		return "East"
	} else if i == 2 {
		return "South"
	} else if i == 3 {
		return "West"
	}
	return "!!"
}

func exists(x, y int, locations [][]int) bool {
	for _, loc := range locations {
		if loc[0] == x && loc[1] == y {
			return true
		}
	}
	return false
}
