package main

import "fmt"

func main() {
	// 切片是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。
	// 切片是一个引用类型，它的内部结构包含地址、长度和容量。切片一般用于快速地操作一块数据集合。
	// 切片的定义
	fmt.Println("切片的定义：")
	slice1()
	// 切片拥有自己的长度和容量，我们可以通过使用内置的len()函数求长度，使用内置的cap()函数求切片的容量。
	// 切片表达式从字符串、数组、指向数组或切片的指针构造子字符串或切片。
	// 它有两种变体：一种指定low和high两个索引界限值的简单的形式，另一种是除了low和high索引界限值外还指定容量的完整的形式。

	//简单切片表达式
	fmt.Println("简单切片表达式：")
	slice2()
	// 对于数组或字符串，如果0 <= low <= high <= len(a)，则索引合法，否则就会索引越界（out of range）。
	// 对切片再执行切片表达式时（切片再切片），high的上限边界是切片的容量cap(a)，而不是长度。
	// 常量索引必须是非负的，并且可以用int类型的值表示;对于数组或常量字符串，常量索引也必须在有效范围内。
	// 如果low和high两个指标都是常数，它们必须满足low <= high。如果索引在运行时超出范围，就会发生运行时panic。
	fmt.Println("切片再切片：")
	slice3()

}

func slice1() {
	// 声明切片类型
	var a []string              //声明一个字符串切片
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化
	// var d = []bool{false, true} //声明一个布尔切片并初始化
	fmt.Println(a)        //[]
	fmt.Println(b)        //[]
	fmt.Println(c)        //[false true]
	fmt.Println(a == nil) //true
	fmt.Println(b == nil) //false
	fmt.Println(c == nil) //false
	// fmt.Println(c == d)   //切片是引用类型，不支持直接比较，只能和nil比较
}

func slice2() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]                                                 // s := a[low:high] 左闭右开：[1,3)
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s)) //s:[2 3] len(s):2 cap(s):4
	// 容量cap：从切片左边的index到数组最后一位，即数组长度减index：5-1=4

	// 为了方便起见，可以省略切片表达式中的任何索引。省略了low则默认为0；省略了high则默认为切片操作数的长度:
	// a[2:]  // 等同于 a[2:len(a)]
	// a[:3]  // 等同于 a[0:3]
	// a[:]   // 等同于 a[0:len(a)]

	//完整切片表达式
	// 对于数组，指向数组的指针，或切片a(注意不能是字符串)支持完整切片表达式：a[low : high : max]
}

func slice3() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3]                                                       // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))       //s:[2 3] len(s):2 cap(s):4
	s2 := s[3:4]                                                      // 索引的上限是cap(s)而不是len(s)
	fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2)) //s2:[5] len(s2):1 cap(s2):1
}
