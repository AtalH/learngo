package main

import "fmt"

// 数组是值传递，方法内对数组进行修改，不影响调用方的原值
func forArray(arr [4]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("arr index %d value %d\n", i, arr[i])
	}
}

func rangeArray(arr [5]int) {
	for i, v := range arr {
		fmt.Printf("arr index %d value %d\n", i, v)
	}
}

// 传递数组地址，可以实现引用传递
func arrayAddr(addr *[4]int) {
	addr[0] = 1
}

func main() {
	var arr0 [5]int
	arr1 := [4]int{0, 1, 2, 3}
	arr2 := [...]int{0, 1, 2, 3}
	var arr3 [7]int = [...]int{0, 1, 2, 3, 4, 5, 6}
	var arr4 [2][3]string

	fmt.Println(arr0)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	forArray(arr2)

	rangeArray(arr0)

	arrayAddr(&arr2)
	fmt.Println(arr2)
}
