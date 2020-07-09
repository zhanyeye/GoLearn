package main

import "fmt"
/*
变量概念和使用
 */
func main()  {
	fmt.Println("hello world")
	fmt.Println("GO GO GO")

	// 先定义，后赋值
	var num1 int
	num1 = 30
	fmt.Printf("num1的数值是: %d\n", num1)

	// 定义并赋值
	var num2 int = 15
	fmt.Printf("num2的数值是: %d\n", num2)

	// 类型推断
	var name = "zhanyeye"
	fmt.Printf("类型是：%T, 数值是：%s\n", name, name)

	// 简短声明, 省略 var
	sum := 100
	fmt.Println(sum)

	// 多个变量同时定义
	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println(a, b, c)

	// 定义多个变量并赋值
	var m, n int = 100, 200
	fmt.Println(m, n)

	// 定义不同类型变量并赋值
	var n1, f1, s1 = 10, 3.14, "GO"
	fmt.Println(n1, f1, s1)

	// 定义一组值
	var (
		stuName = "zhanyeye"
		age = 18
		sex = "man"
	)
	fmt.Printf("学生姓名：%s, 年龄：%d, 性别：%s", stuName, age, sex)


}
