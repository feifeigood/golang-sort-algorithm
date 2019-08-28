package bubble

func Sort(data []int) []int {
	// 外层循环 每次 - 1 内层循环后会把 前i个元素中最大的移动到数组尾部 该位置不再需要排序
	for i := len(data) - 1; i > 0; i-- {
		flag := 0
		// 内层循环 每次交换都会出现一个最大值被排序好
		for j := 0; j < i; j++ {
			if data[j] > data[j+1] {
				tmp := data[j+1]
				data[j+1] = data[j]
				data[j] = tmp
				flag = 1
			}
		}

		// 没有交换 说明排序完成了
		if flag == 0 {
			break
		}
	}

	return data
}
