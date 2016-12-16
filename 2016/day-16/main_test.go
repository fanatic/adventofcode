package main

import "testing"

func TestGenerate(t *testing.T) {
	var tests = []struct {
		in  string
		out string
	}{
		{"1", "100"},
		{"0", "001"},
		{"11111", "11111000000"},
		{"111100001010", "1111000010100101011110000"},
	}
	for _, tt := range tests {
		a := NewData(tt.in)
		a.Generate()
		s := a.String()
		if s != tt.out {
			t.Errorf("Generate(%q) => %q, want %q", tt.in, s, tt.out)
		}
	}
}

func BenchmarkGenerate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := NewData("10011111011011001001100100100000110")
		a.Generate()
	}
}
func BenchmarkChecksum(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := NewData("0111110100111010100111110110111010")
		a.Checksum()
	}
}
