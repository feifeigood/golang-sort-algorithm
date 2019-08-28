package shell

import "fmt"

func Sort(data []int) []int {
	fmt.Println(fmt.Sprintf("unsort: %v", data))

	// 求 delta 循环 每排序依次 delta 缩小一次 用delta来切割待排序列
	for delta := len(data) / 2; delta > 0; delta /= 2 {
		// 从delta开始排序
		for i := delta; i < len(data); i++ {
			j := i
			// 按delta间隔分组排序
			for j-delta >= 0 && data[j] < data[j-delta] {
				// 交换位置
				tmp := data[j]
				data[j] = data[j-delta]
				data[j-delta] = tmp
				j -= delta
			}
		}
		fmt.Println(fmt.Sprintf("delta: %d, data: %v", delta, data))
	}

	return data
}
