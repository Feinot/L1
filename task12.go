package main

import (
	"fmt"
)

/*Имеется последовательность строк - (cat, cat, dog, cat, tree)
создать для нее собственное множество
*/
func main() {
	set := []string{"cat", "cat", "dog", "cat", "tree"}
	temp := make(map[string]struct{})
	res := make([]string, 0)

	for _, val := range set {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
			res = append(res, val)

		}
	}
	fmt.Println(res)
}
