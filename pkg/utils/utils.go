package utils

func CompareAndSwapNum(max, min *int) {
	if *max < *min {
		max, min = min, max
	}
}
