package day1

type Myfloat float64

func (f Myfloat) Abs() float64 {
	return float64(f * -1)
}

func (f Myfloat) Pangkatdua() float64 {
	return float64(f * 2)
}
