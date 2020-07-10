package main

import "fmt"

func main() {

	/*
	数据类型转换
	常量：在有需要时自动转换
	变量：需要手动转换
	 */
	var a int8
	a = 10

	var b int16
	b = int16(a)
	fmt.Println(a, b)

	f1 := 3.14
	var c int
	c = int(f1)
	fmt.Println(f1, c)

	f1 = float64(a)
	fmt.Println(f1, a)

	// 不可以将布尔类型转换成数值类型
	// b1 := true
	// a = int8(b1)  cannot convert b1 (type bool) to type int8

	sum := f1 + 100
	fmt.Printf("%T, %f\n", sum, sum)
}