package algorithm

type Queue struct{
	Values []int
	Len int
}

func NewQueue() *Queue{
	return &Queue{
		Values: make([]int, 0),
	}
}

func (q *Queue) InQueue(val int) {
	q.Values = append(q.Values, val)
	q.Len++
}

func (q *Queue) DeQueue() int{
	val := q.Values[0]
	q.Values = q.Values[1:]
	q.Len--
	return val
}