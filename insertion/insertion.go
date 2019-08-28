package insertion

func Sort(data []int) []int {
	size := len(data)

	// 插入排序 将待排序元素与前面所有元素比较找到插入位置并将后面的元素后移
	for i := 1; i < size; i++ {
		// 如果乱序
		if data[i] < data[i-1] {
			tmp := data[i]
			j := i - 1
			// 与前面所有元素进行比较
			for ; j >= 0 && data[j] > tmp; j-- {
				data[j+1] = data[j]
			}
			// 插入正确位置
			data[j+1] = tmp
		}
	}

	return data
}
