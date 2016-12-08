package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		str, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		doorID := strings.TrimSpace(str)

		password := []rune{'_', '_', '_', '_', '_', '_', '_', '_'}
		for i := 0; i < 100000000000; i++ {
			hashable := doorID + fmt.Sprintf("%d", i)
			hash := fmt.Sprintf("%x", md5.Sum([]byte(hashable)))
			if hash[0:5] == "00000" {
				pos, err := strconv.Atoi(string(hash[5]))
				if err != nil {
					continue
				}
				char := rune(hash[6])
				if pos < 8 && password[pos] == '_' {
					password[pos] = char
					if password[0] != '_' &&
						password[1] != '_' &&
						password[2] != '_' &&
						password[3] != '_' &&
						password[4] != '_' &&
						password[5] != '_' &&
						password[6] != '_' &&
						password[7] != '_' {
						break
					}
				}
				fmt.Printf("Hashing %s -> hash %s. pos: %d. char: %q  %q\n", hashable, hash, pos, char, string(password))
			}
		}
		fmt.Printf("Password: %q\n", string(password))
	}
}
