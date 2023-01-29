package main

import "fmt"

/*
3.定义一个长度是10的数值型数组，并对其进行按逆序进行排序
*/
func main() {
	num1 := [10]int{4, 2, 1, 5, 9, 3, 7, 6, 10, 15}
	var x int
	for i := 0; i < len(num1); i++ {
		f := false
		for j := 1; j < len(num1)-i; j++ {
			if num1[j] > num1[j-1] {
				x = num1[j]
				num1[j] = num1[j-1]
				num1[j-1] = x
				f = true
			}
		}
		if !f {
			break
		}
	}
	fmt.Println(num1)
}
