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
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(a[3:])
}
