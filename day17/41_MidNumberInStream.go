package day17

import "math/rand"

type MedianFinder struct {
	arr []int
}

// Constructor
// initialize your data structure here
func Constructor() MedianFinder {
	return MedianFinder{[]int{}}
}

func (f *MedianFinder) AddNum(num int) {
	f.arr = append(f.arr, num)
}

func (f *MedianFinder) FindMedian() float64 {
	if len(f.arr) == 0 {
		return 0
	}
	mid := (len(f.arr) + 1) / 2
	qsort1(f.arr, 0, len(f.arr)-1)
	if len(f.arr)%2 == 0 {
		return float64(f.arr[mid-1]+f.arr[mid]) / 2
	} else {
		return float64(f.arr[mid-1])
	}
}

func qsort1(nums []int, l, r int) {
	partition := func(nums []int, l, r, p int) int {
		swap := func(i, j int) { nums[i], nums[j] = nums[j], nums[i] }
		pivot := nums[p]

		swap(r, p)
		p = l // pointer
		for i := l; i < r; i++ {
			if nums[i] < pivot {
				swap(i, p)
				p++
			}
		}
		swap(r, p)

		return p
	}
	if l < r {
		p := partition(nums, l, r, rand.Int()%(r-l+1)+l)
		qsort1(nums, l, p-1)
		qsort1(nums, p+1, r)
	}
}
