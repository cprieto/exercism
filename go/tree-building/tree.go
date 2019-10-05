package tree

import (
	"errors"
	"sort"
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
	}

	// NOTE: Now we order the records
	sort.Sort(byRecord(input))
	nodes := make(map[int]*Node)

	if r := input[0]; r.ID != 0 {
		return nil, errors.New("you need to provide a root node")
	} else if r.ID == 0 && r.Parent != 0 {
		return nil, errors.New("Root node should not have a parent")
	}

	for i, r := range input {
		if r.ID != i {
			return nil, errors.New("non-continous nodes are a bug")
		} else if _, ok := nodes[r.ID]; ok {
			return nil, errors.New("no duplicate nodes are allowed")
		}

		p, ok := nodes[r.Parent]

		if r.ID != 0 && !ok {
			return nil, errors.New("parent node does not exist")
		} else if r.ID == 0 {
			nodes[0] = &Node{ID: 0}
			continue
		}

		n := &Node{ID: r.ID}
		p.Children = append(p.Children, n)
		nodes[r.ID] = n
	}

	// NOTE: We only need the root node
	return nodes[0], nil
}
