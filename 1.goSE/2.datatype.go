package main

import "fmt"

func main() {

	//两种声明变量方式：
	//1. 显式声明，关键词var
	var a int = 1
	//2. 隐式声明，:= ，类型自动注入
	b := 1
	// = 只能复制，不能声明
	b = 3
	fmt.Println(a, b)

	//字符串：
	str1 := "hello"
	str2 := "world"
	//len()输出字符串长度，可以和Java一样进行随意拼接
	fmt.Println(len(str1), str1+" "+str2)

	//数组，容量固定
	arr := [5]int{1, 2, 3, 4, 5}
	arr[1] = 100
	fmt.Println(len(arr), arr, arr[1], arr[2])

	//切片，容量可扩充，相当于动态数组
	slice := []int{1, 2, 3}
	slice[1] = 100
	//添加元素
	slice = append(slice, 4, 5, 6)
	//len返回切片长度，cap返回切片容量
	fmt.Println(len(slice), cap(slice), slice)

	//map key-value结构
	score := map[string]int{
		"zhangsan": 100,
		"lisi":     29,
		"wangwu":   34,
	}
	//赋值，添加新key
	score["tom"] = 90
	//读取，s表示返回值，ok表示存在与否
	s, ok := score["tom"]
	//删除元素
	delete(score, "lisi")
	fmt.Println(s, ok, score)
}
