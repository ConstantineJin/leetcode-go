package main

// 遍历，使用哈希表与空结构体实现set
// type EST struct {}

// func singleNumber(nums []int) int {
// 	set := make(map[int]EST)
// 	for _, v := range nums {
// 		if _, ok := set[v]; ok {
// 			delete(set, v)
// 		} else {
// 			set[v] = EST{}
// 		}
// 	}
// 	for k, _ := range set {
// 		return k
// 	}
// 	return 0
// }

// 位运算，使用异或
func singleNumber(nums []int) (single int) {
	for _, num := range nums {
		single ^= num
	}
	return
}
