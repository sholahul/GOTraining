package Latihan

import "fmt"

func Array() {
	var names = [4]string{"apple", "pie", "banana", "grepe"}
	var names2 [3]string
	names2[0] = "Albert"
	names2[1] = "Handika"
	var names3 = [2]string{
		"Yuharis",
		"Taca Rosa",
	}

	names4 := [2]string{"Rio", "Febby"}
	var angka = [...]int{1, 2, 3, 4}
	var angka2 = [5]int{
		1: 10,
		3: 30,
	}
	var angka3 = [3]int{
		1: 30,
		2: 60,
	}
	//Matrix
	var matrix = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	//make
	var fruit = make([]string, 2)
	fruit[0] = "cherry"
	fruit[1] = "belle"

	fmt.Println("Panjang array name1 : ", len(names))

	for i := 0; i < len(names); i++ {
		fmt.Println("Array Ke -", i, ": ", names[i])
	}

	fmt.Println("\nPanjang array name2 : ", len(names2))
	for i := 0; i < len(names2); i++ {
		fmt.Println("Nama Ke -", i, ": ", names2[i])
	}

	fmt.Println("\nPanjang array name3 : ", len(names3))
	for i := 0; i < len(names3); i++ {
		fmt.Println("Nama Ke -", i, ": ", names3[i])
	}

	fmt.Println("\nPanjang array name3 : ", len(names4))
	for i := 0; i < len(names4); i++ {
		fmt.Println("Nama Ke -", i, ": ", names4[i])
	}

	fmt.Println("\nPanjang array angka : ", len(angka))
	for index, value := range angka {
		fmt.Println("Angka Ke-", index, ": ", value)
	}

	fmt.Println("\nPanjang array angka2 : ", len(angka2))
	for index2, value2 := range angka2 {
		fmt.Println("Angka Ke-", index2, ": ", value2)
	}

	//underscore (pengangguran wkwk)
	fmt.Println("\nPanjang array angka3 : ", len(angka3))
	for _, value3 := range angka3 {
		fmt.Println("Angka : ", value3)
	}

	fmt.Printf("\nMatrix :")
	fmt.Println(matrix)
	fmt.Println()

	fmt.Println("\nPanjang array fruit using 'make' : ", len(fruit))
	for _, value4 := range fruit {
		fmt.Println("Angka : ", value4)
	}
}
