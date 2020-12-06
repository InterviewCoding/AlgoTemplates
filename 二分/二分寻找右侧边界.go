package main

/**
 * @Author: yirufeng
 * @Date: 2020/12/6 2:52 下午
 * @Desc: ![o2kYdO](https://cdn.jsdelivr.net/gh/sivanWu0222/ImageHosting@master/uPic/o2kYdO.png)
 **/

//获取目标值右边界的对应位置
func GetRightBounder(data []int, k int) int {
	l, r := 0, len(data)-1
	for l <= r {
		mid := l + (r-l)>>1
		if data[mid] > k {
			r = mid - 1
		} else if data[mid] < k {
			l = mid + 1
		} else if data[mid] == k {
			l = mid + 1
		}
	}

	if r >= 0 && data[r] == k {
		return r
	}

	//说明没有出现
	return -1
}
