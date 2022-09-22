package main

import (
	"fmt"
	"sort"
)

type fraction struct {
	i int
	j int
}

func (f *fraction) div() float64 {
	return float64(f.i) / float64(f.j)
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	i, j := 0, 1
	var fracts []fraction
	for i < len(arr)-1 {
		fracts = append(fracts, fraction{arr[i], arr[j]})
		j++
		if j == len(arr) {
			i++
			j = i + 1
		}
	}
	sort.Slice(fracts, func(i, j int) bool {
		return fracts[i].div() < fracts[j].div()
	})
	fmt.Println(fracts)
	return []int{fracts[k-1].i, fracts[k-1].j}
}

func main() {
	arr := []int{1, 2, 3, 5}
	k := 3
	fmt.Println(kthSmallestPrimeFraction(arr, k))
}
