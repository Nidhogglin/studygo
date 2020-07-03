package main

import "fmt"

func main() {
	// 1.编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用fmt.Printf()搭配%T分别打印出上述变量的值和类型。
	answer1()
	fmt.Println()
	// 2.编写代码统计出字符串"hello沙河小王子"中汉字的数量。
	answer2()
}

func answer1() {
	var i int = 666
	var f float64 = 666.666
	var b bool
	var str string = "Nidhogg"
	fmt.Printf("%d(%T) %f(%T) %v(%T) %s(%T)", i, i, f, f, b, b, str, str)
}

func answer2() {
	var str = "hello沙河小王子"
	count := 0
	s := []rune(str)
	for i := 0; i < len(s); i++ {
		if s[i] >= 0x4E00 && s[i] <= 0x9FA5 {
			count++
		}
	}
	fmt.Println(count)
}
