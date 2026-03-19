package search

// Search returns the index of the first occurrence of find in arr, or -1 if not found.
// Loop version of binary search.
func Search(arr []int, find int) int {
	if len(arr) == 0 {
		return -1
	}

	low, high := 0, len(arr)-1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == find {
			return mid
		}
		if arr[mid] < find {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

// SearchRecursive returns the index of the first occurrence of find in arr, or -1 if not found.
// Recursive version of binary search.
func SearchRecursive(arr []int, find int) int {
	if len(arr) == 0 {
		return -1
	}

	return searchRecursive(arr, find, 0, len(arr)-1)
}

func searchRecursive(arr []int, find int, low int, high int) int {
	if low > high {
		return -1
	}

	mid := (low + high) / 2
	if arr[mid] == find {
		return mid
	}
	if arr[mid] < find {
		return searchRecursive(arr, find, mid+1, high)
	} else {
		return searchRecursive(arr, find, low, mid-1)
	}
}
