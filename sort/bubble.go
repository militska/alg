package main

import "fmt"

func main() {
	fmt.Print("hi! \n")
	x := []int8{3, 20, 5, 32, 1, 16, 98, 76, 2}
	var iteration int = 0

	sortArray := sortBubble(x, iteration)
	fmt.Print(sortArray)
}

func sortBubble(x []int8, iteration int) []int8 {
	for {
		needRepeat := false
		internalIter := 0
		for i, v := range x {
			preview := x[i]
			if i != 0 {
				preview = x[i-1]
				if v > preview {
					internalIter += 1
					needRepeat = true
					x[i], x[i-1] = x[i-1], v
				}
			}
		}
		if internalIter != 0 {
			iteration += 1
		}
		fmt.Printf("проход : %d \n", iteration)

		if needRepeat == true {
			continue
		}
		break
	}

	fmt.Printf("Количество проходов : %d \n", iteration)
	return x
}
