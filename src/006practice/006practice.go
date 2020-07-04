package main

import "fmt"

func main() {
	// 1.编写代码打印9*9乘法表。
	answer1()

}

func answer1() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d X %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
