package clicker

// this was added in go 1.21, which is not available on Debian 12
func max[T int | int64 | float32 | float64](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func min[T int | int64 | float32 | float64](a T, b T) T {
	if a < b {
		return a
	}
	return b
}
