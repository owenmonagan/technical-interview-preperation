package quicksort

func Quicksort(slice []int) []int {
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
	sliceLess = append(Quicksort(sliceLess), slice[pivotIndex])
	return append(sliceLess, Quicksort(sliceMore)...)
}
