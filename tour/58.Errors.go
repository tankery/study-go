/*
 * http://tour.golang.org/#58
 * Exercise: Errors
 * 
 * Copy your Sqrt function from the earlier exercises
 * and modify it to return an error value.
 * 
 * Sqrt should return a non-nil error value when given
 * a negative number, as it doesn't support complex numbers.
 * 
 * Create a new type
 *      type ErrNegativeSqrt float64
 * and make it an error by giving it a
 *      func (e ErrNegativeSqrt) Error() string
 * method such that ErrNegativeSqrt(-2).Error() returns
 * "cannot Sqrt negative number: -2".
 * 
 * Note: a call to fmt.Print(e) inside the Error method
 * will send the program into an infinite loop.
 * You can avoid this by converting e first:
 * fmt.Print(float64(e)). Why?
 * 
 * Change your Sqrt function to return an ErrNegativeSqrt value
 * when given a negative number.
 */

package main

import (
	"fmt"
	"math"
)

// 定义自己的Error类型，并实现error接口的Error()函数
type ErrNagativeSqrt float64

func (e ErrNagativeSqrt) Error() string {
	return fmt.Sprintf(
		"cannot Sqrt negative number: %f",
		e)
}

// 平方根函数，跟练习的区别是加入了查错机制
func Sqrt(f float64) (float64, error) {
	// 只有输入为正才能工作，否则，返回一个错误
	if f > 0 {
		var z = f
		for {
			var oz = z
			z = z - (z*z -f)/(2*z)
			if math.Abs(oz - z) < 1e-7 {
				break
			}
		}
		return z, nil
	} else {
		var e = ErrNagativeSqrt(f)
		return 0, e
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

