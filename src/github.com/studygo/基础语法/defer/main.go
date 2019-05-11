package main

import (
	"fmt"
)

func testDefer() {
	//defer定义的代码会在函数执行结束前执行，最后最后执行，多个defer执行的顺序遵循堆栈规则，先定义的后执行
	defer fmt.Println("Defer1")
	defer fmt.Println("Defer2")
	defer fmt.Println("Defer3")
	fmt.Println("defer")
}

func testDefer2()  {
	//defer定义时的变量不会被后面的代码影响,只是执行的顺序变了
	i:=1
	defer fmt.Println(i)
	i=100
}
func main() {
	testDefer()
	testDefer2()
}
