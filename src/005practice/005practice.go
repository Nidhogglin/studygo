package main

import "fmt"

func main() {

	// 1.有一堆数字，如果除了一个数字以外，其他数字都出现了两次，那么如何找到出现一次的数字？
	answer1()
	//拓展
	exchange(5, 3)
}

//用位运算相异或求只出现单次的数
func answer1() {
	a := []int{9, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 7, 7, 8, 8, 9, 1}
	result := 0
	for _, r := range a {
		result ^= r
	}
	fmt.Println(result)
}

//用位运算相异或交换两个数
func exchange(a int, b int) {
	a ^= b
	b ^= a
	a ^= b
	fmt.Println(a, b)
}
