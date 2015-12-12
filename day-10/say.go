package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	sequence := []int{}
	reader := bufio.NewReader(os.Stdin)
	for {
		r, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		sequence = append(sequence, int(r-'0'))
	}
	fmt.Printf("%d: %v\n", 0, sequence)
	for i := 0; i < 50; i++ {
		sequence = sequenceToSay(sequence)
		fmt.Printf("%d: %v\n", i+1, sequence)
	}
	fmt.Printf("Length of 50th: %d\n", len(sequence))
}

func sayToSequence(sequence []int) []int {
	nextSequence := []int{}
	for j := 0; j < len(sequence); j += 2 {
		for k := 0; k < sequence[j]; k++ {
			nextSequence = append(nextSequence, sequence[j+1])
		}
	}
	return nextSequence
}

func sequenceToSay(sequence []int) []int {
	nextSequence := []int{}
	lastDigit := -1
	numberOfLastDigit := 1
	for j := 0; j < len(sequence); j++ {
		if lastDigit == -1 {
			lastDigit = sequence[j]
			numberOfLastDigit = 1
		} else if lastDigit == sequence[j] {
			numberOfLastDigit++
		} else {
			nextSequence = append(nextSequence, []int{numberOfLastDigit, lastDigit}...)
			lastDigit = sequence[j]
			numberOfLastDigit = 1
		}
	}
	nextSequence = append(nextSequence, []int{numberOfLastDigit, lastDigit}...)

	return nextSequence
}
