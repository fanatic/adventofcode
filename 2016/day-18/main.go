package main

import "fmt"

func main() {
	row := "..^^."
	row = ".^^.^.^^^^"
	row = ".^.^..^......^^^^^...^^^...^...^....^^.^...^.^^^^....^...^^.^^^...^^^^.^^.^.^^..^.^^^..^^^^^^.^^^..^"
	numRows := 400000
	safeTileCount := numSafeTiles(row)

	//fmt.Printf("%s %d\n", row, numSafeTiles(row))
	for j := 0; j < numRows-1; j++ {
		nextRow := ""
		for i := 0; i < len(row); i++ {
			tile := ""
			if i == 0 {
				tile = isTrap('.', row[i], row[i+1])
			} else if i == len(row)-1 {
				tile = isTrap(row[i-1], row[i], '.')
			} else {
				tile = isTrap(row[i-1], row[i], row[i+1])
			}
			nextRow += tile
		}
		safeTileCount += numSafeTiles(nextRow)
		//fmt.Printf("%s %d\n", nextRow, numSafeTiles(nextRow))
		row = nextRow
	}
	fmt.Printf("Safe tiles: %d\n", safeTileCount)
}

func numSafeTiles(row string) int {
	safe := 0
	for _, tile := range row {
		if tile == '.' {
			safe++
		}
	}
	return safe
}

func isTrap(l, c, r byte) string {
	left := tileType(l)
	center := tileType(c)
	right := tileType(r)
	trap := (left && center && !right) ||
		(!left && center && right) ||
		(left && !center && !right) ||
		(!left && !center && right)
	if trap {
		return "^"
	}
	return "."
}

func tileType(tile byte) bool {
	return tile == '^'
}
