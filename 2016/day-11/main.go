package main

import (
	"fmt"
	"sort"
)

type Object struct {
	Floor   int
	Element rune
	Type    rune
}

var possibleElements []rune

func main() {
	// The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
	// The second floor contains a hydrogen generator.
	// The third floor contains a lithium generator.
	// The fourth floor contains nothing relevant.
	var objects = []Object{
		Object{2, 'H', 'G'},
		Object{1, 'H', 'M'},
		Object{3, 'L', 'G'},
		Object{1, 'L', 'M'},
	}
	possibleElements = []rune{'H', 'L'}

	// The first floor contains a polonium generator, a thulium generator, a thulium-compatible microchip,
	// a promethium generator, a ruthenium generator, a ruthenium-compatible microchip,
	// a cobalt generator, and a cobalt-compatible microchip.
	// The second floor contains a polonium-compatible microchip and a promethium-compatible microchip.
	// The third floor contains nothing relevant.
	// The fourth floor contains nothing relevant.
	// objects = []Object{
	// 	Object{1, 'O', 'G'},
	// 	Object{1, 'T', 'G'},
	// 	Object{1, 'T', 'M'},
	// 	Object{1, 'R', 'G'},
	// 	Object{1, 'U', 'G'},
	// 	Object{1, 'U', 'M'},
	// 	Object{1, 'C', 'G'},
	// 	Object{1, 'C', 'M'},

	// 	Object{2, 'O', 'M'},
	// 	Object{2, 'R', 'M'},
	// }
	// possibleElements = []rune{'O', 'T', 'R', 'U', 'C'}

	printObjects(objects, 1)
	for i := 0; ; i++ {
		success := step(objects, 1, 0, i, State{})
		if success {
			fmt.Printf("Minimum number of steps: %d\n", i)
			break
		} else {
			fmt.Printf("[%d] seen states = %d\n", i, len(seenStates))
		}
	}
}

func step(objects []Object, elevatorFloor, stepCount, stepCountLimit int, lastState State) bool {
	if stepCount > stepCountLimit {
		return false
	}
	currentState := State{elevatorFloor, genPairs(objects)}
	if stateEqual(currentState, lastState) || seenBefore(currentState) {
		return false
	}
	if isFailure(objects, elevatorFloor) {
		updateState(currentState)
		return false
	}
	//fmt.Printf("step(floor: %d, count: %d)\n", elevatorFloor, stepCount)
	//printObjects(objects, elevatorFloor)

	if isSuccess(objects) {
		return true
	}

	// Move elevator with single object
	for i, obj := range objects {
		if obj.Floor != elevatorFloor {
			continue
		}
		if elevatorFloor > 1 {
			s := step(moveObjectDown(objects, i), elevatorFloor-1, stepCount+1, stepCountLimit, currentState)
			if s {
				return true
			}
		}
		if elevatorFloor < 4 {
			s := step(moveObjectUp(objects, i), elevatorFloor+1, stepCount+1, stepCountLimit, currentState)
			if s {
				return true
			}
		}
	}

	// Move elevator with two objects
	for i, obj1 := range objects {
		for j, obj2 := range objects {
			if i == j {
				//same object
				continue
			}
			if obj1.Floor != elevatorFloor || obj2.Floor != elevatorFloor {
				continue
			}
			if elevatorFloor > 1 {
				s := step(move2ObjectDown(objects, i, j), elevatorFloor-1, stepCount+1, stepCountLimit, currentState)
				if s {
					return true
				}
			}
			if elevatorFloor < 4 {
				s := step(move2ObjectUp(objects, i, j), elevatorFloor+1, stepCount+1, stepCountLimit, currentState)
				if s {
					return true
				}
			}
		}
	}
	return false
}

func isLegalMove(objects []Object, elevatorFloor, stepCount, stepCountLimit int) bool {
	return false
}

type State struct {
	ElevatorFloor int
	Pairs         []Pair
}
type Pair struct {
	GeneratorFloor int
	ElementFloor   int
}

