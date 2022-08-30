package main

import "fmt"

func main() {
	fmt.Println(myPow(2,10))
	fmt.Println(4&1)
}

func myPow(x float64, n int) float64 {
	var pow func(float64, int) float64
	pow = func(x float64, n int) float64 {
		if n == 0 {
			return 1
		}
		y := pow(x, n/2)
		// n 为偶数，可以分解完全
		if n&1 == 0 {
			return y * y
		}
		// n 为奇数，多出一个x
		return y * y * x
	}
	if n >= 0 {
		return pow(x, n)
	}
	return 1.0 / pow(x, -n)
}

