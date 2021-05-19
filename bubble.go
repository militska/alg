package main

import "fmt"

func sortBubble() []int8 {
	fmt.Print("hi! \n")
	x := []int8{3, 20, 5, 32, 1, 16, 98, 76, 2}

	sortArray := _sortBubble(x)
	//fmt.Print(sortArray)

	return sortArray
}

func _sortBubble(x []int8) []int8 {
	needRepeat := true
	for needRepeat {
		needRepeat = false
		for i, v := range x {
			prevKey := i
			if i != 0 {
				prevKey = i - 1
			}
			if v < x[prevKey] {
				needRepeat = true
				x[i], x[i-1] = x[i-1], v
			}

		}
	}
	return x
}
