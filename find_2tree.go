package main

import "fmt"

func find2Tree(list []int8, num int8) {
	find := true
	inc := 0
	for find {
		inc++
		cnt := len(list) / 2

		if list[cnt] == num {
			fmt.Printf("Найдено за %d прохода \n", inc)
			find = false
			break
		}

		if num < list[cnt] {
			list = list[0 : cnt+1]
		} else {
			list = list[cnt-1 : len(list)]
		}

		find = true
	}
}
