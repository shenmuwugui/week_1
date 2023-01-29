package main

import "fmt"

/*
2. 定义一个map入参函数, 对入参进行遍历，其遍历结果是否有序，为什么会	有这样的结果，请从源码层次进行解释
*/

func main() {
	map1 := make(map[int]string, 5)
	map1[0] = "qwertzxcvb"
	map1[1] = "asdf"
	map1[2] = "dsfg"
	map1[3] = "asjghfdf"
	map1[4] = "assdrtdf"
	fb(map1)
}

func fb(a map[int]string) {
	for k, v := range a {
		fmt.Println(k, v)
	}
}
