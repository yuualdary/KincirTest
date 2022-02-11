package main

import (
	"fmt"
	"strconv"
)

func main() {

	//CountSame := make(map[int]int)

	var arr1 = map[int]int{1: 0, 2: 1, 3: 2, 4: 3, 5: 4}
	var arr2 = []int{2, 4, 6}

	hasil := []string{}

	for _, check := range arr2 {
		var x, isExist = arr1[check]
		if isExist {
			hasil = append(hasil, strconv.Itoa(check)+" di index ke : "+strconv.Itoa(x)+",")
		}
	}

	fmt.Println(hasil)

}
