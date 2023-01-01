package collection

func Contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func IncrementElements(arr []int, val int) {
	for i := 0; i < len(arr); i++ {
		arr[i] += val
	}
}
