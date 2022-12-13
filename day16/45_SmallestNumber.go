package day16

import (
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

func minNumber1(nums []int) string {
	strs := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		strs[i] = strconv.Itoa(nums[i])
	}
	sort.SliceStable(strs, func(i, j int) bool {
		str1, str2 := strs[i], strs[j]
		return strings.Compare(str1+str2, str2+str1) < 0
	})
	//fmt.Printf("after sorted: %v\n", nums)
	return strings.Join(strs, "")
}

func minNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		strs[i] = strconv.Itoa(nums[i])
	}
	qsort(strs, 0, len(strs)-1)
	//fmt.Printf("after sorted: %v\n", nums)
	return strings.Join(strs, "")
}

func qsort(nums []string, l, r int) {
	if r > l {
		p := partition(nums, l, r, rand.Int()%(r-l+1)+l)
		qsort(nums, l, p-1)
		qsort(nums, p+1, r)
	}
}

func partition(nums []string, l, r, p int) int {
	swap := func(a, b int) { nums[a], nums[b] = nums[b], nums[a] }

	pivot := nums[p]
	swap(p, r) // move pivot to the end

	p = l // [l, p) less than pivot; (p, r] more than or equal to pivot
	for i := l; i < r; i++ {
		if strings.Compare(nums[i]+pivot, pivot+nums[i]) < 0 {
			swap(p, i)
			p++
		}
	}
	swap(r, p)
	return p
}
