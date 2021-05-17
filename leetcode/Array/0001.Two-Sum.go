package main

import "fmt"

func main() {
	arr := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(arr[0], arr[1])
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

