package main

import "fmt"

//数组的key 是整形
//map的key 是字符串
func main() {
	studentsAge := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	fmt.Println(studentsAge)
}
