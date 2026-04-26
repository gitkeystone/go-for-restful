// 闭包函数返回另一个函数
package main

import "fmt"

// 生成器模式：根据给定的条件，每次生成一个新项目
func generator() func() int { // 外部函数：返回值可以是一个接口，此时，内部函数要实现这个接口
	var i = 0
	return func() int { // 内部函数
		i++
		return i
	}
}
func main() {
	numGenerator := generator() // 生成器可以记住状态，而不用每次根据入参返回对应的值
	for i := 0; i < 5; i++ {
		fmt.Print(numGenerator(), "\t")
	}
}
