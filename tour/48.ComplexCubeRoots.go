/*
 * http://tour.golang.org/#48
 * Advanced Exercise: Complex cube roots
 * 
 * Let's explore Go's built-in support for complex numbers
 * via the complex64 and complex128 types.
 * For cube roots, Newton's method amounts to repeating:
 * 	z = z - (z^3 - x) / (3*z^2)
 * Find the cube root of 2,
 * just to make sure the algorithm works.
 * There is a Pow function in the math/cmplx package.
 */
package main

import (
	"fmt"
	"math/cmplx"
)

// 利用牛顿法计算立方根，支持复数运算
func Cbrt(x complex128) complex128 {
	var z = x
	for {
		var oz = z
		z = z - ((z*z*z - x)/ (3*z*z))
		if cmplx.Abs(oz - z) < 1e-7 {
			break
		}
	}
	return z
}

// 支持复数的牛顿法平方根计算
func Sqrt(x complex128) complex128 {
	var z = x
	for {
		var oz = z
		z = z - ((z*z - x)/ (2*z))
		if cmplx.Abs(oz - z) < 1e-7 {
			break
		}
	}
	return z
}

func main() {
	x := -1 + 0i
	fmt.Println(Cbrt(x))
	fmt.Println(Sqrt(x))
}

