package leecode

import (
	"fmt"
	"testing"
)

/*
	There are two sorted arrays nums1 and nums2 of size m and n respectively.

	Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

	You may assume nums1 and nums2 cannot be both empty.

	给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

	请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

	你可以假设 nums1 和 nums2 不会同时为空。
*/

/*
	nums1 = [1, 3]
	nums2 = [2]

	The median is 2.0

	nums1 = [1, 2]
	nums2 = [3, 4]

	The median is (2 + 3)/2 = 2.5
*/

func TestMedian(t *testing.T) {
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("the median of the two arrays is ", findMedianSortedArrays1(array[0:5], array[0:6]))
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	arrayMap := map[int][]int{
		1: nums1,
		2: nums2,
	}
	var tempArray []float64
	var temp float64
	for _, nums := range arrayMap {
		numsLen, midIdx := len(nums), 0
		midIdx = numsLen / 2
		temp = float64(nums1[midIdx])
		if numsLen%2 == 0 {
			temp = float64(nums1[midIdx-1]+nums1[midIdx]) / 2
		}
		tempArray = append(tempArray, temp)
	}
	return (tempArray[0] + tempArray[1]) / 2
}
