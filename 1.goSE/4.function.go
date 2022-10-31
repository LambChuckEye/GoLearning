package main

import "fmt"

// 外部显式定义函数
// 前参数，后返回值
func addSub(a, b int) (int, int) {
	return a + b, a - b
}

// ...表示参数数目可变
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func main() {
	fmt.Println()
	a, b := addSub(4, 3)
	fmt.Println(a, b)

	res1 := sum(1, 2, 3, 4)
	fmt.Println(res1)

	nums := []int{1, 2, 3, 4, 5, 6}
	//使用...把切片转换为参数（解包）
	res2 := sum(nums...)
	fmt.Println(res2)

	//声明f为函数类型
	f := func(in string) {
		fmt.Println(in)
	}
	f("hello world!")

	//声明匿名函数
	func(in string) {
		fmt.Println(in)
	}("hello world!")

}
