package main

import "fmt"


var a = 100 // 全局变量
var b int = 200
// 全局变量不支持简短定义
// c := 300

/*
变量的内存分析和注意
 */
func main()  {
	/*
	注意点：
	1. 变量必须先定义再使用
	2. 变量类型和赋值类型一致
	3. 同一作用域变量名不能冲突
	4. 简短定义方式，左边的变量至少有一个是新的
	5. 简短定义不支持全局变量
	6. 变量定义就要使用
	 */


	var num int

	num = 100
	fmt.Printf("num的数值是：%d, 地址是：%p\n", num, &num)

	num = 200
	fmt.Printf("num的数值是：%d, 地址是：%p\n", num, &num)

	// 定义的类型和赋值的类型必须保持一致
	var name string
	// name = 100 // 报错
	// fmt.Println(name)
	name = "zhanyeye"
	fmt.Println(name)

	// 简洁声明左边至少有一个新变量
	// 报错
	// num, name := 100, “zhanyeye"
	num, name, sex := 100, "zhanyeye", "man"
	fmt.Println(num, name, sex)

	var m int      // 0
	var k string   // ""
	var f float64  // 0.0
	var g []int    // []
	fmt.Printf("m: %d\n", m)
	fmt.Printf("k: %s\n", k)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(g == nil)

}