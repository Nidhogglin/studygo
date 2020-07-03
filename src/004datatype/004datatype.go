package main

import (
	"fmt"
	"math"
)

func main() {
	// 整型
	var x uint = 10 //无符号整型 32位操作系统默认是uint32 64位系统默认是uint64
	fmt.Println(x)
	//十进制
	var a int = 10         //32位操作系统默认是int32 64位系统默认是int64
	fmt.Printf("%d \n", a) //10
	fmt.Printf("%b \n", a) //1010 占位符%b表示二进制

	//八进制 以0开头
	var b int = 077
	fmt.Printf("%o \n", b) //77
	fmt.Printf("%d \n", b) //63 以10进制输入

	// 十六进制 以0x开头
	var c int = 0xff
	fmt.Printf("%x %X %d \n", c, c, c) // ff FF 255

	// 浮点型
	var y float32 = 30.25
	fmt.Printf("%f %f %.2f \n", y, math.Pi, math.Pi) //30.250000 3.141593 3.14

	//负数 复数有实部和虚部，complex64的实部和虚部为32位，complex128的实部和虚部为64位。
	var c1 complex64 = 1 + 2i
	var c2 complex128 = 2 + 3i
	fmt.Println(c1, c2) //(1+2i) (2+3i)

	//Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值。
	// 1. 布尔类型变量的默认值为false。
	// 2. Go 语言中不允许将整型强制转换为布尔型.
	// 3. 布尔型无法参与数值运算，也无法与其他类型进行转换。
	var isok bool = true
	fmt.Println(isok) //true

	//字符串
	var str1 string = "hello"
	str2 := "你好"
	fmt.Println(str1, str2)

	//字符串转义符
	//Go 语言的字符串常见转义符包含回车(\r)、换行(\n)、单(\')双引(\")号、制表符(\t)，反斜杠(\\)等
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")

	//多行字符串
	// Go语言中要定义一个多行字符串时，就必须使用反引号字符：
	// 反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。
	str3 := `	第一行\"
				第二行\n
				第三行
	`
	fmt.Println(str3)

	//字符 byte和rune类型
	var ch1 byte = 's' //uint8类型，或者叫 byte 型，代表了ASCII码的一个字符。
	var ch2 rune = '林' //rune类型，代表一个 UTF-8字符。
	fmt.Printf("%v(%c) %v(%c)", ch1, ch1, ch2, ch2)
	fmt.Println()

	//调用函数，无参数
	traversalString()
	changeString()
	//调用函数，有参数
	sqrtDemo(3, 4)
}

//遍历字符串
func traversalString() {
	s := "hello林倾"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	} //104(h) 101(e) 108(l) 108(l) 111(o) 230(æ) 158() 229(å) 128() 190(¾) 后面乱码
	fmt.Println()

	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	} //104(h) 101(e) 108(l) 108(l) 111(o) 26519(林) 20542(倾) 正常输出
	fmt.Println()
	// 因为UTF8编码下一个中文汉字由3~4个字节组成，所以我们不能简单的按照字节去遍历一个包含中文的字符串，否则就会出现上面输出中第一行的结果。

}

//修改字符串
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))

	s2 := "白萝卜"
	runeS2 := []rune(s2)
	runeS2[0] = '红'
	fmt.Println(string(runeS2))
}

func sqrtDemo(a int, b int) {
	var c int
	//math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b))) //把计算结果转换为int类型
	fmt.Println(c)
}
