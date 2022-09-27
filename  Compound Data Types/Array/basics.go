package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3}
	changeArr(arr)
	fmt.Printf("外层的数组%v\n", arr)
	changeArrOK(&arr)
	fmt.Printf("外层的数组-指针%v\n", arr)
}

// 初始化数组
func initArr() {
	var arr [3]int
	a := [3]int{1, 2, 3}
	b := [...]int{5, 4, 7}
	c := [...]int{99: 1}

	fmt.Printf("%v\n", arr)
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("arr = %d,a =%d,arr的容量%d,a的容量%d", len(arr), len(a), cap(arr), len(a))
}

// 类型的比较
func typeArr() {
	a := [3]int{1, 2, 3}
	b := [4]int{}
	c := [3]int{1, 2, 3}
	//无法将 'b' (类型 [4]int) 用作类型 [3]int
	//a = b
	fmt.Printf("%v,%v\n", a, b)
	fmt.Printf("%v\n", a == c)
}

// 传参
func changeArr(arr [3]int) {
	arr[0] = 99
	fmt.Printf("函数里的数组%v\n", arr)
}

// 传指针
func changeArrOK(arr *[3]int) {
	arr[0] = 99
	fmt.Printf("函数的数组-指针%v\n", arr)
}
