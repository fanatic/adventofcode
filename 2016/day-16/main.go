package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := NewData("10011111011011001")
	diskLength := 35651584

	for i := 0; a.Length() < diskLength; i++ {
		fmt.Printf("Generate #%d: %s\n", i, a.StringWithLength())
		a.Generate()
	}
	a.Truncate(diskLength)
	fmt.Printf("Final: %s\n", a.StringWithLength())

	// Checksum
	for i := 0; a.length%2 == 0; i++ {
		a.Checksum()
		fmt.Printf("Checksum #%d: %s\n", i, a.StringWithLength())
	}
	fmt.Printf("Final Checksum: %s\n", a.StringWithLength())
}

type Data struct {
	length int
	data   []byte
}

func NewData(input string) Data {
	data := Data{
		length: len(input),
		data:   []byte{},
	}
	for i, c := range input {
		if c == '1' {
			data.set(i, 1)
		} else {
			data.set(i, 0)
		}
	}
	return data
}

func (r *Data) Generate() {
	l := r.length

	// Append 0
	r.set(r.length, 0)

	// Read in reverse order
	for i := l - 1; i >= 0; i-- {
		b := r.get(i)

		// Write bit flip
		if b == 0 {
			r.set(r.length, 1)
		} else {
			r.set(r.length, 0)
		}
	}
}

func (r *Data) Checksum() {
	checksum := NewData("")
	for i := 0; i < r.length; i += 2 {
		x := r.get(i)
		y := r.get(i + 1)
		if (x == 0 && y == 0) || (x == 1 && y == 1) {
			checksum.set(checksum.length, 1)
		} else {
			checksum.set(checksum.length, 0)
		}
	}
	*r = checksum
}

func (r *Data) Length() int {
	return r.length
}

func (r *Data) Truncate(l int) {
	r.length = l
}

func (r *Data) set(idx int, val uint8) {
	if idx < 0 || idx > r.Length() {
		panic(fmt.Sprintf("index %d is out of bounds", idx))
	} else if idx == r.Length() {
		r.length++
	}
	bidx := idx / 8
	if bidx >= len(r.data) {
		r.data = append(r.data, byte(0))
	}
	num := uint8(r.data[bidx])
	off := uint8(7 - (idx % 8))
	mask := uint8(1) << off
	if num&mask == mask {
		//is set
		if val == 0 {
			//so unset it
			r.data[bidx] = num ^ mask
		}
	} else {
		//not set
		if val == 1 {
			//so set it
			r.data[bidx] = num | mask
		}
	}
}

func (r *Data) get(idx int) uint8 {
	if idx < 0 || idx >= r.Length() {
		panic(fmt.Sprintf("index %d is out of bounds", idx))
	}
	bidx := idx / 8
	num := uint8(r.data[bidx])
	off := uint8(7 - (idx % 8))
	mask := uint8(1) << off
	if num&mask == mask {
		return 1
	}
	return 0
}

func (r *Data) String() string {
	var buf bytes.Buffer
	for _, v := range r.data {
		buf.WriteString(fmt.Sprintf("%08b", v))
	}
	return buf.String()[0:r.length]
}

func (r *Data) StringWithLength() string {
	if r.Length() > 50 {
		return fmt.Sprintf("[%d]", r.Length())
	}
	return fmt.Sprintf("%s [%d]", r.String(), r.Length())
}
