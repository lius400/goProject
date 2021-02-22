// 这个示例程序展示无法从另一个包里
// 访问未公开的标识符
package main

import (
	"fmt"
	"temProject/counters"
)

// main 是应用程序的入口
func main() {
	// 使用 counters 包公开的 New 函数来创建
	// 一个未公开的类型的变量
	counter := counters.New(10)
	Counter := counters.AlertCounter(11)

	// ./listing.go:: 不能引用未公开的名字
	// counters.alertCounter
	// ./listing.go:: 未定义：counters.alertCounter

	fmt.Printf("Counter: %d\n", counter)
	fmt.Printf("Public Counter: %d\n", Counter)
}
