package selection

func Sort(data []int) []int {

	for i := 0; i < len(data); i++ {
		k := i
		// 从i+1开始找最小的元素
		for j := i + 1; j < len(data); j++ {
			// 找到相对小的元素给k
			if data[j] < data[k] {
				k = j
			}
		}
		// k=i 说明没有找到比i更小的 不需要替换
		if k != i {
			// 将i+1起始的序列中最小的元素与当前i元素替换位置
			tmp := data[i]
			data[i] = data[k]
			data[k] = tmp
		}
	}

	return data
}
