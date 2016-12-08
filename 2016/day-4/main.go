package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	sektorSum := 0
	for {
		str, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		str = strings.TrimSpace(str)

		r, _ := regexp.Compile(`^([a-z-]+)-([0-9]+)\[([a-z]+)\]$`)
		match := r.FindStringSubmatch(str)
		fmt.Printf("Analyzing %s: ", strings.TrimSpace(str))

		common := mostCommonLetters(match[1])[0:5]
		fmt.Printf("SectorID: %s, Freq: %s, Checksum: %s", match[2], string(common), match[3])
		if string(common) == match[3] {
			fmt.Printf(" ✓")
			sectorID, _ := strconv.Atoi(match[2])
			sektorSum += sectorID

			fmt.Printf("Decrypted: %s", decrypt(match[1], sectorID))
		}
		fmt.Printf("\n")
		//fmt.Printf("Name: %s, SectorID: %s, Checksum: %s\n", match[1], match[2], match[3])

	}
	fmt.Printf("Sum: %d\n", sektorSum)
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

func decrypt(str string, rotate int) string {
	s := []rune(str)
	for i := 0; i < rotate; i++ {
		for i := range s {
			if s[i] == '-' || s[i] == ' ' {
				s[i] = ' '
				continue
			}
			s[i]++
			if s[i] > 'z' {
				s[i] = 'a'
			}
		}
	}
	return string(s)
}
