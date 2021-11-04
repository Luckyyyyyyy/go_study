package main

/*
查找质数
编写一个程序来查找小于 20 的所有质数。 质数是大于 1 的任意数字，
只能被它自己和 1 整除。 “整除”表示经过除法运算后没有余数。
与大多数编程语言一样，Go 还提供了一种方法来检查除法运算是否产生余数。
我们可以使用模数 %（百分号）运算符。

在本练习中，你将更新一个名为 findprimes 的函数，以检查数值是否为质数。
该函数有一个整数参数，并返回一个布尔值。 函数通过检查是否有余数来测试输入数字是否为质数。
如果数字为质数，则该函数返回 true。
*/
import "fmt"

func findprimes(number int) bool {
	for i := 2; i < number; i++ {
		//fmt.Printf("%v ", number%i)
		if number%i == 0 {
			return false
		}
	}
	if number > 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("Prime numbers less than 20:")

	for number := 1; number <= 20; number++ {
		//fmt.Printf("%v ", number)
		if findprimes(number) {
			fmt.Printf("%v ", number)
		}
	}
}
