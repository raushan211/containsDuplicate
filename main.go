package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4, 1}
	result := containsDuplicate(input)
	fmt.Println(result)

}
func containsDuplicate(nums []int) bool {
	result := false
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i] == nums[j] && i != j {

				return true

			}
		}
	}
	return result
}
