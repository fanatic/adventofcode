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

func main() {
	reader := bufio.NewReader(os.Stdin)

	possible := 0
	triangles := [][]int{
		[]int{},
		[]int{},
		[]int{},
	}
	for {
		str, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		str = strings.TrimSpace(str)
		re := regexp.MustCompile(` +`)

		for i, sideStr := range re.Split(str, 3) {
			side, _ := strconv.Atoi(sideStr)
			triangles[i] = append(triangles[i], side)
		}

		if len(triangles[0]) == 3 {
			for _, triangle := range triangles {
				fmt.Printf("Checking: %d %d %d\n", triangle[0], triangle[1], triangle[2])
				if triangle[0]+triangle[1] > triangle[2] &&
					triangle[1]+triangle[2] > triangle[0] &&
					triangle[0]+triangle[2] > triangle[1] {
					possible++
				}
			}
			triangles = [][]int{
				[]int{},
				[]int{},
				[]int{},
			}
		}
	}
	fmt.Printf("Possible triangles: %d\n", possible)
}
