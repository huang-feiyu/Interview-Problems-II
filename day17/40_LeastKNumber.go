package day17

import "math/rand"

func getLeastNumbers(arr []int, k int) []int {
	qsort(arr, 0, len(arr)-1, k)
	return arr[0:k]
}

func qsort(arr []int, l, r, k int) {
	if l < r {
		p := partition(arr, l, r, rand.Int()%(r-l+1)+l)
		if p == k {
			return
		} else if p > k {
			qsort(arr, l, p-1, k)
		} else {
			qsort(arr, p+1, r, k)
		}
	}
}

func partition(arr []int, l, r, p int) int {
	swap := func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }

	pivot := arr[p]
	swap(p, r) // move to the end

	p = l // [l, p) less than pivot; (p, r] more than or equal to pivot
	for i := l; i < r; i++ {
		if arr[i] < pivot {
			swap(p, i)
			p++
		}
	}
	swap(p, r)

	return p
}
