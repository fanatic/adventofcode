package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	grid := [1000][1000]int{}

	for {
		var action string
		x1, y1, x2, y2 := 0, 0, 0, 0

		_, err := fmt.Scanf("%s", &action)
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if action == "turn" {
			_, err := fmt.Scanf("%s", &action)
			if err != nil && err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
		}

		_, err = fmt.Scanf("%d,%d through %d,%d\n", &x1, &y1, &x2, &y2)
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if action == "on" {
			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					grid[i][j] += 1
				}
			}
		}

		if action == "off" {
			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					grid[i][j] -= 1
					if grid[i][j] < 0 {
						grid[i][j] = 0
					}
				}
			}
		}

		if action == "toggle" {
			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					grid[i][j] += 2
				}
			}
		}
	}

	brightness := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			brightness += grid[i][j]
		}
	}
	fmt.Println("Brightness: ", brightness)

}
