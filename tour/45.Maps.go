/*
 * http://tour.golang.org/#45
 * Exercise: Maps
 * 
 * Implement WordCount.
 * It should return a map of the counts of each "word"
 * in the string s. The wc.Test function runs a test suite
 * against the provided function and prints success or failure.
 * 
 * You might find strings.Fields helpful.
 */
// 类似Java，用包名来组织代码
package main

import (
	"tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	// 创建一个键为string，值为int的map
	// make可以用来创建任何类型的变量。
	// 比如make([]int, 3)是创建3个元素的int数组
	m := make(map[string]int)
	// 变量的使用可以不用显示的指明类型
	// 这里，words的类型即Fields的返回值类型，是个字符串数组
	words := strings.Fields(s)

	// Go语言没有while、do-while
	// for 条件 { 执行体 } 即相当于while
	// for { 执行体 } 即无限循环
	// 这里，使用for的range特性，取words的索引和值
	// 分别给_和word，下划线_相当于一个占位符，不赋值给具体的变量
	// 同样，还可以使用：i, _ := range words，表示只需要其索引
	// 甚至可以使用：_, _ := range words，表示只需要循环相应次数即可
	for _, word := range words {
		// 根据键取map中的值，并修改
		// Go是内存安全的语言，如果m中不存在word键
		// 将会自动创建一个word，并初始化其值为0
		m[word]++;
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

