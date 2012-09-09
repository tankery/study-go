/*
 * http://tour.golang.org/#69
 * Exercise: Equivalent Binary Trees
 * 
 * 1. Implement the Walk function.
 * 
 * 2. Test the Walk function.
 * 
 * The function tree.New(k) constructs a randomly-structured
 * binary tree holding the values k, 2k, 3k, ..., 10k.
 * 
 * Create a new channel ch and kick off the walker:
 * go Walk(tree.New(1), ch)
 * 
 * Then read and print 10 values from the channel.
 * It should be the numbers 1, 2, 3, ..., 10.
 * 
 * 3. Implement the Same function using Walk to determine
 * whether t1 and t2 store the same values.
 * 
 * 4. Test the Same function.
 * 
 * Same(tree.New(1), tree.New(1)) should return true,
 * and Same(tree.New(1), tree.New(2)) should return false.
 */
package main

import (
	"fmt"
	"tour/tree"
	"sort"
)

// Walk 遍历t，将其所有的内容由ch发送出去
// 我使用递归的方式实现了它
// 注意，Go语言的channel是有类型的
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	if t.Right != nil {
		Walk(t.Right, ch)
	}
	// 使用 <- 符号将变量值发送到chan
	ch <- t.Value
}

// Same 决定t1、t2是否是具有相同内容的两棵树
func Same(t1, t2 *tree.Tree) bool {
	// 创建两个具有缓存的管道
	// 管理发送的线程将不停的发送，直至缓存溢出，
	// 等到管理接收的线程取出值以后，才能继续发送
	// 这相当于一个固定大小的队列。
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	// 使用两个数组来存储树的内容	
	a1  := make([]int, 10)
	a2  := make([]int, 10)
	
	// 新建两个线程，同时开始遍历
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	// 主线程将不停的接收另外两个线程传来的数据
	for i := 0; i < 10; i++ {
		a1[i] = <- ch1
		a2[i] = <- ch2
	}
	
	// 排序以便于检查其内容是否一致
	sort.Ints(a1)
	sort.Ints(a2)
	for i := 0; i < 10; i++ {
		if a1[i] != a2[i] {
			return false
		}
	}
	
	return true
}

func main() {
	// tree.New(k)可以创建内容包含k, 2k, 3k, ..., nk的树
	t1 := tree.New(1)
	t2 := tree.New(2)
	ch := make(chan int, 10)
	
	// 打印t1树
	fmt.Print("t1: ")
	go Walk(t1, ch)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2d, ", <- ch)
	}
	fmt.Println()
	
	// 打印t2树
	fmt.Print("t2: ")
	go Walk(t2, ch)
	for i := 0; i < 10; i++ {
		fmt.Printf("%2d, ", <- ch)
	}
	fmt.Println()
	
	// 测试Same函数
	fmt.Printf("Is %v that t1 equal t1.\n", Same(t1, t1))
	fmt.Printf("Is %v that t1 equal t2.\n", Same(t1, t2))
}

