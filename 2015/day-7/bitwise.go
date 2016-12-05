package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
)

var signal = regexp.MustCompile(`^(\w+) -> (\w+)\n$`)
var and = regexp.MustCompile(`^(\w+) (AND|OR|LSHIFT|RSHIFT) (\w+) -> (\w+)\n$`)
var not = regexp.MustCompile(`^NOT (\w+) -> (\w+)\n$`)

type Connection struct {
	In1 string
	Op  string
	In2 string
}

var conns = map[string]Connection{}
var valueCache = map[string]int{}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		if res := signal.FindStringSubmatch(str); res != nil {
			conns[res[2]] = Connection{res[1], "WIRE", ""}
		}
		if res := and.FindStringSubmatch(str); res != nil {
			conns[res[4]] = Connection{res[1], res[2], res[3]}
		}
		if res := not.FindStringSubmatch(str); res != nil {
			conns[res[2]] = Connection{res[1], "NOT", ""}
		}
	}

	// Resolve value for "a"
	a := resolve("a")
	fmt.Printf("Value of a is: %d\n", a)

	// Part Two
	valueCache = map[string]int{}
	valueCache["b"] = a
	fmt.Printf("Value of a is: %d\n", resolve("a"))
}

func resolveAndCache(wire string) int {
	if v, ok := valueCache[wire]; ok {
		return v
	}
	result := resolve(wire)
	valueCache[wire] = result
	return result
}

func resolve(wire string) int {
	if wire == "" {
		panic("oops")
	}
	// First, check if wire is int
	if i, err := strconv.Atoi(wire); err == nil {
		valueCache[wire] = i
		fmt.Printf("V %s is %d\n", wire, i)
		return i
	}

	// Okay, get connection
	in1 := resolveAndCache(conns[wire].In1)
	op := conns[wire].Op

	switch op {
	case "WIRE":
		fmt.Printf("W %s -> %d\n", wire, in1)
		return in1
	case "NOT":
		fmt.Printf("N %s -> %d\n", wire, ^in1)
		return ^in1
	}

	in2 := resolveAndCache(conns[wire].In2)
	switch op {
	case "AND":
		fmt.Printf("A %s -> %d\n", wire, in1&in2)
		return in1 & in2
	case "OR":
		fmt.Printf("O %s -> %d\n", wire, in1|in2)
		return in1 | in2
	case "LSHIFT":
		fmt.Printf("L %s -> %d\n", wire, in1<<uint(in2))
		return in1 << uint(in2)
	case "RSHIFT":
		fmt.Printf("R %s -> %d\n", wire, in1>>uint(in2))
		return in1 >> uint(in2)
	}
	panic("Oops")

}
