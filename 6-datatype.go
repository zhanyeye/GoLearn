package main

import "fmt"

func main() {
	/*
	GO 语言基本数据类型：
	布尔类型： bool
	数值类型：
		整数：int
		有符号：最高位表示符号位，0正数，1负数，其余位表示数值
			int8, int16, int32, int64
		无符号：所有位数表示数值
			uint8, uint16, uint32, uint64
		byte:unit8
		rune:int32
	 */

	// 布尔类型
	var b1 bool
	b1 = true
	fmt.Printf("%T, %t\n", b1, b1)
	b2 := false
	fmt.Printf("%T, %t\n", b2, b2,)
	// b2 = 100 // 报错

	// 整数
	var i1 int8   // -128 ~ 127
	i1 = -128
	fmt.Println(i1)
	var i2 uint8
	i2 = 255      // 0 ~ 255
	fmt.Println(i2)

	// 语法角度，int, int64 不是同一种类型
	var i3 int
	i3 = 1000
	fmt.Println(i3)
	// var i4 int64
	// i4 = i3

	// uint8 -> byte
	var i5 uint8
	i5 = 100
	var i6 byte
	i6 = i5
	fmt.Println(i5, i6)

	var f1 float32
	f1 = 3.14
	var f2 float64
	f2 = 4.78
	fmt.Printf("%T, %.2f\n", f1, f1)
	fmt.Printf("%T, %.3f\n", f2, f2)
	fmt.Println(f1)

	var f3 = 9.23
	fmt.Printf("%T, %f", f3, f3)  // 默认是float64






}
