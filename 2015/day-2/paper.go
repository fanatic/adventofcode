package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	paper := 0
	ribbon := 0
	//reader := bufio.NewReader(os.Stdin)
	for {
		var l int
		var w int
		var h int

		_, err := fmt.Scanf("%dx%dx%d\n", &l, &w, &h)
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		surfaceArea := 2*l*w + 2*w*h + 2*h*l
		areaOfSmallestSide := Min(l*w, Min(w*h, h*l))
		paper += surfaceArea + areaOfSmallestSide

		shortestPerimeter := Min(2*l+2*w, Min(w+w+h+h, h+h+l+l))
		volume := l * w * h
		ribbon += shortestPerimeter + volume
	}
	fmt.Println("Total wrapping paper: ", paper)
	fmt.Println("Total ribbon: ", ribbon)

}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
