package main

import "fmt"

func main() {

	nums := make([]int, 3 /*初始长度*/, 5 /*初始容量&每次增加的容量*/)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(nums) /*长度*/, cap(nums) /*容量*/, nums /*元素*/)
	nums = append(nums, 1)
	nums = append(nums, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(nums), cap(nums), nums)
	nums = append(nums, 3) // 满元素时, 会自动增加容量
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(nums), cap(nums), nums)

	fmt.Println("-----------------------------------")

	nums2 := make([]int, 3) // 默认cap=len,当len=0时cap=1
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(nums2), cap(nums2), nums2)

	fmt.Println("-----------------------------------")

	// 切片
	s0 := []int{1, 2, 3} // len=cap=3
	s1 := s0[0:2]        // 截取左闭右开区间[0,2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(s1), cap(s1), s1)
	/*
		[startIndex:endIndex]
		- startIndex 留空, 默认值为0
		- endIndex 留空, 默认值为len()
	*/

	// 切片和原数组使用的时同块内存, 即浅拷贝, 本质上时数组指向的头尾指针改变
	s1[0] = 666
	fmt.Println("s0 =", s0)
	fmt.Println("s1 =", s1)

	s2 := make([]int, 3)
	copy(s2, s0) // 深拷贝
	fmt.Println("s2 =", s2)
	s2[0] = 0
	fmt.Println("s0 =", s0)
	fmt.Println("s2 =", s2)
}
