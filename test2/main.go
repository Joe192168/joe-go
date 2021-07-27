package main

import "fmt"

var(
	name string
	age int
	isOk bool
)

func main() {
	name = "张三"
	age = 12
	isOk = true
	fmt.Printf("name:%s\n",name)
	fmt.Print(age)
	fmt.Println()
	fmt.Println(isOk)
}
