/*
 * http://tour.golang.org/#47
 * Exercise: Fibonacci closure
 * 
 * Let's have some fun with functions.
 * 
 * Implement a fibonacci function that returns a function (a closure)
 * that returns successive fibonacci numbers.
 */
package main

import "fmt"

// fibonacci 返回一个返回值为int的函数指针
// 利用闭包的特性，使得闭包内变量值可以保存，以待下次调用时叠加
func fibonacci() func() int {
	var i1, i2 = 1, 1
	return func() int {
		var f = i1 + i2
		i1 = i2
		i2 = f
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

