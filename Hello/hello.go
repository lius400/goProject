package main

import (
	"fmt"
	//"example/greetings"
	"rsc.io/quote"
)

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
	myColors1 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	fmt.Println(myColors1["Red"])
	for key, value := range myColors1 {
		fmt.Println("Color: ", key, value)
	}
}
