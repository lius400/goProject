//这个示例程序展示如何将一个类型嵌入另一个类型，以及
//内部类型和外部类型之间的关系
package main

import (
	"fmt"
)

//user在程序里定义一个用户类型
type User struct {
	name  string
	email string
}

//notify实现了一个可以通过user类型值的指针
//调用的方法
func (u *User) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

//admin代表一个拥有权限的管理员用户
type admin struct {
	User  //嵌入类型
	level string
}

//main是应用程序的入口
func main() {
	//创建一个admin用户
	ad := admin{
		User: User{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	//我们可以直接访问内部类型的方法
	ad.User.notify()

	//内部类型的方法也被提升到外部类型
	ad.notify()
}
