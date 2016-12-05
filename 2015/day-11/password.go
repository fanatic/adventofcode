package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	password := ""
	reader := bufio.NewReader(os.Stdin)
	password, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	password = strings.Trim(password, "\n")
	fmt.Printf("Starting: %q\n", password)
	password = increment(password)
	for !hasIncreasingSequence(password) || hasInvalidLetter(password) || !hasTwoNonOverlappingPairs(password) {
		password = increment(password)
		//fmt.Printf("%q\n", password)
	}
	fmt.Printf("%q\n", password)
}

func increment(str string) string {
	for i := len(str) - 1; i > 0; i-- {
		if str[i] == 'z' {
			str = str[:i] + "a" + str[i+1:]
			continue
		}
		str = str[:i] + string(str[i]+1) + str[i+1:]
		return str
	}
	return str
}

func hasIncreasingSequence(str string) bool {
	if len(str) < 3 {
		return false
	}
	for i := 2; i < len(str); i++ {
		if str[i-2]+1 == str[i-1] &&
			str[i-1]+1 == str[i] {
			return true
		}
	}
	return false
}

func hasInvalidLetter(str string) bool {
	for _, c := range str {
		if c == 'i' || c == 'o' || c == 'l' {
			return true
		}
	}
	return false
}

func hasTwoNonOverlappingPairs(str string) bool {
	foundFirstPair := false
	if len(str) < 4 {
		return false
	}
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			if foundFirstPair {
				return true
			}
			foundFirstPair = true
			i++
		}
	}
	return false
}
