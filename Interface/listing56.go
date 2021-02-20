// 这个示例程序展示如何将嵌入类型应用于接口
package main

import (
	"fmt"
)

// notifier 是一个定义了
// 通知类行为的接口
type Notifier interface {
	notify()
}

// Reuser 在程序里定义一个用户类型
type Reuser struct {
	name  string
	email string
}

// 通过 Reuser 类型值的指针
// 调用的方法
func (u *Reuser) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// Admin 代表一个拥有权限的管理员用户
type Admin struct {
	Reuser
	level string
}

// main 是应用程序的入口
func main() {
	// 创建一个 admin 用户
	ad := Admin{
		Reuser: Reuser{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// 给 Admin 用户发送一个通知
	// 用于实现接口的内部类型的方法，被提升到
	// 外部类型
	SendNotification(&ad)
}

// sendNotification 接受一个实现了 Notifier 接口的值
// 并发送通知
func SendNotification(n Notifier) {
	n.notify()
}
