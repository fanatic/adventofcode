package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Registers struct {
	A int
	B int
	C int
	D int
}

func (r *Registers) CopyInt(x int, y rune) {
	if y == 'a' {
		r.A = x
	} else if y == 'b' {
		r.B = x
	} else if y == 'c' {
		r.C = x
	} else if y == 'd' {
		r.D = x
	}
}

func (r *Registers) Get(x rune) int {
	if x == 'a' {
		return r.A
	} else if x == 'b' {
		return r.B
	} else if x == 'c' {
		return r.C
	} else if x == 'd' {
		return r.D
	}
	return -1
}

func (r *Registers) Copy(x, y rune) {
	if y == 'a' {
		r.A = r.Get(x)
	} else if y == 'b' {
		r.B = r.Get(x)
	} else if y == 'c' {
		r.C = r.Get(x)
	} else if y == 'd' {
		r.D = r.Get(x)
	}
}

func (r *Registers) Inc(x rune) {
	if x == 'a' {
		r.A++
	} else if x == 'b' {
		r.B++
	} else if x == 'c' {
		r.C++
	} else if x == 'd' {
		r.D++
	}
}

func (r *Registers) Dec(x rune) {
	if x == 'a' {
		r.A--
	} else if x == 'b' {
		r.B--
	} else if x == 'c' {
		r.C--
	} else if x == 'd' {
		r.D--
	}
}

func (r *Registers) NotZero(x rune) bool {
	if x == 'a' {
		return r.A != 0
	} else if x == 'b' {
		return r.B != 0
	} else if x == 'c' {
		return r.C != 0
	} else if x == 'd' {
		return r.D != 0
	}
	return false
}

func main() {

	r := &Registers{C: 1}
	program := []string{}

	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		program = append(program, strings.TrimSpace(str))
	}

	for pc := 0; pc < len(program); {
		//fmt.Printf("PC: %d   A: %d B: %d C: %d D: %d\n", pc, r.A, r.B, r.C, r.D)
		//fmt.Printf("%s\n", program[pc])
		terms := strings.Split(program[pc], " ")
		if terms[0] == "cpy" {
			x, err := strconv.Atoi(terms[1])
			if err != nil {
				r.Copy(rune(terms[1][0]), rune(terms[2][0]))
			} else {
				r.CopyInt(x, rune(terms[2][0]))
			}
		} else if terms[0] == "inc" {
			r.Inc(rune(terms[1][0]))
		} else if terms[0] == "dec" {
			r.Dec(rune(terms[1][0]))
		} else if terms[0] == "jnz" {
			x, err := strconv.Atoi(terms[1])
			if err != nil {
				if r.NotZero(rune(terms[1][0])) {
					y, _ := strconv.Atoi(terms[2])
					pc += y
					continue
				}
			} else {
				if x != 0 {
					y, _ := strconv.Atoi(terms[2])
					pc += y
					continue
				}
			}
		}
		pc++
	}
	fmt.Printf("Value in register A: %d\n", r.A)
}
