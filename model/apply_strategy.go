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
		heap.Push(pq, *rankedStudent)
		jc.DecSpots()
		return true
	}

	heap.Push(pq, *rankedStudent)
	rs := heap.Pop(pq).(RankedStudent)
	rs.Student().ClearCourse()
	return rankedStudent.Student().CitizenId() != rs.Student().CitizenId()
}

func (strategy *BaseStrategy) countEdgeReplicas() uint16 {
	pq := strategy.jointCourse.Students()
	students := []RankedStudent{
		heap.Pop(pq).(RankedStudent),
	}
	rank := students[0].Rank()
	for ; students[len(students)-1].Rank() == rank; students = append(students, heap.Pop(pq).(RankedStudent)) {
	}
	for _, student := range students {
		heap.Push(pq, student)
	}
	return uint16(len(students) - 1)
}
