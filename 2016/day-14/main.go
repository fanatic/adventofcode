package main

import (
	"crypto/md5"
	"fmt"
	"sort"
)

func main() {
	salt := "abc"
	salt = "qzyelonm"

	tripleHashes := map[int]rune{}
	quadHashes := map[int]rune{}
OUTER:
	for i := 0; ; i++ {
		// Take MD5 of pre-arranged salt and incrasing integer
		hash := salt + fmt.Sprintf("%d", i)

		// Key streching
		for j := 0; j < 2017; j++ {
			hash = fmt.Sprintf("%x", md5.Sum([]byte(hash)))
		}

		foundQuad := false
		// Check for consequtive characters
		for j := range hash {
			if j >= 2 &&
				hash[j] == hash[j-1] &&
				hash[j-1] == hash[j-2] {
				tripleHashes[i] = rune(hash[j])
				fmt.Printf("[%d] triple: %c\n", i, hash[j])
			}
			if j >= 4 &&
				hash[j] == hash[j-1] &&
				hash[j-1] == hash[j-2] &&
				hash[j-2] == hash[j-3] &&
				hash[j-3] == hash[j-4] {
				quadHashes[i] = rune(hash[j])
				fmt.Printf("[%d] quadruple: %c\n", i, hash[j])
				foundQuad = true
			}
		}

		if foundQuad {

			valids := []int{}
			for i, ir := range tripleHashes {
				for j, jr := range quadHashes {
					if ir == jr && j > i && j < i+1000 {
						valids = append(valids, i)
						if len(valids) == 64 {
							sort.Ints(valids)
							fmt.Printf("Valid: %d %v\n", len(valids), valids)
							fmt.Printf("64th key: %d\n", valids[63])
							break OUTER
						}
						break
					}
				}
			}
			fmt.Printf("Valid: %d %v\n", len(valids), valids)
		}
	}
}
