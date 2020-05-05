package main

import "fmt"

// 定义一个 key 为 string 类型，value 为 int 类型的 map 并赋值
var map0 = map[string]int{
	"k0": 0,
	"k1": 1,
}

func defMap() {
	// 创建一个为空 map
	map1 := make(map[int]string)

	//定义一个 map，零值为 nil
	var map2 map[string]int
	fmt.Printf("map0=%v, map1=%v, map2=%v\n", map0, map1, map2)
}

func mapOp() {
	// 取得 key 为 k0 的 value
	v0 := map0["k0"]
	// 不存在的 key，返回值为 value 类型 int 零值 0
	v3 := map0["noKey"]

	// 可以返回两个值
	if v, isExist := map0["k1"]; isExist {
		fmt.Printf("key k1 exist, value is %v\n", v)
	} else {
		fmt.Printf("key k1 dose not exist")
	}

	delete(map0, "k1")
	// 删除不存在的 key 不会报错
	delete(map0, "noKey")

	fmt.Printf("v0=%v, v3=%v\n", v0, v3)
	fmt.Printf("map0=%v\n", map0)

	// 增加 key value
	map0["k3"] = 3
	fmt.Printf("map0 len=%d", len(map0))
	// map 的遍历，不保证顺序
	for k, v := range map0 {
		fmt.Printf("map0 has key=%v, value=%v\n", k, v)
	}
}

func main() {
	defMap()
	mapOp()
}
