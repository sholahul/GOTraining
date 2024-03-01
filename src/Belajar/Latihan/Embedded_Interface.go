package Latihan

import (
	"fmt"
	"math"
)

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

// embedd interface
type hitungd interface {
	hitung2d
	hitung3d
}

// kubus USING POINTER
type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

func Embedded_Interface() {
	var bangunRuang2 hitungd = &kubus{4}

	fmt.Println("===== kubus")
	fmt.Println("luas      :", bangunRuang2.luas())
	fmt.Println("keliling  :", bangunRuang2.keliling())
	fmt.Println("volume    :", bangunRuang2.volume())
}
