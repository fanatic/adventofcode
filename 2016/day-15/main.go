package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`Disc #\d+ has (\d+) positions; at time=0, it is at position (\d+).`)

type Disc struct {
	Index           int
	CurrentPosition int
	Positions       int
}

func (d *Disc) DropPositions() {
	// Drop position is the times where this
	// disc is at position 0 assuming we're at time 0 now

}

func main() {
	discs := []Disc{}
	reader := bufio.NewReader(os.Stdin)
	for i := 0; ; i++ {
		str, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		m := re.FindStringSubmatch(str)
		s, _ := strconv.Atoi(m[2])
		p, _ := strconv.Atoi(m[1])
		//fmt.Printf("Disk %d %d %d\n", i, s, p)
		discs = append(discs, Disc{i, s, p})
	}

	// Part Two
	discs = append(discs, Disc{len(discs), 0, 11})

	for time := 0; ; time++ {
		fmt.Printf("--[time=%d]--\n", time)
		allDiscsFallThrough := true
		for i, disc := range discs {
			innerTime := time + i + 1
			offset := (disc.Positions + disc.CurrentPosition + innerTime) % disc.Positions
			if offset == 0 {
				fmt.Printf("[time=%d] Capsule falls through disc %d (position: %d).\n", innerTime, i+1, offset)
			} else {
				fmt.Printf("[time=%d] Capsule blocked by disc %d (position: %d).\n", innerTime, i+1, offset)
				allDiscsFallThrough = false
				break
			}
		}
		if allDiscsFallThrough {
			fmt.Printf("Press the button at time=%d\n", time)
			break
		}
	}

}
