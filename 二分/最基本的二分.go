package main

/**
 * @Author: yirufeng
 * @Date: 2020/12/6 2:52 下午
 * @Desc:
 **/

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] == target {
			return mid
		}
	}

	return -1
}
