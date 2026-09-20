package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if j, ok := seen[complement]; ok {
			return []int{j, i}
		}

		seen[num] = i
	}

	return nil
}

func main() {
	myNums := []int{2, 7, 11, 15}
	myTarget := 9

	result := twoSum(myNums, myTarget)
	fmt.Println("Индексы:", result)
}
