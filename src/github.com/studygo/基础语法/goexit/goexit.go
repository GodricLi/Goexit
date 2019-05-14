package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("ccc")
	// return
	// 终止所在的协程,函数在结束前会执行defer，所在的协程之后的代码不会执行
	runtime.Goexit()

	fmt.Println("ddd")
}

func main() {
	// 创建新的协程
	go func ()  {
		fmt.Println("aaa")
		
		// 调用别的函数
		test()
		
		fmt.Println("bbb")
	
		}()
	
		//写一个死循环不让主程序结束
	for {

	}
}