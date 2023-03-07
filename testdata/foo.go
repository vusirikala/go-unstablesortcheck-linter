package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 2, 1}
	sort.Sort(sort.IntSlice(nums))
	fmt.Println(nums)

	strs := []string{"c", "b", "a"}
	sort.Slice(strs, func(i, j int) bool { return strs[i] < strs[j] })
	fmt.Println(strs)
}
