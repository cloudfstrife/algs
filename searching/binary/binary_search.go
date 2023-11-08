package binary

// Search binary search
// Page number - Chinese Edition: 5
// Page number : 9
func Search(i int, a []int) int {
	start := 0
	end := len(a) - 1
	for start <= end {
		mid := start + (end-start)/2
		if i < a[mid] {
			end = mid - 1
		} else if i > a[mid] {
			start = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// RecursionSearch recursion binary search
// Page number - Chinese Edition: 15
// Page number : 25
func RecursionSearch(key int, a []int, start, end int) int {
	mid := start + (end-start)/2
	if start > end {
		return -1
	}
	if key < a[mid] {
		return RecursionSearch(key, a, start, mid-1)
	} else if key > a[mid] {
		return RecursionSearch(key, a, mid+1, end)
	} else {
		return mid
	}
}
