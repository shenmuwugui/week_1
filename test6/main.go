package main

import "fmt"

func main() {
	//var i *int
	//fmt.Println(i, &i)
	//*i = 10
	//fmt.Println(i, &i)

	//var a [5]int
	//for i := 0; i < 6; i++ {
	//	a[i] = i
	//}

	//ch := make(chan int, 5)
	//close(ch)
	//ch <- 0   //往已经关闭的通道里发数据
	//close(ch) //关闭已关闭的通道
	diyfun()
}

func diyfun() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("圣诞节的快感不符合接口的")
		}
	}()
	panic("错误")
}
