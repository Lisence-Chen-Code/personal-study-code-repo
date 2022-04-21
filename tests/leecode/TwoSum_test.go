package leecode

import (
	"fmt"
	"testing"
)

/**
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.
*/

/**
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1]
*/

func BenchmarkSth1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum([]int{1, 2, 3, 4, 5, 6, 7}, 8)
	}
}

func TestSth1(t *testing.T) {
	twoSum([]int{1, 2, 3, 4, 5, 6, 7}, 8)
}

func twoSum(nums []int, target int) (res []int) {
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

func TestMap(t *testing.T) {
	m := make(map[int]int, 9)
	for i := 0; i < 9; i++ {
		m[i] = i * i
	}
	if _, ok := m[64]; ok {
		fmt.Println(m[81])
	}
	fmt.Println("nothing")
}
