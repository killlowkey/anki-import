package util

// Filter 过滤函数，返回一个过滤后的结构体数组
func Filter[T any](items []T, filterFunc func(T) bool) []T {
	var filteredItems []T
	for _, item := range items {
		if filterFunc(item) {
			filteredItems = append(filteredItems, item)
		}
	}
	return filteredItems
}
