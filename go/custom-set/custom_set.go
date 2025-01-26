package stringset

import "hash/fnv"

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

const (
	initialSize = 20
)

type Set struct {
	elements [][]string
	size     int
}

func New() Set {
	return Set{
		elements: make([][]string, initialSize),
	}
}

func NewFromSlice(l []string) Set {
	s := Set{}
	s.resize(len(l) * 2)
	for _, v := range l {
		add(&s, v)
	}
	return s
}

func hash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}

func elementsForVal(s *Set, v string) []string {
	h := hash(v)
	i := h % len(s.elements)
	els := s.elements[i]
	if els == nil {
		return []string{}
	}

	return els
}

func add(s *Set, e string) {
	if s.size < len(s.elements) || s.size < initialSize {
		newSize := s.size * 2
		if s.size == 0 {
			newSize = initialSize
		}
		s.resize(newSize)
	}
	h := hash(e)
	i := h % len(s.elements)
	els := elementsForVal(s, e)
	for _, v := range els {
		if v == e {
			return
		}
	}
	s.size++
	s.elements[i] = append(els, e)
}

func (s *Set) resize(size int) Set {
	newElements := make([][]string, size)
	for i := 0; i < size; i++ {
		newElements[i] = []string{}
	}
	oldElements := s.elements
	s.elements = newElements
	if s.size > 0 {
		for _, v := range oldElements {
			for _, w := range v {
				add(s, w)
			}
		}
	}
	return *s
}

func (s Set) String() string {
	out := "{"
	for _, v := range s.elements {
		for _, w := range v {
			out += "\"" + w + "\"" + ", "
		}
	}
	if len(out) > 1 {
		out = out[:len(out)-2]
	}
	return out + "}"
}

func (s Set) IsEmpty() bool {
	return s.size == 0
}

func (s Set) Has(elem string) bool {
	if s.size == 0 {
		return false
	}
	els := elementsForVal(&s, elem)
	for _, v := range els {
		if v == elem {
			return true
		}
	}
	return false
}

func (s Set) Add(elem string) {
	add(&s, elem)
}

func Subset(s1, s2 Set) bool {
	for _, els := range s1.elements {
		for _, e := range els {
			if !s2.Has(e) {
				return false
			}
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for _, els := range s1.elements {
		for _, e := range els {
			if s2.Has(e) {
				return false
			}
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	return Subset(s1, s2) && Subset(s2, s1)
}

func Intersection(s1, s2 Set) Set {
	panic("Please implement the Intersection function")
}

func Difference(s1, s2 Set) Set {
	panic("Please implement the Difference function")
}

func Union(s1, s2 Set) Set {
	panic("Please implement the Union function")
}
