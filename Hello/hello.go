package main

import (
	"fmt"
	"rsc.io/quote"
	"temProject/example"
)

type User example.User

var helloString string
var attack int = 40
var defence int = 20
var damageRate float32 = 0.17
var damage = float32(attack-defence) * damageRate

func Hello() string {
	helloString := "Hello, World!"
	return helloString
}

func main() {

	fmt.Println(Hello())
	fmt.Println(helloString)
	fmt.Println(quote.Go())
	fmt.Println(quote.Opt())
	//var damageString = strconv.FormatFloat(damage, 'f', -1, 32)
	fmt.Printf("damage value: %G\n", damage)

	point1 := new(int)
	*point1 = 111
	fmt.Println("point1 value:", *point1)

	// 创建一个映射，键的类型是 string，值的类型是 int
	myColors := make(map[string]string)
	myColors["Blue"] = "#0000FF"
	// 创建一个映射，键和值的类型都是 string
	// 使用两个键值对初始化映射
	myColors1 := map[string]string{"Red": "#da1337", "Orange": "#e95a22", "AliceBlue": "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22"}
	fmt.Println(myColors1["Red"])
	for key, value := range myColors1 {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

	// 获取键 Blue 对应的值
	colorValue, exists := myColors1["Blue"]
	// 这个键存在吗？
	if !exists {
		fmt.Println("color Blue not in myColors1", colorValue)
	}

	// 删除键为 ForestGreen 的键值对
	removeColor(myColors1, "ForestGreen")
	fmt.Println("删除 ForestGreen 后")
	for key, value := range myColors1 {
		printLn, _ := fmt.Printf("Key: %s Value: %s\n", key, value)
		fmt.Println(printLn)
	}

	// user 类型的值可以用来调用
	// 使用值接收者声明的方法
	bill := User{Name: "Bill", Email: "lius400@163.com", Age: 33}
	bill.notify()

	// 指向 user 类型值的指针也可以用来调用
	// 使用值接收者声明的方法
	lisa := &User{Name: "Lisa", Email: "lisa@email.com", Age: 32, Privileged: true}
	lisa.notify()

	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}

func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}

// notify 使用值接收者实现了一个方法
func (u User) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n\n", u.Name, u.Email)
}

// changeEmail 使用指针接收者实现了一个方法
func (u *User) changeEmail(email string) {
	u.Email = email
}
