package main

import (
	"fmt"
	"sort"
)

type Elf struct {
	index    int
	presents int
}

func main() {
	elfCount := 100
	elfCount = 3004953
	elves := make([]Elf, elfCount, elfCount)
	for i := range elves {
		elves[i] = Elf{i + 1, 1}
	}
OUTER:
	for {
		deleteList := []int{}
		for i := 0; i < len(elves); i++ {
			fmt.Printf("Elf %d", elves[i].index)

			nextElfWithPresents := nextElfWithPresents(i, elves)

			fmt.Printf(" takes Elf %d's %d presents.\n", elves[nextElfWithPresents].index, elves[nextElfWithPresents].presents)
			elves[i].presents += elves[nextElfWithPresents].presents

			if elves[i].presents == elfCount {
				fmt.Printf("All done.  Elf %d gets all the presents.\n", elves[i].index)
				break OUTER
			}
			i++
			deleteList = append(deleteList, nextElfWithPresents)
		}
		// Now delete all at once
		sort.Sort(sort.Reverse(sort.IntSlice(deleteList)))
		for _, i := range deleteList {
			elves = append(elves[:i], elves[i+1:]...)
		}
	}
}

// part 1
func nextElfWithPresents(i int, elves []Elf) int {
	next := i + 1
	if next == len(elves) {
		next = 0
	}
	return next
}

// part 2
func nextElfWithPresents2(i int, elves []Elf) int {
	// pre-shrink circle
	nextElves := []Elf{}
	for i := range elves {
		if elves[i].presents > 0 {
			nextElves = append(nextElves, elves[i])
		}
	}
	next := len(nextElves)/2 + i
	if next == len(elves) {
		next = 0
	}
	return next
}
