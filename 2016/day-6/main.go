package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	msgs := []string{}
	for {
		str, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		msgs = append(msgs, strings.TrimSpace(str))
	}

	decrypted := ""
	for i := 0; i < len(msgs[0]); i++ {
		letters := ""
		for _, msg := range msgs {
			letters += string(msg[i])
		}
		letterFreq := mostCommonLetters(letters)
		fmt.Printf("For index %d, most common letters are: %q, least common are: %q\n", i, letterFreq[:5], letterFreq[len(letterFreq)-6:len(letterFreq)-1])
		// decrypted += string(letterFreq[0])
		decrypted += string(letterFreq[len(letterFreq)-1])
	}
	fmt.Printf("Decrypted: %q\n", decrypted)
}

type Freq struct {
	Letter rune
	Count  int
}

type FreqList []Freq

func (f FreqList) Len() int { return len(f) }
func (f FreqList) Less(i, j int) bool {
	if f[i].Count == f[j].Count {
		return f[i].Letter > f[j].Letter
	}
	return f[i].Count < f[j].Count
}
func (f FreqList) Swap(i, j int) { f[i], f[j] = f[j], f[i] }

func mostCommonLetters(str string) []rune {
	freq := map[rune]int{}
	for _, c := range str {
		if c != '-' {
			freq[c]++
		}
	}
	fl := make(FreqList, len(freq))
	i := 0
	for k, v := range freq {
		fl[i] = Freq{k, v}
		i++
	}
	sort.Sort(sort.Reverse(fl))
	letters := make([]rune, len(freq))
	i = 0
	for _, f := range fl {
		letters[i] = f.Letter
		i++
	}
	return letters
}
