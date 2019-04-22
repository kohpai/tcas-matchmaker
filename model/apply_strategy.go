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

func NewApplyStrategy(condition Condition) ApplyStrategy {
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
		}
	}

	return &base
}

func (strategy *BaseStrategy) SetJointCourse(jc *JointCourse) {
	strategy.jointCourse = jc
}

func (strategy *BaseStrategy) Apply(rankedStudent *RankedStudent) bool {
	pq := strategy.jointCourse.Students()
	heap.Push(pq, rankedStudent)
	rs := heap.Pop(pq).(*RankedStudent)
	rs.Student().ClearCourse()
	return rankedStudent != rs

	// @ASSERTION, this shouldn't happen
	// if !course.JointCourse().Apply() {
	// 	log.Println("Apply returns false")
	// }
}
