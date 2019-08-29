package heap

func swap(data []int, i, j int) []int {
	temp := data[i]
	data[i] = data[j]
	data[j] = temp
	return data
}

// 将给定数组调整为大顶堆
func heapAdjust(data []int, i, length int) {
	tmp := data[i]

	// 以节点i为根的左子树在数组中为2i+1
	for k := 2*i + 1; k < length; k = k*2 + 1 {
		// 对比左子树和右子树 选出大的
		if k+1 < length && data[k] < data[k+1] {
			k++
		}

		// 最大的子树与根比较
		if tmp < data[k] {
			data[i] = data[k]
			i = k // 将根移动到大的子树
		} else {
			break
		}
	}

	// 插入原父节点数据
	data[i] = tmp
}

func Sort(data []int) {

	// 初始化大顶堆 把最大的元素作为二叉树的根
	for i := len(data)/2 - 1; i >= 0; i-- {
		heapAdjust(data, i, len(data))
	}

	// 调整交换结构 每一次都将最大的子节点替换为根并与最小的节点交换
	for j := len(data) - 1; j > 0; j-- {
		swap(data, 0, j)
		heapAdjust(data, 0, j)
	}

}
