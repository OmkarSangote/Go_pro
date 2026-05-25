package main

type MinStack struct {
	items [][2]int
}

func (s *MinStack) Push(val int) {
	if len(s.items) == 0 {
		s.items = append(s.items, [2]int{val, val})
		return
	}

	currentMin := s.items[len(s.items)-1][1]

	if val < currentMin {
		currentMin = val
	}

	s.items = append(s.items, [2]int{val, currentMin})
}

func (s *MinStack) Pop() {
	if len(s.items) == 0 {
		return
	}
	s.items = s.items[:len(s.items)-1]
}

func (s *MinStack) Top() int {
	return s.items[len(s.items)-1][0]
}

func (s *MinStack) GetMin() int {
	return s.items[len(s.items)-1][1]
}
