package main

import "fmt"

func main() {
	/**
	算术运算符： +, -, *, /, %, ++, --
	 */
	a := 10
	b := 3
	sum := a + b
	fmt.Printf("%d + %d = %d \n", a, b, sum)

	sub := a - b
	fmt.Printf("%d - %d = %d \n", a, b, sub)

	mul := a * b
	fmt.Printf("%d * %d = %d \n", a, b, mul)

	div := a / b
	mod := a % b
	fmt.Printf("%d / %d = %d \n", a, b, div)
	fmt.Printf("%d %% %d = %d \n", a, b, mod)

	c := 3
	c++
	fmt.Println(c)
	c--
	fmt.Println(c)
	// go 语言中 ++， --  不支持参加运算

	/**
	关系运算符： > , < , >= , <= , == , !=
	 */


	/*
	逻辑运算符： 操作数必须时bool,运算结果也是bool
	&&, ||, !
	 */

	/**
	位运算符：
	按位&： 对应位置上都为1，才为1，有一个为0，就为0
	按位|： 对应位置上都是0，才为0，有一个为一1，就为1
	异或^:
		二元： a^b
			对应位值不同为1，相同为0
		一元：^a
			按位取反  有些语言用~来表示
	位清空：&^
		对于 a &^ b
		对于b上的每个数值，如果为0，取对应位上的数值， 如果为1，结果位就取0
	 */
	a1 := 60
	b1 := 13
	fmt.Printf("%d, %b\n", a1, a1)
	fmt.Printf("%d, %b\n", b1, b1)
	fmt.Printf("%b & %b = %b\n", a1, b1, a1 & b1)
	fmt.Printf("%b | %b = %b\n", a1, b1, a1 | b1)
	fmt.Printf("%b ^ %b = %b\n", a1, b1, a1 ^ b1)
	fmt.Printf("%b &^ %b = %b\n", a1, b1, a1 &^ b1)
	fmt.Printf("^%b = %b, %d\n", b1, ^b1, ^b1)

	/**
	位移运算符：
		<< 按位左移
		>> 按位右移
	 */


	/*
	赋值运算符
	=, +=, -=, *=, /=, %=, <<=, >>=, &=, |=, ^= ...
	 */


	/*
	GO运算符优先级
	7 ! ++ -- ~(^单元)
	6 * / % << >> & &^
	5 + - ^
	4 == !=  < <= >= >
	3 <-
	2 &&
	1 ||
	 */

}