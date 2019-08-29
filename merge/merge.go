package merge

func merge(data, temp []int, left, right, mid int) {
	i := left    // 左下标
	j := mid + 1 // 右下标

	t := 0 // temp 数组位置

	// 比较两边数组
	for i <= mid && j <= right {
		if data[i] < data[j] {
			temp[t] = data[i]
			t++
			i++
		} else {
			temp[t] = data[j]
			t++
			j++
		}
	}

	// 左边剩余元素推入临时
	for i <= mid {
		temp[t] = data[i]
		t++
		i++
	}

	// 右边剩余元素推入临时
	for j <= right {
		temp[t] = data[j]
		t++
		j++
	}

	// temp 数组数据放入原数组
	t = 0
	for left <= right {
		data[left] = temp[t]
		left++
		t++
	}
}

func sort(data, temp []int, left, right int) {
	if left < right {
		mid := (left + right) / 2           // 选取数组中间值对半切
		sort(data, temp, left, mid)         // 左边数组继续分治
		sort(data, temp, mid+1, right)      // 右边数组继续分治
		merge(data, temp, left, right, mid) // 归并数组
	}
}

func Sort(data []int) {
	temp := make([]int, len(data))   // 建立一个临时数组
	sort(data, temp, 0, len(data)-1) // 开始分治
}
