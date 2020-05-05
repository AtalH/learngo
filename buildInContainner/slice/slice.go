package main

import "fmt"

var arr = [...]int{0, 1, 2, 3, 4, 5, 6}

func sliceAppend() {
	slice := arr[3:5]
	// 对 slice 进行追加，会改变源数组的内容
	slice0 := append(slice, 50)
	slice1 := append(slice0, 60)
	// 当追加超出源数组的长度时，底层会给这个 slice 创建一个新的数组
	// 扩容 cap 按 2 的指数增加
	slice2 := append(slice1, 70)
	fmt.Printf("\nafter append, arr=%v\n", arr)
	fmt.Println(slice, "\n", slice0, "\n", slice1, "\n", slice2)
}

func createSlice() {
	fmt.Println("in createSlice()")
	// 定义一个 int 型 slice ，零值是 nil
	var slice []int
	// nil 也不会错误
	slice = append(slice, 1)

	// 定义一个 int 型 slice
	slice0 := []int{2, 4, 5, 6}

	// 创建一个空的 slice，并指定长度为5（容量也是5）
	slice1 := make([]int, 5)

	// 创建一个空的 slice，并指定长度为5，容量是10
	slice2 := make([]int, 5, 10)
	fmt.Println(slice, "\n", slice0, "\n", slice1, "\n", slice2)
	fmt.Printf("slice1 value=%v, len=%d, cap=%d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 value=%v, len=%d, cap=%d\n", slice2, len(slice2), cap(slice2))
}

func sliceDel() {
	fmt.Println("in sliceDel()")
	slice := arr[:]
	// 删除索引为 3 的元素
	slice = append(slice[:3], slice[4:]...)
	fmt.Printf("slice value=%v, len=%d, cap=%d\n", slice, len(slice), cap(slice))

	// 删除第一个元素
	slice = slice[1:]
	fmt.Printf("slice value=%v, len=%d, cap=%d\n", slice, len(slice), cap(slice))

	// 删除最后一个元素
	slice = slice[:len(slice)-1]
	fmt.Printf("slice value=%v, len=%d, cap=%d\n", slice, len(slice), cap(slice))
}

func sliceCopy() {
	// 此处只是定义了一个为 nil 的 slice，没有空间，无法拷贝
	//var targetSlice []int
	// 使用 make 得到 slice 容量为 7
	targetSlice := make([]int, 7)
	srcSlice := arr[:]
	copy(targetSlice, srcSlice)
	fmt.Printf("srcSlice=%v, targetSlice=%v\n", srcSlice, targetSlice)
}

// 切片是数组的视图，对切片的修改会影响源数组的值，是引用传递
func main() {
	// slice0 value=[2 3 4], len=3, cap=5
	var slice0 []int = arr[2:5]
	slice1 := arr[:4]
	slice2 := arr[:]
	// 对切片进行重切片，虽然索引超出了 slice0 的长度，但没超出容量值
	slice3 := slice0[1:4]
	fmt.Printf(" arr=%v\n slice0=%v\n slice1=%v\n slice2=%v\n slice3=%v\n", arr, slice0, slice1, slice2, slice3)
	fmt.Printf("slice0 value=%v, len=%d, cap=%d", slice0, len(slice0), cap(slice0))

	sliceAppend()
	createSlice()
	sliceDel()
	sliceCopy()
}
