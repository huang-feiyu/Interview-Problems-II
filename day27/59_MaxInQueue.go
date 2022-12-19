package day27

type MaxQueue struct {
	queue []int
	mono  []int
}

func Constructor() MaxQueue {
	return MaxQueue{[]int{}, []int{}}
}

func (q *MaxQueue) Max_value() int {
	if len(q.queue) > 0 {
		return q.mono[0]
	}
	return -1
}

func (q *MaxQueue) Push_back(x int) {
	for len(q.mono) > 0 && x > q.mono[len(q.mono)-1] {
		q.mono = q.mono[:len(q.mono)-1]
	}
	q.mono = append(q.mono, x)
	q.queue = append(q.queue, x)
}

func (q *MaxQueue) Pop_front() int {
	if len(q.queue) == 0 {
		return -1
	}
	x := q.queue[0]
	if x == q.mono[0] {
		q.mono = q.mono[1:]
	}
	q.queue = q.queue[1:]
	return x
}
