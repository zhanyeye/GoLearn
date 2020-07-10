package main

import "fmt"

func main() {
	/*
	字符串：
	1. 概念： 多个byte的集合，理解为一个字符序列
	2. 语法：使用双引号, 也可以是 ``
		"abc"
	 */

	// 定义字符串
	var s1 string
	s1 = "zhanyeye"
	fmt.Printf("%T, %s \n", s1, s1)

	s2 := `Hello world`
	fmt.Printf("%T, %s \n", s2, s2)

	// "" 和 '' 的区别
	v1 := 'A'
	v2 := "A"

	fmt.Printf("%T, %d \n", v1, v1)
	fmt.Printf("%T, %s \n", v2, v2)

	v3 := '中'
	fmt.Printf("%T, %d, %c \n", v3, v3, v3)   // %T 类型， %d 整形值， %c 字符

	// 转义字符
	fmt.Println("\"hello world\"")
	fmt.Println("hello\nwor \tld")
	fmt.Println(`hello "world"`)

}