package main

import "fmt"

//编写一个程序来计算斐波纳契数列
func main() {
	val := 0

	var numbers []int
	numbers = append(numbers, 1)
	numbers = append(numbers, 1)

	fmt.Print("Enter number: ")
	fmt.Scanf("%d", &val)
	if val < 2 {
		panic("number must be > 2")
	}

	for i := 2; i < val; i++ {
		numbers = append(numbers, numbers[i-1]+numbers[i-2])

	}
	fmt.Printf("%v\n", numbers)
}
