package main

import (
	"fmt"
	"io"
	"log"
)

var goalTime = 2503

type Reindeer struct {
	Score            int
	Speed            int
	Distance         int
	Rest             int
	Position         int
	lastActionFlight bool
	flyingTimeLeft   int
	restingTimeLeft  int
}

func main() {
	r := map[string]*Reindeer{}
	for {
		var reindeer string
		var speed int
		var distance int
		var restTime int

		_, err := fmt.Scanf("%s can fly %d km/s for %d seconds, but then must rest for %d seconds.\n", &reindeer, &speed, &distance, &restTime)
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		r[reindeer] = &Reindeer{
			Speed:            speed,
			Distance:         distance,
			Rest:             restTime,
			lastActionFlight: true,
			flyingTimeLeft:   distance,
		}
	}

	for time := 1; time <= goalTime; time++ {
		for _, reindeer := range r {
			reindeer.Move()
			//fmt.Printf("%s: time:%ds position:%d\n", name, time, reindeer.Position)
		}
		topPosition := 0
		for _, reindeer := range r {
			if reindeer.Position > topPosition {
				topPosition = reindeer.Position
			}
		}

		// All reindeer in the lead get the point
		for name, reindeer := range r {
			if reindeer.Position == topPosition {
				fmt.Printf("%s wins with %d position\n", name, topPosition)
				reindeer.Score++
			}
		}
	}
	topScore := 0
	topReindeer := ""
	for name, reindeer := range r {
		if reindeer.Score > topScore {
			topScore = reindeer.Score
			topReindeer = name
		}
	}
	fmt.Printf("%s with %d\n", topReindeer, topScore)
}

func (r *Reindeer) Move() {
	if r.flyingTimeLeft > 0 {
		r.lastActionFlight = true
		r.Position += r.Speed
		r.flyingTimeLeft--
	} else if r.restingTimeLeft > 0 {
		r.lastActionFlight = false
		r.restingTimeLeft--
	} else {
		if r.lastActionFlight {
			r.restingTimeLeft = r.Rest - 1
		} else {
			r.Position += r.Speed
			r.flyingTimeLeft = r.Distance - 1
		}
	}
}
