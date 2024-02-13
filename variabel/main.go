package main

import "fmt"

func main() {
	//deklrasi string
	var firstName string = "john"
	var lastName string
	lastName = "wick"

	// deklrasi tanpa var
	middlename := "pantau"

	//deklrasi multi var
	var awal, tengah, akhir string
	awal = "sholahul"
	tengah = "fajri"
	akhir = "tsi"

	// multiple var
	awal1, tengah1, akhir1 := "Senang", "Berpetualang", "Jungler"

	//multiple var string, integer, float
	integ, float, str := 1, 1.1, "stringv"

	fmt.Printf("halo %s %s %s!\n", firstName, middlename, lastName)
	fmt.Printf("halo %s %s %s!\n", awal, tengah, akhir)
	fmt.Printf("halo %s %s %s!\n", awal1, tengah1, akhir1)
	fmt.Printf("halo %d %f %s!\n", integ, float, str)
}
