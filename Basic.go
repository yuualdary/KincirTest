// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

func main() {

	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{2, 4, 6}

	hasil := []string{}

	for i := 0; i < len(arr2); i++ {

		for j := 0; j < len(arr1); j++ {

			if arr2[i] == arr1[j] {

				hasil = append(hasil, "angka "+strconv.Itoa(arr2[i])+" di index ke "+strconv.Itoa(j)+", ")
			}
		}

	}

	fmt.Println(hasil)

}
