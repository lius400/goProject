// 这个示例程序展示公开的结构类型中未公开的字段
// 无法直接访问
package main

import (
	"fmt"
	"temProject/entities"
)

// main 是应用程序的入口
func main() {
	u := entities.user{
		Name:  "Bill",
		Email: "Bill@email.com",
	}

	fmt.Printf("User: %v\n", u)
}
