package Latihan

import "fmt"

func Selec() {
	var nilai int

	fmt.Printf("Masukkan Nilai : ")
	_, err := fmt.Scanf("%d", &nilai)

	if err != nil {
		fmt.Println("Error Reading Nilai :", err)
		return
	}

	fmt.Printf("Nilai anda : %d \n", nilai)

	//using if
	if nilai > 80 {
		fmt.Printf("Nilai Anda A")
	} else if nilai > 70 {
		fmt.Printf("Nilai Anda B")
	} else if nilai > 50 {
		fmt.Printf("Nilai Anda C")
	} else if nilai > 40 {
		fmt.Printf("Nilai Anda D")
	} else {
		fmt.Printf("Nilai Anda E")
	}

	switch {
	case nilai == 80:
		fmt.Printf("Lulus dengan Nilai Anda A")
	case nilai > 70:
		fmt.Printf("Lulus dengan Nilai Anda B")
	default:
		{
			fmt.Printf("\nTidak Lulus")
		}
	}
}
