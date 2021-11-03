package main

import "fmt"

/*
如果你声明了一个局部变量却没有在相同的代码块中使用它，同样会得到编译错误
但是全局变量是允许声明但不使用的。
*/

func main() {

	//变量声明
	//第一种，指定变量类型，如果没有初始化，则变量默认为零值。
	var a = "RUNOOB"
	fmt.Println(a)

	// 没有初始化就为零值
	var b int
	fmt.Println(b)

	// bool 零值为 false
	var c bool
	fmt.Println(c)

	//第二种，根据值自行判定变量类型。
	var d = true
	fmt.Println(d)
	/*
		第三种，  := 声明变量，格式：
		这种不带声明格式的只能在函数体中出现
		intVal := 1
		相等于：
		var intVal int
		intVal =1
	*/
	intVal := 1
	fmt.Println(intVal)

	/*
		_ 实际上是一个只写变量，你不能得到它的值。
	*/
	_, b = 5, 7

	//显式类型定义：
	const e string = "abc"
	//隐式类型定义：
	const f = "abc"

}
