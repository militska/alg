package main

import "fmt"

func main() {
	fmt.Print("hi! \n")
	x := []int8{3, 20, 5, 32, 1, 16, 98, 76, 2}
	sortArray := sort(x)
	fmt.Print(sortArray)
}

func sort(x []int8) []int8 {
	var needRec bool = false
	var iteration int8 = 0
	for i, v := range x {
		current := x[i]
		preview := x[i]
		if i != 0 {
			preview = x[i-1]
			if current > preview {
				needRec = true
				iteration += 1
				x[i] = x[i-1]
				x[i-1] = v
			}
		}
	}
	if needRec {
		sort(x)
	}
	return x
}
