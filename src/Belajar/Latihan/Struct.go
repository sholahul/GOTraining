package Latihan

import (
	"fmt"
)

type respon struct {
	responCode string
	message    string
	data       data
}

type data struct {
	id    string
	nama  string
	email string
	role  string
}

type respon2 struct {
	responCode string
	message    string
	GetAllUser []map[string]string
}

type respon3 struct {
	username string
	password string
}

func Struct() {
	var R respon
	R.responCode = "00"
	R.message = "Success"

	var D data
	D.id = "123123123"
	D.nama = "Taca Rosa"
	D.email = "tacarosa21@gmail.com"
	D.role = "1"

	R.data = D

	// log.Println(R)
	fmt.Print(R)

	fmt.Println("\n\n =========== DEFINE STRUCT =========")
	var R3 = respon3{
		username: "Sholahul",
		password: "Fajri",
	}
	fmt.Println(R3)

	fmt.Println("\n\n =========== Struct and NMAP =========")

	var R2 respon2
	R2.responCode = "01"
	R2.message = "failed"
	R2.GetAllUser = []map[string]string{
		{"id": "124124124", "nama": "Albert Ivando"},
		{"id": "125125125", "nama": "Jonouchu"},
	}

	// log.Println(R2)
	fmt.Print(R2)
}
