package stringset

// Implement Set as a collection of unique string values.
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements. For example, a set with 2 elements, "a" and "b", should be formatted as {"a", "b"}.
// Format the empty set as {}.

const (
	initialCapacity = 20
	loadFactor      = 0.75
)

type Set struct {
	elements map[string]bool
}

func New() Set {
	return Set{elements: map[string]bool{}}
}

func NewFromSlice(l []string) Set {
	s := New()
	for _, v := range l {
		s.Add(v)
	}
	return s
}

func (s Set) String() string {
	out := "{"
	for e := range s.elements {
		out += "\"" + e + "\"" + ", "
	}
	if len(out) > 1 {
		out = out[:len(out)-2]
	}
	return out + "}"
}

func (s Set) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s Set) Has(elem string) bool {
	return s.elements[elem]
}

func (s Set) Add(elem string) {
	s.elements[elem] = true
}

func Subset(s1, s2 Set) bool {
	for e := range s1.elements {
		if !s2.Has(e) {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for e := range s1.elements {
		if s2.Has(e) {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	return Subset(s1, s2) && Subset(s2, s1)
}

func Intersection(s1, s2 Set) Set {
	intersection := New()
	for e := range s1.elements {
		if s2.Has(e) {
			intersection.Add(e)
		}
	}
	return intersection
}

func Difference(s1, s2 Set) Set {
	difference := New()
	for e := range s1.elements {
		if !s2.Has(e) {
			difference.Add(e)
		}
	}
	return difference
}

func Union(s1, s2 Set) Set {
	union := New()
	for e := range s1.elements {
		union.Add(e)
	}
	for e := range s2.elements {
		union.Add(e)
	}
	return union
}
