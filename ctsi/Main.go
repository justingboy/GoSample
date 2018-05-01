package main

import (
	"fmt"
	wq "GoSample/waiqin"
	"math/bits"
	"time"
)

const COUNT = 100

var name = "zhangsan"

type student struct {
}

type IView interface {
}

func len() int {
	return 100
}

func learn() {
	fmt.Println("开始学习GO")
}

const IP string = "192.168.100.14"
const address = "回龙观东大街"
const (
	cat  string = "猫"
	dong        = "狗"
)
const v1, v2 = 1, "wl"

const name1 = iota
const name2 = iota
const (
	n1 = iota * 2
	n2
	n3
)

func main() {

one:
	fmt.Println("大码1")
	time.Sleep(1 * time.Second)
	goto one

	for {
		fmt.Println("A")
		time.Sleep(1 * time.Second)
	}

	fmt.Println("n1 = ", n1)
	fmt.Println("n2 = ", n2)
	fmt.Println("n3 = ", n3)
	fmt.Println(name1)
	fmt.Println(name2)

	var (
		a int    = 23
		b string = "zhansan"
	)

	c, d := 1, 2
	c = bits.Len(2)
	if a > 1 {
		fmt.Println("a >1")
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("c = %", c)
	fmt.Println(d)
	fmt.Println(wq.Age)

	wq.Register() //别名用法
	fmt.Println("hello go !")
	fmt.Println(COUNT)
	fmt.Println(name)
	learn()
	fmt.Println("end")

}
