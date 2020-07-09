package main

import "fmt"

func main() {

	// 1. 定义一组常量
	const PATH string = "www.baidu.com"
	const PI = 3.14
	fmt.Println(PATH)

	// 2. PATH = "www.github.com" //cannot assign to PATH

	// 3. 定义一组常量
	const C1, C2, C3 = 100, 3.12, "哈哈"
	const (
		MALE = 0
		FEMALE = 1
		UNKNOW = 3
	)

	// 4. 一组常量中，如果某个常量没有初始值，默认和上一行一致
	const (
		a int = 100
		b
		c string = "zhanyeye"
		d
		e
	)
	fmt.Printf("%T, %d \n", a, a)
	fmt.Printf("%T, %d \n", b, b)
	fmt.Printf("%T, %s \n", c, c)
	fmt.Printf("%T, %s \n", d, d)
	fmt.Printf("%T, %s \n", e, e)

	// 5. 枚举类型：使用常量组作为枚举类型
	const (
		SPRING = 0
		SUMMER = 1
		AUTUTMN
	)

}
