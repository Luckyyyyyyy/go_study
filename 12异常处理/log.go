package main

import (
	"log"
	"os"
)

func main() {

	//记录到文件
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//关闭文件
	defer file.Close()

	log.SetOutput(file)
	log.Print("Hey, I'm a log!")

	//可以使用 log.Fatal() 函数记录错误并结束程序，
	//就像使用 os.Exit(1) 一样。
	log.Fatal("Hey, I'm an error log!")

}
