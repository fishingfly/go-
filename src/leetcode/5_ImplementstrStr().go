package main

import (
	"container/list"
	"fmt"
)

func strStr(haystack string, needle string) int {
	if len(needle) == 0 || len(haystack) == 0 {
		if needle == haystack || needle == "" {
			return 0
		}
		return -1
	}
	var map_index_string map[int]string
	l := list.New()
	map_index_string = make(map[int]string)
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] && i + len(needle) <= len(haystack) {
			map_index_string[i] = haystack[i : i + len(needle)]
			l.PushBack(i)
		}
	}
	for index := l.Front(); index != nil; index = index.Next(){
		if map_index_string[index.Value.(int)] == needle {
			return index.Value.(int)
		}
	}
	return -1
}

func main () {
	haystack := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	needle :=  "ab"
	fmt.Println(strStr(haystack,needle))

}