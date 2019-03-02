package quicksort

func quicksort(slice []int) []int {
	if len(slice) <= 1 {
		return slice
	}
	pivotIndex := 0
	var sliceLess, sliceMore []int
	for index, val := range slice {
		if index != pivotIndex {
			if val < slice[pivotIndex] {
				sliceLess = append(sliceLess, val)
			} else {
				sliceMore = append(sliceMore, val)
			}
		}
	}
	sliceLess = append(quicksort(sliceLess), slice[pivotIndex])
	return append(sliceLess, quicksort(sliceMore)...)
}
