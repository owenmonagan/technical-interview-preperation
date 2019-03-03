package mergesort

func Mergesort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}
	pivotIndex := len(slice) / 2
	leftSlice := slice[:pivotIndex]
	rightSlice := slice[pivotIndex:]
	leftSlice = Mergesort(leftSlice)
	rightSlice = Mergesort(rightSlice)
	return merge(leftSlice, rightSlice)
}

func merge(left, right []int) []int {
	size, leftIndex, rightIndex := len(left)+len(right), 0, 0
	slice := make([]int, size, size)
	for k := 0; k < size; k++ {
		if leftIndex >= len(left) {
			slice[k] = right[rightIndex]
			rightIndex++
		} else if rightIndex >= len(right) {
			slice[k] = left[leftIndex]
			leftIndex++
		} else if left[leftIndex] < right[rightIndex] {
			slice[k] = left[leftIndex]
			leftIndex++
		} else {
			slice[k] = right[rightIndex]
			rightIndex++
		}
	}
	return slice
}
