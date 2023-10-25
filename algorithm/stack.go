package algorithm

type Stack struct{
	Values []int
	Len int
}

func (s *Stack) Push(val int) {
	s.Values = append(s.Values, val)
	s.Len++
}

func (s *Stack) Pop() int {
	val := s.Values[s.Len-1]
	s.Values = s.Values[:s.Len-2]
	s.Len--
	return val
}