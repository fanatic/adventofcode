package main

import (
	"fmt"
	"io"
	"log"

	"github.com/fighterlyt/permutation"
)

func main() {
	seatOrder := []string{}
	preferences := map[string]map[string]int{}
	for {
		var a string
		var b string
		var c int
		var d string

		_, err := fmt.Scanf("%s would %s %d happiness units by sitting next to %s", &a, &b, &c, &d)
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		if preferences[a] == nil {
			preferences[a] = map[string]int{}
		}
		if b == "lose" {
			c = -c
		}
		preferences[a][d[:len(d)-1]] = c
	}

	// Add myself
	preferences["me"] = map[string]int{}
	for guest := range preferences {
		preferences[guest]["me"] = 0
		preferences["me"][guest] = 0
	}

	// Initialize seat order
	for guest := range preferences {
		seatOrder = append(seatOrder, guest)
	}

	p, err := permutation.NewPerm(seatOrder, nil) //generate a Permutator
	if err != nil {
		log.Fatal(err)
	}
	bestOverallHappiness := -999999999
	for i, err := p.Next(); err == nil; i, err = p.Next() {
		seats := i.([]string)

		// Calculate happiness
		happiness := 0
		for i, guest := range seats {
			leftGuest := seats[len(seats)-1]
			if i != 0 {
				leftGuest = seats[i-1]
			}

			rightGuest := seats[0]
			if i != len(seats)-1 {
				rightGuest = seats[i+1]
			}

			happiness += preferences[guest][leftGuest]
			happiness += preferences[guest][rightGuest]
		}
		if happiness > bestOverallHappiness {
			bestOverallHappiness = happiness
		}
		//fmt.Printf("%5d/%5d: %v --> %d\n", p.Index(), p.Index()+p.Left(), seats, happiness)
	}
	fmt.Printf("Best overall happiness level: %d\n", bestOverallHappiness)

}
