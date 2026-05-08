package util

// 计算最大页数
func ClacMaxPage(total, size int64) int64 {
	return (total + size - 1) / size
}
