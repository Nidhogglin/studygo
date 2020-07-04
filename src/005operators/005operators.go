package main

import "fmt"

func main() {
	//1. 算术运算符
	fmt.Println(10 + 3) //加
	fmt.Println(10 - 3) //减
	fmt.Println(10 * 3) //乘
	fmt.Println(10 / 3) //除,整数除法舍弃余数 3
	fmt.Println(10 % 3) //取余 1
	//++（自增）和--（自减）在Go语言中是单独的语句，并不是运算符。
	a := 10
	a++ //a=a+1 = 11
	fmt.Println(a)
	a-- //a=a-1 = 10
	fmt.Println(a)

	//2. 关系运算符
	fmt.Println(10 > 3)
	fmt.Println(10 < 3)
	fmt.Println(10 == 3)
	fmt.Println(10 != 3)

	//3.逻辑运算符
	// 	&&	逻辑 AND 运算符。 如果两边的操作数都是 True，则为 True，否则为 False。
	//  ||	逻辑 OR 运算符。 如果两边的操作数有一个 True，则为 True，否则为 False。
	//   !	逻辑 NOT 运算符。 如果条件为 True，则为 False，否则为 True。
	fmt.Println(10 > 3 && 1 < 2) //true
	fmt.Println(10 < 3 || 1 < 2) //true
	fmt.Println(!(10 > 3))       //false

	// 4.位运算符
	x := 2              //010
	y := 5              //101
	fmt.Println(x & y)  //000 => 0 		两数相与：对应的二进制位都位1，对应位为1，否则为0
	fmt.Println(x | y)  //111 => 7 		两数相或：对应的二进制位有一个为1或两个都为1，对应为位1，否则为0
	fmt.Println(x ^ y)  //111 => 7 		两数相异或：对应二进制位不相等则为1，想等则为0
	fmt.Println(x << y) //1000000 => 64	左移n位就是乘以2的n次方。本例在10后面加5个0 => 2*2的5次方
	fmt.Println(y >> x) //001 => 1		右移n位就是除以2的n次方。本例将101前两位变成0 =>5/2的2次方

	//5. 赋值运算符
	// 	=	简单的赋值运算符，将一个表达式的值赋给一个左值
	// +=	相加后再赋值
	// -=	相减后再赋值
	// *=	相乘后再赋值
	// /=	相除后再赋值
	// %=	求余后再赋值
	// <<=	左移后赋值
	// >>=	右移后赋值
	// &=	按位与后赋值
	// |=	按位或后赋值
	// ^=	按位异或后赋值
	var n int
	n = 10  //10
	n += 10 //20
	n &= 15 //10100 & 1111 = 100 =>4
	fmt.Println(n)

}
