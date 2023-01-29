package main

import "fmt"

func main() {
	slice1 := make([]int, 0)
	slice1 = append(slice1, 1, 2, 3, 4)
	arr := fb(slice1)
	fmt.Println(arr)
}

func fb(arr []int) []int {
	var ret []int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			for k := 0; k < len(arr); k++ {
				if arr[i] != arr[j] && arr[i] != arr[k] && arr[j] != arr[k] {
					ret = append(ret, arr[i]*100+arr[j]*10+arr[k])
				}
			}
		}
	}
	return ret
}
