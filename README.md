# util
##Introduction
工具方法，用于收个人使用中一些小问题的解决方法类和函数。
## Project Status
Under developing and testing.

###[paging.go](https://github.com/JiangXuanYi/util/blob/master/paging.go)

2015-08-14

    1、通过数组Array、切片Slice、对所返回的数据进行切片分页 
    2、使用swich作为逻辑判定，减少了无线if&for的嵌套循环，增加代码的可读性。
    
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