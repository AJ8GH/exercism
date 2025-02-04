package tree

import (
	"errors"
	"slices"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Parent   int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	slices.SortFunc(records, func(a, b Record) int {
		switch {
		case a.ID > b.ID:
			return 1
		case a.ID < b.ID:
			return -1
		default:
			return 0
		}
	})

	nodesById := map[int]*Node{}
	rootId := -1

	for _, record := range records {
		_, exists := nodesById[record.ID]
		if exists {
			return nil, errors.New("dupe node")
		}
		if record.Parent > record.ID {
			return nil, errors.New("parent higher than child")
		}
		nodesById[record.ID] = &Node{ID: record.ID, Parent: record.Parent}
		if record.ID == record.Parent {
			rootId = record.ID
		}
	}

	if rootId == -1 {
		return nil, errors.New("no root")
	}

	for _, record := range records {
		if record.ID == rootId {
			continue
		}
		_, prevExists := nodesById[record.ID-1]
		if !prevExists {
			return nil, errors.New("non-continuous")
		}
		parent := nodesById[record.Parent]
		parent.Children = append(parent.Children, nodesById[record.ID])
	}

	return nodesById[rootId], nil
}

//
// slices.SortFunc(records, func(a, b Record) int {
// 	switch {
// 	case a.ID > b.ID:
// 		return 1
// 	case b.ID < b.ID:
// 		return -1
// 	default:
// 		return 0
// 	}
// })
