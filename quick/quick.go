package quick

func partition(data []int, low, high int) int {
	i := low
	j := high
	pivot := data[low]

	for i < j {
		// 从后向前扫描
		for i < j && data[j] >= pivot {
			j--
		}
		data[i] = data[j] // 比pivot小的数向前移动

		// 从前向后扫描
		for i < j && data[i] <= pivot {
			i++
		}
		data[j] = data[i] // 比pivot大的数向后移动
	}

	data[i] = pivot

	return i
}

func Sort(data []int, low, high int) {
	if low < high {
		loc := partition(data, low, high) // 划分半区
		Sort(data, low, loc-1)            // 排序前半区
		Sort(data, loc+1, high)           // 排序后半区
	}
}
