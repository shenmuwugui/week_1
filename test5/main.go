package main

import "fmt"

type Human struct {
	Name string
}

type GFS struct {
	Human
	Height float64
	Wealth string
	Face   int
}

func (s *Human) pay() {
	fmt.Println("支付")
}

func (s *GFS) Say(num float64) {
	fmt.Printf("今年%v实现了一个小目标，金额是%f亿\n", s.Name, num)
}

func (s *GFS) Eat(str string) {
	fmt.Printf("过年%v吃了%v\n", s.Name, str)
}

type Parrot struct {
	Name string
}

type Bird interface {
	Fly()
}

func (a Parrot) Fly() {
	fmt.Println(a.Name, "飞行")
}

func main() {
	a := Parrot{
		Name: "鹦鹉",
	}
	var a1 Bird
	a1 = a
	a1.Fly()

	as := GFS{Human{"张三"}, 1.87, "几个亿", 94}
	as.Say(1.2)
	as.Eat("满汉全席")
}
