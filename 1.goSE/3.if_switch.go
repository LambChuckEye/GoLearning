package main

import "fmt"

func main() {

	//if判断
	condition := true
	if condition {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	//switch分支
	expr := "zhangsan"
	switch expr {
	case "zhangsan":
		fmt.Println("zhangsan")
	case "lisi":
		fmt.Println("lisi")
	default: //无匹配结果时执行
		fmt.Println("None")
	}

	//for 循环
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i, "double")
		}
	}
	
	//死循环
	i :=0
	for{
		i++
		fmt.Println(i)
		if i >100{
			break
		}
	}
	
}
