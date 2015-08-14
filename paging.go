// main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	//	var arrayInt []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	//	arr, err := paging(1, 55, arrayInt)
	//	fmt.Println("arr:", arr, "err:", err)
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
