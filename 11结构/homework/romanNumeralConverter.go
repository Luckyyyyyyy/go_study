package main

import (
	"fmt"
	"strconv"
)

func main() {
	val := ""
	total := 0
	list := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}
	//临时映射

	tem := make(map[string]int)
	fmt.Print("Enter string: ")
	fmt.Scanf("%s", &val)
	//遍历字符串 判断是否存在
	for index, digit := range val {

		//转编码 遍历后为unicode编码
		value, exist := list[string(digit)]
		if !exist {
			fmt.Printf("%s value couldn't be found\n", string(digit))
			return
		}
		tem[strconv.Itoa(index)] = value
		//判断大小

		//计算和

		fmt.Println("value is", value)

		fmt.Println("total is", total)
	}
	fmt.Println(tem)
}
