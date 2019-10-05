package tree

import (
	"sort"
	"errors"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type byRecord []Record

func (r byRecord) Len() int           { return len(r) }
func (r byRecord) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r byRecord) Less(i, j int) bool { return r[i].ID < r[j].ID }

func Build(input []Record) (*Node, error) {
	// NOTE: Shortcuts for specific conditions
	if len(input) == 0 {
		return nil, nil
	} else if len(input) == 1 {
		return &Node{ID: input[0].ID}, nil
	}

	// NOTE: Now we order the records
	sort.Sort(byRecord(input))
	if input[0].ID != 0 {
		return nil, errors.New("no root node")
	}

	nodes := make(map[int]*Node)
	for _, r := range input {

	}

	// NOTE: We only need the root node
	return nodes[0], nil
}
