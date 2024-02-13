package main

import "fmt"

func main() {
	/*format deklrasi
	var <nama-variabel> <tipe-data>
	var <nama-variabel> <tipe-data> = <nilai>
	*/

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

	//var boolean
	var bool1 bool = false
	bool2 := true

	//define string
	var message = `Nama saya adalah fajri.
	Bagian Pengembangan TSI.
	Pengelolaan core banking`

	//undefined variabel string blackhole bangke wkwk
	_ = "hayo"

	//pointer
	name := new(string)
	name2 := "jono"

	name = &name2

	fmt.Printf("halo %s %s %s!\n", firstName, middlename, lastName)
	fmt.Printf("halo %s %s %s!\n", awal, tengah, akhir)
	fmt.Printf("halo %s %s %s!\n", awal1, tengah1, akhir1)
	fmt.Printf("halo %d %f %s!\n", integ, float, str)
	// using %t untuk boolean dan printf, tidak dapat menggunakan Println
	fmt.Printf("terkadang %t dan %t \n", bool1, bool2)
	fmt.Println(message)
	//concat string menggunakan printf
	// fmt.Printf("%s", awal, +"ganteng")
	fmt.Println(name2, name)
	fmt.Printf("nama : %s \n", name2)
	// print pointer using *
	fmt.Printf("nama : %s", *name)

}
