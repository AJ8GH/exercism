package robotname

import (
	"errors"
	"strconv"
)

const (
	numCeil  = 9
	numFloor = 0
	letCeil  = 'Z'
	letFloor = 'A'
)

var (
	lets        = []string{}
	nums        = []string{}
	names       = []string{}
	lastNameIdx = -1
)

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if lastNameIdx == -1 {
		genNames()
	}

	if r.name == "" {
		if lastNameIdx >= len(names)-1 {
			return "", errors.New("too many names")
		}

		lastNameIdx++
		r.name = names[lastNameIdx]
	}

	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = ""
}

func genNames() {
	for i := numFloor; i <= numCeil; i++ {
		nums = append(nums, strconv.Itoa(i))
	}
	for i := letFloor; i <= letCeil; i++ {
		lets = append(lets, string(i))
	}

	for _, l1 := range lets {
		for _, l2 := range lets {
			for _, n1 := range nums {
				for _, n2 := range nums {
					for _, n3 := range nums {
						names = append(names, l1+l2+n1+n2+n3)
					}
				}
			}
		}
	}
}
