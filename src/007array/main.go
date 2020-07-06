package main

import "fmt"

func main() {
	// 数组是同一种数据类型元素的集合。 在Go语言中，数组从声明时就确定，使用时可以修改数组成员，但是数组大小不可变化。
	// 数组定义：var 数组变量名 [元素数量]T,eg：var a [5]int
	// 数组的长度必须是常量，并且长度是数组类型的一部分。一旦定义，长度不能变。 [5]int和[10]int是不同的类型。

	// 数组的初始化
	// 方法一：初始化数组时可以使用初始化列表来设置数组元素的值。
	fmt.Printf("方法一：\n")
	method1()
	// 方法二：按照上面的方法每次都要确保提供的初始值和数组长度一致，一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度，例如：
	fmt.Printf("方法二：\n")
	method2()
	// 方法三：可以使用指定索引值的方式来初始化数组
	fmt.Printf("方法三：\n")
	method3()

	// 数组遍历
	fmt.Printf("数组遍历：\n")
	arraybianli()

	//多维数组
	// 二维数组的定义和遍历
	fmt.Printf("二维数组：\n")
	erweiarray()

	// 数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。
	fmt.Println("数组是值类型：")
	a := [3]int{10, 20, 30}
	modifyArray(a) //在modify中修改的是a的副本x
	fmt.Println(a) //[10 20 30]
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	modifyArray2(b) //在modify中修改的是b的副本x
	fmt.Println(b)  //[[1 1] [1 1] [1 1]]
}

// 方法一：初始化数组时可以使用初始化列表来设置数组元素的值。
func method1() {
	var a [3]int                              //数组会初始化为int类型的零值
	var b = [4]int{1, 2}                      //使用指定的初始值完成初始化
	var c = [4]string{"北京", "上海", "广州", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(a)                            //[0 0 0]
	fmt.Println(b)                            //[1 2 0 0]
	fmt.Println(c)                            //[北京 上海 深圳]
}

// 方法二：一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度
func method2() {
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)                          //[0 0 0]
	fmt.Println(numArray)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
	fmt.Println(cityArray)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string
}

// 方法三：可以使用指定索引值的方式来初始化数组
func method3() {
	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a) //type of a:[4]int
}

// 遍历数组a有以下两种方法：
func arraybianli() {
	var a = [...]string{"北京", "上海", "深圳"}
	// 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 方法2：for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}
}

func erweiarray() {
	a := [...][2]string{ //多维数组只有第一层可以使用...来让编译器推导数组长度
		{"江西", "南昌"},
		{"浙江", "杭州"},
		{"广东", "广州"},
	}
	fmt.Println(a)
	fmt.Printf("二维数组遍历：\n")
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}

func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}
