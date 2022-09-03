package main

func QuickSort(arr []int) []int {
	sorted_array := make([]int, len(arr))

	for i, v := range arr {
		sorted_array[i] = v
	}

	sort(sorted_array, 0, len(arr)-1)

	return sorted_array
}

func sort(arr []int, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			temp := arr[splitIndex]

			arr[splitIndex] = arr[i]
			arr[i] = temp

			splitIndex++
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	sort(arr, start, splitIndex-1)
	sort(arr, splitIndex+1, end)
}
