// main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	//	var arrayInt []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	array := []map[string]interface{}{}
	arrayString := map[string]interface{}{"2": 3, "344g": 4}
	arrayString2 := map[string]interface{}{"323": 455, "fgrgr": 444324}
	for i := 0; i < 15; i++ {
		if i%2 == 0 {
			array = append(array, arrayString2)
		} else {
			array = append(array, arrayString)
		}

	}
	fmt.Println("array:", array)
	arr, err := paging_map_interface2(2, 3, array)
	//	arr, err :=
	//	web_paging(1, 5, arrayInt)
	//	fmt.Println("arr:", arr, "err:", err)
	fmt.Println("arr:", arr, "err:", err)
}

func paging(page_element_start int, page_size int, array []int) ([]int, string) {
	//	var arrayErr []int
	var arraySlice []int
	pageEnd := page_size + page_element_start - 1
	switch {
	case page_element_start > len(array):
		return arraySlice, "初始元素位置大于查询数组长度"
	case page_element_start+page_size-1 > len(array):
		arraySlice = array[page_element_start-1 : len(array)]
		return arraySlice, "已经到了最后一页"
		//	case page_size > len(array):
		//		arraySlice = array
	case page_element_start < len(array) && page_element_start+page_size-1 <= len(array) && page_size <= len(array):
		arraySlice = array[page_element_start-1 : pageEnd]
		return arraySlice, ""
	}
	return arraySlice, "未知情况"
}

func paging_map_interface(page_element_start int, page_size int, array []map[string]interface{}) ([]map[string]interface{}, string) {
	//	var arrayErr []int
	arraySlice := []map[string]interface{}{}
	pageEnd := page_size + page_element_start - 1
	switch {
	case page_element_start > len(array):
		return arraySlice, "初始元素位置大于查询数组长度"
	case page_element_start+page_size-1 > len(array):
		arraySlice = array[page_element_start-1 : len(array)]
		return arraySlice, "已经到了最后一页"
		//	case page_size > len(array):
		//		arraySlice = array
	case page_element_start < len(array) && page_element_start+page_size-1 <= len(array) && page_size <= len(array):
		arraySlice = array[page_element_start-1 : pageEnd]
		return arraySlice, ""
	}
	return arraySlice, "未知情况"
}

func paging_map_interface2(page_start int, page_size int, array []map[string]interface{}) ([]map[string]interface{}, string) {
	//	var arrayErr []int
	arraySlice := []map[string]interface{}{}

	page_start_element := (page_start - 1) * page_size
	//	pageEnd := page_size*page_element_start - 1

	switch {
	case page_start_element > len(array): //开始元素大于总数据个数
		return arraySlice, "初始元素位置大于查询数组长度"
	case page_start*page_size >= len(array): //
		arraySlice = array[page_start*page_size : len(array)]
		return arraySlice, "已经到了最后一页"

	case page_start_element < len(array) && page_start_element+page_size-1 <= len(array) && page_size <= len(array):
		//		arraySlice = array[page_element_start-1 : pageEnd]
		arraySlice = array[page_start_element : page_start_element+page_size]
		return arraySlice, ""
	}
	return arraySlice, "未知情况"
}

//func web_paging(page_start int, page_size int, array ...interface{}) {
//	length := len(array)
//	fmt.Println("length:", length)
//	page_end_element := page_start*page_size - 1 //每页最后一个元素
//	if page_end_element < length {
//		arraySlice := array[page_size*(page_start-1) : page_end_element]
//		fmt.Println("page_end_element : ", page_end_element)
//		fmt.Println("arr:", arraySlice)
//	} else {
//		arraySlice := array[page_size*(page_start-1) : length]
//		fmt.Println("arr:", arraySlice)
//	}
//	//	arraySlice := array[page_size*(page_start-1) : page_end_element]

//}
