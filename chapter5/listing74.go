// 这个示例程序展示公开的结构类型中如何访问
// 未公开的内嵌类型的例子
package main

import (
	"fmt"
	"temProject/entities"
)

// main 是应用程序的入口
func main() {

	a := entities.Admin{
		Rights: 10,
	}

	// 设置未公开的内部类型的
	// 公开字段的值
	a.Name = "Bill"
	a.Email = "Bill@email.com"

	fmt.Printf("User: %v\n", a)
}
