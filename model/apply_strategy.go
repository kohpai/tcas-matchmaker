package model

import (
	"container/heap"
)

type ApplyStrategy interface {
	SetJointCourse(*JointCourse)
	Apply(*RankedStudent) bool
}

type BaseStrategy struct {
	jointCourse *JointCourse
}

func NewApplyStrategy(condition Condition, exceedLimit uint16) ApplyStrategy {
	base := BaseStrategy{
		nil,
	}

	conditions := Conditions()

	switch condition {
	case conditions.AllowAll:
		return &AllowAllStrategy{
			base,
		}
	case conditions.DenyAll:
		return &DenyAllStrategy{
			base,
			0,
		}
	case conditions.AllowSome:
		return &AllowSomeStrategy{
			base,
			0,
			exceedLimit,
		}
	}

	return &base
}

func (strategy *BaseStrategy) SetJointCourse(jc *JointCourse) {
	strategy.jointCourse = jc
}

func (strategy *BaseStrategy) Apply(rankedStudent *RankedStudent) bool {
	jc := strategy.jointCourse
	pq := jc.Students()

	if !pq.IsFull() {
		heap.Push(pq, rankedStudent)
		pq.DecSpots()
		return true
	}

	heap.Push(pq, rankedStudent)
	rs := heap.Pop(pq).(*RankedStudent)
	rs.Student().ClearCourse()
	return rankedStudent != rs
}

func (strategy *BaseStrategy) countEdgeReplicas(pq *PriorityQueue) int {
	lastStudent := heap.Pop(pq).(*RankedStudent)
	if pq.Len() < 1 {
		heap.Push(pq, lastStudent)
		return 1
	}

	nextStudent := heap.Pop(pq).(*RankedStudent)
	students := []*RankedStudent{lastStudent}
	for ; lastStudent.Rank() == nextStudent.Rank(); nextStudent = heap.Pop(pq).(*RankedStudent) {
		students = append(students, nextStudent)
		if pq.Len() < 1 {
			break
		}
	}

	for _, rs := range students {
		heap.Push(pq, rs)
	}
	if students[len(students)-1] != nextStudent {
		heap.Push(pq, nextStudent)
	}

	return len(students)
}