type ByPair []Pair

func (a ByPair) Len() int      { return len(a) }
func (a ByPair) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByPair) Less(i, j int) bool {
	if a[i].GeneratorFloor == a[j].GeneratorFloor {
		return a[i].ElementFloor < a[j].ElementFloor
	}
	return a[i].GeneratorFloor < a[j].GeneratorFloor
}

var seenStates = []State{}

func genPairs(objects []Object) []Pair {
	pairs := []Pair{}
	for _, obj1 := range objects {
		for _, obj2 := range objects {
			if obj1.Type == 'G' && obj2.Type == 'M' && obj1.Element == obj2.Element {
				pairs = append(pairs, Pair{obj1.Floor, obj2.Floor})
			}
		}
	}
	sort.Sort(ByPair(pairs))
	return pairs
}

func updateState(state State) {
	seenStates = append(seenStates, state)
}

func seenBefore(stateA State) bool {
	for _, stateB := range seenStates {
		if stateEqual(stateA, stateB) {
			return true
		}
	}
	return false
}

func stateEqual(a, b State) bool {
	if a.ElevatorFloor != b.ElevatorFloor {
		return false
	}

	if len(a.Pairs) != len(b.Pairs) {
		return false
	}

	for i := range a.Pairs {
		if a.Pairs[i] != b.Pairs[i] {
			return false
		}
	}
	return true

}

// moveObject returns an array with one object modified
func moveObjectUp(objects []Object, idx int) []Object {
	objs := make([]Object, len(objects))
	copy(objs, objects)
	objs[idx].Floor++
	return objs
}

// moveObject returns an array with one object modified
func moveObjectDown(objects []Object, idx int) []Object {
	objs := make([]Object, len(objects))
	copy(objs, objects)
	objs[idx].Floor--
	return objs
}

// moveObject returns an array with one object modified
func move2ObjectUp(objects []Object, idx1, idx2 int) []Object {
	objs := make([]Object, len(objects))
	copy(objs, objects)
	objs[idx1].Floor++
	objs[idx2].Floor++
	return objs
}

// moveObject returns an array with one object modified
func move2ObjectDown(objects []Object, idx1, idx2 int) []Object {
	objs := make([]Object, len(objects))
	copy(objs, objects)
	objs[idx1].Floor--
	objs[idx2].Floor--
	return objs
}

func isSuccess(objects []Object) bool {
	success := true
	for _, obj := range objects {
		if obj.Floor != 4 {
			success = false
		}
	}
	return success
}

func isFailure(objects []Object, floor int) bool {
	microchipsOnFloor := []rune{}
	generatorsOnFloor := []rune{}

	for _, obj := range objects {
		if obj.Floor != floor {
			continue
		}
		if obj.Type == 'M' {
			microchipsOnFloor = append(microchipsOnFloor, obj.Element)
		}
		if obj.Type == 'G' {
			generatorsOnFloor = append(generatorsOnFloor, obj.Element)
		}
	}

	for _, microchip := range microchipsOnFloor {
		hasOwnGenerator := false
		hasDangerGenerator := false
		for _, generator := range generatorsOnFloor {
			if microchip == generator {
				hasOwnGenerator = true
			} else {
				hasDangerGenerator = true
			}
		}
		if !hasOwnGenerator && hasDangerGenerator {
			//fmt.Printf("Failed because RTG %c M:%q G:%q\n", microchip, microchipsOnFloor, generatorsOnFloor)
			return true
		}
	}
	return false
}

func printObjects(objects []Object, elevatorFloor int) {
	for i := 4; i > 0; i-- {
		fmt.Printf("F%d ", i)
		if elevatorFloor == i {
			fmt.Printf("E  ")
		} else {
			fmt.Printf(".  ")
		}
		for _, obj := range objects {
			if obj.Floor == i {
				fmt.Printf("%c%c ", obj.Element, obj.Type)
			} else {
				fmt.Printf(".  ")
			}
		}
		fmt.Printf("\n")
	}
}
