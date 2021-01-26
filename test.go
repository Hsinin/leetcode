package main

import "fmt"

func main() {
	// 变量声明
	// var a int = 99
	// var a = [5]int{1, 2, 3, 4, 5}

	// var b = "hello world"
	// b := [5]int{1, 2, 3, 4, 5}
	// b := [...]int{1, 2, 3, 4, 5}

	// func rotate(nums []int, k int) {
	//    newNums := make([]int, len(nums))
	//    for i, v := range nums {
	//        newNums[(i+k)%len(nums)] = v
	//    }
	//    copy(nums, newNums)
	// }

	// func reverse(a []int) {
	//   for i, n := 0, len(a); i < n/2; i++ {
	//       a[i], a[n-1-i] = a[n-1-i], a[i]
	//   }
	// }

	// func reverse(array []int) []int {
	// 	for i, n := 0, len(array); i < n/2; i++ {
	// 		tmp := array[i]
	// 		array[i] = array[n-1-i]
	// 		array[n-1-i] = tmp
	// 	}
	// 	return array
	// }
	// func main() {
	// 	a := []int{1, 2, 3, 4, 5}  //正确
	// 	a := [5]int{1, 2, 3, 4, 5} //报错
	// 	fmt.Println(reverse(a))
	// }

	// 遍历数组
	// for _,v := range(a) {
	// }

	// ********** map **********
	// tmp := make(map[string]int)
	// tmp["a"] = 1

	// 遍历map
	// for k,v := range(tmp) {

	// }

	// // golang判断key是否在map中
	// if _, ok := tmp[key]; ok {
	// //存在
	// }

	// 切片追加
	res := [11]int{}
	// var res []int
	// res = append(res, 1)
	res[10] = 1
	fmt.Println(res)
	// a := [5]int{1, 2, 3, 4, 5}
	// for _, v := range a {
	// 	fmt.Println(v)
	// }

}
