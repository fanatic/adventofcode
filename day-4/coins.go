package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

var secretKey = "bgvyzdsv"

//var secretKey = "abcdef"

func main() {
	suffix := 1
	//fmt.Printf("%x\n", md5.Sum([]byte("abcdef609043")))
	for {
		input := fmt.Sprintf("%s%d", secretKey, suffix)
		hash := fmt.Sprintf("%x", md5.Sum([]byte(input)))
		if strings.HasPrefix(hash, "000000") {
			break
		}
		suffix += 1
		if suffix%100 == 0 {
			fmt.Printf("\rWorking on %d. Last hash: %s", suffix, hash)
		}
	}
	fmt.Println("\nAnswer: ", suffix)
}
