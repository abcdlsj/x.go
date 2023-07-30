package skiplist

import "math/rand"

const MAX_LEVEL = 16

type Skiplist struct {
	CurMaxLevel int
	DummyHead   *SkiplistNode

	max int
	min int
}

type SkiplistNode struct {
	Elem      int
	NextLevel []*SkiplistNode
}

func New() Skiplist {
	return Skiplist{
		CurMaxLevel: 0,
		DummyHead: &SkiplistNode{
			Elem:      -1, // dummy SkiplistNode
			NextLevel: make([]*SkiplistNode, MAX_LEVEL),
		},
	}
}

func (s *Skiplist) Search(target int) bool {
	p := s.DummyHead
	for i := s.CurMaxLevel - 1; i >= 0; i-- {
		for p.NextLevel[i] != nil {
			if p.NextLevel[i].Elem == target {
				return true
			}

			if p.NextLevel[i].Elem > target {
				break
			}

			p = p.NextLevel[i]
		}
	}
	return false
}

func (s *Skiplist) Add(num int) {
	update := make([]*SkiplistNode, MAX_LEVEL)
	p := s.DummyHead

	for i := s.CurMaxLevel - 1; i >= 0; i-- {
		for p.NextLevel[i] != nil && p.NextLevel[i].Elem < num {
			p = p.NextLevel[i]
		}
		update[i] = p // insert position, SkiplistNode will insert after this.
	}

	level := 1
	for rand.Intn(2) == 1 {
		level++
	}
	if level > MAX_LEVEL {
		level = MAX_LEVEL
	}

	if level > s.CurMaxLevel {
		for i := s.CurMaxLevel; i < level; i++ {
			update[i] = s.DummyHead
		}
		s.CurMaxLevel = level
	}

	n := &SkiplistNode{
		Elem:      num,
		NextLevel: make([]*SkiplistNode, level),
	}

	for i := 0; i < level; i++ {
		n.NextLevel[i] = update[i].NextLevel[i]
		update[i].NextLevel[i] = n
	}

	s.updateMinMax(num)
}

func (s *Skiplist) Erase(num int) bool {
	if s.DummyHead.NextLevel[0] == nil {
		return false
	}

	update := make([]*SkiplistNode, MAX_LEVEL)
	p := s.DummyHead
	for i := s.CurMaxLevel - 1; i >= 0; i-- {
		for p.NextLevel[i] != nil && p.NextLevel[i].Elem < num {
			p = p.NextLevel[i]
		}
		update[i] = p
	}

	if update[0].NextLevel[0] == nil || update[0].NextLevel[0].Elem != num {
		return false
	}

	level := len(update[0].NextLevel[0].NextLevel)
	for i := 0; i < level; i++ {
		update[i].NextLevel[i] = update[i].NextLevel[i].NextLevel[i]
	}

	for i := s.CurMaxLevel - 1; s.DummyHead.NextLevel[i] == nil; i-- {
		s.CurMaxLevel--
	}

	return true
}

func (s *Skiplist) Range(min, max int) []int {
	var res []int

	p := s.DummyHead
	for i := s.CurMaxLevel - 1; i >= 0; i-- {
		for p.NextLevel[i] != nil && p.NextLevel[i].Elem < min {
			p = p.NextLevel[i]
		}
	}

	p = p.NextLevel[0]
	for p != nil && p.Elem <= max {
		res = append(res, p.Elem)
		p = p.NextLevel[0]
	}

	return res
}

func (s *Skiplist) updateMinMax(num int) {
	if s.min > num {
		s.min = num
	}
	if s.max < num {
		s.max = num
	}
}
