package insertion

func Sort(data []int) ([]int, int) {
	size := len(data)

	for i := 1; i < size; i++ {
		if data[i] < data[i-1] {
			tmp := data[i]
			j := i - 1
			for ; j >= 0 && data[j] > tmp; j-- {
				data[j+1] = data[j]
			}
			data[j+1] = tmp
		}
	}

	return data, size
}
