package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, v := range s {
		initial = fn(initial, v)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := len(s) - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	out := []int{}
	for _, v := range s {
		if fn(v) {
			out = append(out, v)
		}
	}
	return out
}

func (s IntList) Length() int {
	l := 0
	for range s {
		l++
	}
	return l
}

func (s IntList) Map(fn func(int) int) IntList {
	out := []int{}
	for _, v := range s {
		out = append(out, fn(v))
	}
	return out
}

func (s IntList) Reverse() IntList {
	out := []int{}
	for i := len(s) - 1; i >= 0; i-- {
		out = append(out, s[i])
	}
	return out
}

func (s IntList) Append(lst IntList) IntList {
	out := make([]int, len(s)+len(lst))
	copy(out, s)
	for i := 0; i < len(lst); i++ {
		out[i+len(s)] = lst[i]
	}
	return out
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, v := range lists {
		s = s.Append(v)
	}
	return s
}
