package main

import (
	"fmt"
	"strconv"
)

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
	// var res []int
	// res = append(res, 1)

	// a := [][]int{}
	// b := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	// a[1][2] = 3

	// str := "abcde"
	// fmt.Println(str[:5])
	// for k, _ := range str {
	// 	fmt.Println(k)
	// }

	// var a int= 20   /* 声明实际变量 */
	// var ip *int        /* 声明指针变量 */

	// ip = &a  /* 指针变量的存储地址 */

	// fmt.Printf("a 变量的地址是: %x\n", &a  )

	// /* 指针变量的存储地址 */
	// fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

	// /* 使用指针访问值 */
	// fmt.Printf("*ip 变量的值: %d\n", *ip )

	str := "3322251"
	start, end := 0, 0
	var res string
	for end < len(str) {
		for end < len(str) && str[start] == str[end] {
			end++
		}
		res = res + strconv.Itoa(end-start) + string(str[start])
		start = end
	}
	fmt.Printf(res)

}
