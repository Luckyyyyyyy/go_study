package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string
	Address   string `json:"address,omitempty"`
}

type Employee struct {
	Person
	ManagerID int
}

type Contractor struct {
	Person
	CompanyID int
}

//使用 & 运算符生成指向结构的指针，如以下代码所示：
func main() {

	//结构初始化
	person := Person{LastName: "Doe", FirstName: "John"}
	fmt.Println(person)
	employeeCopy := &person
	employeeCopy.FirstName = "David"
	fmt.Println(person)

	//结构嵌入
	employee := Employee{
		Person: Person{
			FirstName: "John",
		},
	}
	employee.LastName = "Doe"
	fmt.Println(employee.FirstName)

	//最后，可使用结构来对 JSON 中的数据进行编码和解码。
	//Go 对 JSON 格式提供很好的支持，该格式已包含在标准库包中。

	//若要将结构编码为 JSON，请使用 json.Marshal 函数。
	//若要将 JSON 字符串解码为数据结构，
	//请使用 json.Unmarshal 函数。
	employees := []Employee{
		Employee{
			Person: Person{
				LastName: "Doe", FirstName: "John",
			},
		},
		Employee{
			Person: Person{
				LastName: "Campbell", FirstName: "David",
			},
		},
	}

	data, _ := json.Marshal(employees)
	fmt.Printf("%s\n", data)

	var decoded []Employee
	json.Unmarshal(data, &decoded)
	fmt.Printf("%v", decoded)
}
