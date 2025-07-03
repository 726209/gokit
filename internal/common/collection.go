package common

// Map 映射函数：将 []T 映射为 []R
func Map[T any, R any](input []T, fn func(T) R) []R {
	result := make([]R, len(input))
	for i, v := range input {
		result[i] = fn(v)
	}
	return result
}

// Filter 过滤函数：返回满足条件的 []T 子集
func Filter[T any](input []T, fn func(T) bool) []T {
	result := make([]T, 0)
	for _, v := range input {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// Reduce 累加函数：将 []T 聚合成一个 R 值
func Reduce[T any, R any](input []T, initial R, fn func(R, T) R) R {
	acc := initial
	for _, v := range input {
		acc = fn(acc, v)
	}
	return acc
}

// Find 查找函数：返回首个满足条件的元素及是否找到
func Find[T any](input []T, fn func(T) bool) (T, bool) {
	for _, v := range input {
		if fn(v) {
			return v, true
		}
	}
	var zero T
	return zero, false
}
