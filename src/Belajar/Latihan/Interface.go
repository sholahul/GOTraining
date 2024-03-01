package Latihan

import (
	"fmt"
	"math"
)

// Define interface
type hitung interface {
	luas() float64
	keliling() float64
}

// ============================================= Lingkaran
type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

// =============================================Persegi
type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

func Interface() {
	var bangunDatar hitung

	bangunDatar = persegi{10.0} //proses init
	fmt.Println("===== persegi")
	fmt.Println("Sisi : ", bangunDatar.(persegi).sisi)
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0} //proses init
	fmt.Println("===== lingkaran")
	fmt.Println("Diameter : ", bangunDatar.(lingkaran).diameter)
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())
	fmt.Println("jari-jari :", bangunDatar.(lingkaran).jariJari())

	// cara 2
	// var bangunDatar2 hitung = lingkaran{14.0}
	// var bangunLingkaran2 lingkaran = bangunDatar.(lingkaran)

	// fmt.Println("jari-jari :",bangunLingkaran2.jariJari())
	// bangunLingkaran.jariJari()
}
