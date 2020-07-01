package main

import "fmt"

//变量声明后必须使用

//全局变量区域
//单个声明
var a string
var b int
var c bool

//批量声明
var (
	d string
	e int
	f float32
)

//声明时初始化
var g int = 10

//变量类型推导
var h = "Nidhogg" //Go自动推导h的类型为string

//一次初始化多个变量
var i, j = 10, "Nidhogg"

func main() {

	// 局部变量声明区域，只能作用于本函数内
	//短变量声明，只能在函数内使用。其他声明方式也可在函数内使用
	k := true
	l := 3.14159

	//匿名变量“_”,在使用多重赋值时，如果想要忽略某个值，可以使用。
	m, _ := 0, 1
	//对已声明的变量赋值
	a = "测试"
	d = "开发"

	fmt.Println(a, b, c) //println:输入结束时自动换行
	fmt.Println(d, e, f)
	fmt.Println(g, h, i, j)
	fmt.Println(k, l, m)
}
