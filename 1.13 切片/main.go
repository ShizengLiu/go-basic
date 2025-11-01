package main

import (
	"fmt"
)

func main() {
	// 方式1，声明并初始化一个空的切片
	var s1 []int = []int{}

	// 方式2，类型推导，并初始化一个空的切片
	var s2 = []int{}

	// 方式3，与方式2等价
	s3 := []int{}

	// 方式4，与方式1、2、3 等价，可以在大括号中定义切片初始元素
	s4 := []int{1, 2, 3, 4}

	// 方式5，用make()函数创建切片，创建[]int类型的切片，指定切片初始长度为0
	s5 := make([]int, 0)

	// 方式6，用make()函数创建切片，创建[]int类型的切片，指定切片初始长度为2，指定容量参数4
	s6 := make([]int, 2, 4)

	// 方式7，引用一个数组，初始化切片
	arr := [5]int{6, 5, 4, 3, 2}
	// 从数组下标2开始，直到数组的最后一个元素
	s7 := arr[2:]
	// 从数组下标1开始，直到数组下标3的元素，创建一个新的切片
	s8 := arr[1:3]
	// 从0到下标2的元素，创建一个新的切片
	s9 := arr[:2]

	fmt.Println(s1, s2, s3, s4, s5, s6, s7, s8, s9, arr)
	changeArray()
	addSlice()
	removeSlice()
	copySlice()
}

func changeArray() {
	fmt.Println("change Array--------")
	arr := [5]int{1, 2, 3, 4, 5}
	s1 := arr[2:]
	s2 := arr[1:3]
	s3 := arr[:2]

	fmt.Println(s1, s2, s3)

	arr[0] = 9
	arr[1] = 8
	arr[2] = 7

	fmt.Println(s1, s2, s3)

	fmt.Println("len:", len(arr), "cap", cap(arr))
}

func addSlice() {

	fmt.Println("add Slice--------")
	s := []int{1, 2, 3}
	fmt.Println("original slice:", s, "len:", len(s), "cap:", cap(s))

	s = append(s, 4)
	fmt.Println("after append 4:", s, "len:", len(s), "cap:", cap(s))
	s = append(s, 5, 6, 7)
	fmt.Println("after append 5,6,7:", s, "len:", len(s), "cap:", cap(s))
	s = append(s, []int{8, 9, 10}...)
	fmt.Println("after append 8,9,10:", s, "len:", len(s), "cap:", cap(s))

	s4 := []int{11, 12, 13, 14}
	s4 = append(s4[:2], append([]int{3}, s4[2:]...)...)
	fmt.Println("s4 after insert 3 to index 2:", s4)
}

func removeSlice() {
	fmt.Println("remove Slice--------")
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("original slice:", s, "len:", len(s), "cap:", cap(s))

	// 删除下标为3的元素
	s = append(s[:3], s[4:]...)
	fmt.Println("after remove index 3:", s, "len:", len(s), "cap:", cap(s))
}

func copySlice() {

	src1 := []int{1, 2, 3}
	dst1 := make([]int, 4, 5)

	src2 := []int{1, 2, 3, 4, 5}
	dst2 := make([]int, 3, 3)

	fmt.Println("before copy, src1 = ", src1)
	fmt.Println("before copy, dst1 = ", dst1)

	fmt.Println("before copy, src2 = ", src2)
	fmt.Println("before copy, dst2 = ", dst2)

	copy(dst1, src1)
	copy(dst2, src2)

	fmt.Println("before copy, src1 = ", src1)
	fmt.Println("before copy, dst1 = ", dst1)

	fmt.Println("before copy, src2 = ", src2)
	fmt.Println("before copy, dst2 = ", dst2)

}
