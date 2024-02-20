package main

import "fmt"

func main() {
	fmt.Println("SLICE *******")

	//define slice tanpa len
	var fruits = []string{"apple", "grape", "banana", "melon", "semangka"}
	var buahbesar = fruits[3] + fruits[4]

	var newfruits = fruits[0:2]
	var newfruits2 = fruits[:]
	var newfruits3 = fruits[2:]
	var newfruits4 = fruits[:2]

	fmt.Println("\nSlice : ", fruits[0]) // "apple"
	fmt.Println("Buah Besar : ", buahbesar)
	fmt.Println("newfruits : ", newfruits)
	fmt.Println("newfruits2 : ", newfruits2)
	fmt.Println("newfruits3 : ", newfruits3)
	fmt.Println("newfruits4 : ", newfruits4)

	// Perubahan mempengaruhi memori lama
	newfruits2[1] = "grrapppee"
	fmt.Println("\n\nnewfruits : ", newfruits)
	fmt.Println("newfruits2 : ", newfruits2)
	fmt.Println("newfruits3 : ", newfruits3)
	fmt.Println("newfruits4 : ", newfruits4)

	//fungsi len
	fmt.Println("\n\nFungsi Len : ", len(newfruits2))

	//fungsi cap
	fmt.Println("Fungsi Cap : ", cap(newfruits2))

	//jika terdapat operasi slice, cap lebihterbaru
	anewfruits2 := newfruits2[0:4]

	fmt.Println("anewfruits2 :", anewfruits2)
	fmt.Println("\n\nFungsi Len anewfruits2 : ", len(anewfruits2))
	fmt.Println("Fungsi Cap anewfruits2 : ", cap(anewfruits2))
}
