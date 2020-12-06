package main

/**
 * @Author: yirufeng
 * @Date: 2020/12/6 2:52 下午
 * @Desc: ![PMjqtL](https://cdn.jsdelivr.net/gh/sivanWu0222/ImageHosting@master/uPic/PMjqtL.png)
 **/
//获取目标值左边界的对应位置
func GetLeftBounder(data []int, k int) int {
	l, r := 0, len(data)-1
	for l <= r {
		mid := l + (r-l)>>1
		if data[mid] > k {
			r = mid - 1
		} else if data[mid] < k {
			l = mid + 1
		} else if data[mid] == k {
			r = mid - 1
		}
	}

	//最后循环结束一定有 l = r + 1，此时l指向第一个出现目标值的元素或者目标值没有出现时第一个大于目标值的元素，或者在len(data)位置
	if l < len(data) && data[l] == k {
		return l //说明有出现
	}

	//说明没有出现
	return -1
}
