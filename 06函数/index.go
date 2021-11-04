package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 100
	return c
}

//返回多个返回值，匿名的
func foo2(a string, b int) (int, int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	c := 100
	return c, c
}

//返回多个返回值，有形参
func foo3(a string, b int) (re1 int, re2 int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	re1, re2 = 100, 100
	return
}

func main() {
	c := foo1("abc", 777)

	fmt.Println(c)
	re1, re2 := foo2("adc", 123)
	fmt.Println(re1, re2)
}
