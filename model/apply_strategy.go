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

func NewApplyStrategy(condition Condition, exceedLimit int) ApplyStrategy {
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
			make(RankCount),
		}
	case conditions.AllowSome:
		return &AllowSomeStrategy{
			base,
			0,
			make(RankCount),
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

	if !jc.IsFull() {
		heap.Push(pq, rankedStudent)
		jc.DecSpots()
		return true
	}

	heap.Push(pq, rankedStudent)
	rs := heap.Pop(pq).(*RankedStudent)
	rs.Student().ClearCourse()
	return rankedStudent != rs
}

func (strategy *BaseStrategy) countEdgeReplicas() int {
	pq := strategy.jointCourse.Students()
	students := []*RankedStudent{
		heap.Pop(pq).(*RankedStudent),
	}
	count, rank := 0, students[0].Rank()
	for ; students[count].Rank() == rank; students = append(students, heap.Pop(pq).(*RankedStudent)) {
		count++
	}
	for i := 0; i <= count; i++ {
		heap.Push(pq, students[i])
	}
	return count
}
