# util
##Introduction
工具方法，用于收个人使用中一些小问题的解决方法类和函数。
## Project Status
Under developing and testing.
###[ExportToXLS.go]
2015-08-24
1,解决前几日提交未被github计数问题
2,此demo提供一个方法，可将规则json导出至.xls。便于运营阶段的数据分析与统计。
###[paging.go](https://github.com/JiangXuanYi/util/blob/master/paging.go)

2015-08-14

1,通过数组Array、切片Slice、对所返回的数据进行切片分页 
2,使用swich作为逻辑判定，减少了无线if&for的嵌套循环，增加代码的可读性。
	
2015-08-21

1,
```go
	func paging_map_interface(page_element_start int, page_size int, array []map[string]interface{}) 
```
将上面的第一个参数page_element_start 换为 page_start,即
```go
	func paging_map_interface2(page_start int, page_size int, array []map[string]interface{})
```
即函数参数列表中文为
```
	(起始页码，分页大小（数据分页长度)，数据数组）
```
	
    
## Design (How does it works?)

If there are the array of json in my web apps.I want to paging it to display in the HTML on the screen,so I must know the startelement and the size of the page in the HTML.
```go
func paging(page_element_start int, page_size int, array []int) ([]int, string) {
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
	return arraySlice, "未知情况"}
```



For example,there is array of number, I call it [arrayInt]

```go
var arrayInt []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
arr, err := paging(1, 55, arrayInt)
```
Like this I will see the all number in the HTML of the first page

## Credits
[Jean.Jiang](https://github.com/JiangXuanYi)