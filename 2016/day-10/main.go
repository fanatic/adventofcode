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

var specific = regexp.MustCompile(`value (\d+) goes to bot (\d+)`)
var generic = regexp.MustCompile(`bot (\d+) gives low to (bot|output) (\d+) and high to (bot|output) (\d+)`)

type Bot struct {
	Index        int
	Chip         int
	Low          int
	LowisOutput  bool
	High         int
	HighisOutput bool
	Seen         []int
}

func (n *Bot) acceptChip(chip int) {
	n.Seen = append(n.Seen, chip)
	if n.Chip == 0 {
		n.Chip = chip
		return
	}
	// We have two chips now, so do the next part.
	highVal := 0
	lowVal := 0
	if n.Chip > chip {
		highVal = n.Chip
		lowVal = chip
		n.Chip = 0
	} else {
		highVal = chip
		lowVal = n.Chip
		n.Chip = 0
	}

	if n.LowisOutput {
		outputs[n.Low] = append(outputs[n.Low], lowVal)
	} else {
		fmt.Printf("Bot[%d]. Giving lower (%d) to bot[%d].\n", n.Index, lowVal, n.Low)
		bots[n.Low].acceptChip(lowVal)
	}
	if n.HighisOutput {
		outputs[n.High] = append(outputs[n.High], highVal)
	} else {
		fmt.Printf("Bot[%d]. Giving higher (%d) to bot[%d].\n", n.Index, highVal, n.High)
		bots[n.High].acceptChip(highVal)
	}
}

var bots = map[int]*Bot{}
var outputs = map[int][]int{}
var valueToBot = map[int]int{}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		str = strings.TrimSpace(str)

		if m := specific.FindStringSubmatch(str); m != nil {
			val, _ := strconv.Atoi(m[1])
			bot, _ := strconv.Atoi(m[2])
			valueToBot[val] = bot
		} else if m := generic.FindStringSubmatch(str); m != nil {
			idx, _ := strconv.Atoi(m[1])
			low, _ := strconv.Atoi(m[3])
			high, _ := strconv.Atoi(m[5])
			bots[idx] = &Bot{
				Index:        idx,
				Low:          low,
				LowisOutput:  (m[2] == "output"),
				High:         high,
				HighisOutput: (m[4] == "output"),
			}
		}
	}

	// Hand out initial values
	for val, bot := range valueToBot {
		bots[bot].acceptChip(val)
	}

	fmt.Printf("Outputs:\n")
	for i, vals := range outputs {
		fmt.Printf("  %d => %v\n", i, vals)
	}

	fmt.Printf("Bots:\n")
	for i, n := range bots {
		fmt.Printf(" %d => %v\n", i, n.Seen)
	}
}
