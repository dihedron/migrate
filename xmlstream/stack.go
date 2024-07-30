package xmlstream

type Stack []string

func (s *Stack) Len() int {
	if s == nil {
		return 0
	}
	return len(*s)
}

func (s *Stack) Empty() bool {
	return s.Len() == 0
}

func (s *Stack) Push(values ...string) {
	t := append(*s, values...)
	*s = t
}

func (s *Stack) Pop() string {
	value := s.Peek()
	*s = []string(*s)[:len(*s)-1]
	return value
}

func (s *Stack) Peek() string {
	if s.Len() == 0 {
		return ""
	}
	value := []string(*s)[len(*s)-1]
	return value
}

func (s *Stack) At(index int) string {
	if index >= 0 {
		panic("invalid non-negative index")
	}
	if s.Len() < (-1 * index) {
		panic("index larfìger than stack")
	}
	return []string(*s)[len(*s)+index]
}
