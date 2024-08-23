package main

import "fmt"

func main() {
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
            if i != j && nums[i]+nums[j] == target {
				fmt.Println("values sum = ", nums[i], nums[j])
				fmt.Println("index of in slice = ", i, j)
				return []int{i, j}
			}
		}
	}
    return []int{0, 0}
}
