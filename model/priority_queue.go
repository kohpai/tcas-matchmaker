package model

import (
	"fmt"
)

type RankedStudent struct {
	student *Student
	rank    float32
	index   int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue struct {
	limit    uint16
	students []*RankedStudent
}

func (pq *PriorityQueue) Len() int { return len(pq.students) }

func (pq *PriorityQueue) Less(i, j int) bool {
	students := pq.students
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return students[i].rank >= students[j].rank
}

func (pq *PriorityQueue) Swap(i, j int) {
	students := pq.students
	students[i], students[j] = students[j], students[i]
	students[i].index = i
	students[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	students := pq.students
	n := len(students)
	item := x.(*RankedStudent)
	item.index = n
	pq.students = append(students, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := pq.students
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	pq.students = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Students() []*RankedStudent {
	return pq.students
}

func (pq *PriorityQueue) IsFull() bool {
	return uint16(pq.Len()) >= pq.limit
}

func (pq *PriorityQueue) Limit() uint16 {
	return pq.limit
}

func (pq *PriorityQueue) AvailableSpots() uint16 {
	length, limit := uint16(pq.Len()), pq.limit
	if pq.IsFull() {
		return length - limit
	}
	return limit - length
}

func (rs *RankedStudent) Student() *Student {
	return rs.student
}

func (rs *RankedStudent) Rank() float32 {
	return rs.rank
}

func (rs *RankedStudent) String() string {
	return fmt.Sprintf("{student: %v, rank: %v}", rs.student, rs.rank)
}
