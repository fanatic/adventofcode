package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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
		ip := strings.TrimSpace(str)

		supportsTLS := false
		hasABBA := false
		hasABBAInHypernet := false
		inBracket := false

		supportsSSL := false
		allABAs := []string{}
		allBABs := []string{}

		for i := range ip {
			if ip[i] == '[' {
				inBracket = true
			} else if ip[i] == ']' {
				inBracket = false
			}
			if i >= 2 {
				// Same character twice with a different character between them
				if ip[i-2] == ip[i] &&
					ip[i-2] != ip[i-1] &&
					ip[i-2] != '[' && ip[i-2] != ']' &&
					ip[i-1] != '[' && ip[i-1] != ']' &&
					ip[i] != '[' && ip[i] != ']' {
					if !inBracket {
						allABAs = append(allABAs, ip[i-2:i+1])
					} else {
						allBABs = append(allBABs, ip[i-2:i+1])
					}
				}
			}
			if i >= 3 {
				// Two different characters followed by reverse of that pair
				if ip[i-3] == ip[i] && ip[i-2] == ip[i-1] &&
					ip[i] != ip[i-1] &&
					ip[i-3] != '[' && ip[i-3] != ']' &&
					ip[i-2] != '[' && ip[i-2] != ']' &&
					ip[i-1] != '[' && ip[i-1] != ']' &&
					ip[i] != '[' && ip[i] != ']' {
					if inBracket {
						hasABBAInHypernet = true
					} else {
						hasABBA = true
					}
				}
			}
		}
		if hasABBA && !hasABBAInHypernet {
			supportsTLS = true
		}

		for i := range allABAs {
			for j := range allBABs {
				if allABAs[i][0] == allBABs[j][1] &&
					allABAs[i][1] == allBABs[j][0] {
					supportsSSL = true
				}
			}
		}

		fmt.Printf("%s: ABBA: %t. ABBA In Hypernet: %t.  TLS support: %t.\n", ip, hasABBA, hasABBAInHypernet, supportsTLS)
		fmt.Printf("%s: ABAs: %v.  BABs: %v SSL support: %t.\n", ip, allABAs, allBABs, supportsSSL)
	}
}
