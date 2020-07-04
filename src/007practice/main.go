package main

import "fmt"

func main() {
	// 1.求数组[1, 3, 5, 7, 8]所有元素的和
	answer1()
	// 2.找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	answer2()
}

func answer1() {
	a := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range a {
		sum += v
	}
	fmt.Println(sum)
}

func answer2() {
	a := [...]int{1, 3, 5, 7, 8, 9, 11}
	// for index1, v1 := range a {
	// 	for index2, v2 := range a {
	// 		if v1+v2 == 10 && index1 < index2 {
	// 			fmt.Println(index1, index2)
	// 		}
	// 	}
	// }
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == 10 {
				fmt.Println(i, j)
			}
		}
	}
}
