package main

import "fmt"

/*
Реализовать пересечение двух неупорядоченных множеств.
*/

//#1 способ
//func main(){
//
//	set1:= mapset.NewSet("cat","dog","elephant", "lion", "tiger")
//	set2:= mapset.NewSet("bird", "bear", "cat","dog", "fish")
//	fmt.Println(set1.Intersect(set2)) // {cat,dog}
//
//}

// #2 способ
func main() {
	set1 := []string{"cat", "dog", "elephant", "lion", "tiger"}
	set2 := []string{"bird", "bear", "cat", "dog", "fish"}

	m := make(map[string]int)
	intersection := make([]string, 0)
	for _, v := range set1 {
		m[v] += 1
	}

	for _, v := range set2 {
		m[v] += 1
	}

	for k := range m {
		if m[k] > 1 {
			intersection = append(intersection, k)
		}
	}
	fmt.Println(intersection) // {cat,dog}

}
