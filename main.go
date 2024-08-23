package main

import "fmt"

func main() {

	arr := []int{8, 7, 5, 6, 9, 3, 2, 0, 3, 2, 5, 6, 9, 8}

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i]+arr[j] == 10 {
				fmt.Println(arr[i], arr[j])
				break
			}
		}
		break
	}

}
