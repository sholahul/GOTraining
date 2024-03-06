package day1

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func Interface() {
	var a Abser
	f := Myfloat(-math.Sqrt2)
	fmt.Println("F : ", f.Abs())

	v := vertex{3, 4}
	fmt.Println("v : ", v.Abs())

	a = f
	fmt.Println("a = f  : ", a.Abs())

	a = &v
	fmt.Println("a = &v  :", a.Abs())

	a = v
	fmt.Println("a = v :", a.Abs())
}
