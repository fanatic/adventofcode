package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	nice := 0
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		//if strings.Count(str, "a")+strings.Count(str, "e")+strings.Count(str, "i")+strings.Count(str, "o")+strings.Count(str, "u") < 3 {
		//	fmt.Printf("Found less than three vowels: %d\n", strings.Count(str, "aeiou"))
		//	continue
		//}
		if !twice(str) {
			fmt.Printf("t")
			continue
		}
		if !repeat(str) {
			fmt.Printf("r")
			continue
		}
		//if strings.Contains(str, "ab") ||
		//	strings.Contains(str, "cd") ||
		//	strings.Contains(str, "pq") ||
		//	strings.Contains(str, "xy") {
		//	fmt.Printf("Contains bad string\n")
		//	continue
		//}

		nice += 1
	}
	fmt.Println("Nice: ", nice)
}

func twice(str string) bool {
	pairs := map[string]int{}
	for i := 0; i < len(str)-1; i += 1 {
		s := str[i : i+2]
		if _, exists := pairs[s]; exists && pairs[s] < i-1 {
			//fmt.Printf("%d < %d\n", pairs[s], i)
			return true
		} else if !exists {
			pairs[s] = i
		}
	}
	//fmt.Printf("%v\n", pairs)
	return false
}

func repeat(str string) bool {
	for i := 2; i < len(str); i += 1 {
		if str[i] == str[i-2] {
			return true
		}
	}
	return false
}
