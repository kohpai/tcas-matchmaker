package applystrategy

import (
	"container/heap"

	"github.com/kohpai/tcas-3rd-round-resolver/model/common"
	"github.com/kohpai/tcas-3rd-round-resolver/model/course"
	jc "github.com/kohpai/tcas-3rd-round-resolver/model/jointcourse"
)

type ApplyStrategy interface {
	SetJointCourse(*jc.JointCourse)
	Apply(common.RankedStudent) bool
}

type BaseStrategy struct {
	jointCourse *jc.JointCourse
}

func NewApplyStrategy(condition course.Condition, exceedLimit int) ApplyStrategy {
	base := BaseStrategy{
		nil,
	}

	conditions := course.Conditions()

	switch condition {
	case conditions.AllowAll():
		return &AllowAllStrategy{
			base,
		}
	case conditions.DenyAll():
		return &DenyAllStrategy{
			base,
			0,
			make(course.RankCount),
		}
	case conditions.AllowSome():
		return &AllowSomeStrategy{
			base,
			0,
			make(course.RankCount),
			exceedLimit,
		}
	}

	return &base
}

func (strategy *BaseStrategy) SetJointCourse(jc *jc.JointCourse) {
	strategy.jointCourse = jc
}

func (strategy *BaseStrategy) Apply(rankedStudent common.RankedStudent) bool {
	jc := strategy.jointCourse
	pq := jc.Students()

	if !jc.IsFull() {
		heap.Push(pq, rankedStudent)
		jc.DecSpots()
		return true
	}

	heap.Push(pq, rankedStudent)
	rs := heap.Pop(pq).(common.RankedStudent)
	rs.Student().ClearCourse()
	return rankedStudent != rs
}

func (strategy *BaseStrategy) countEdgeReplicas() int {
	pq := strategy.jointCourse.Students()
	students := []common.RankedStudent{
		heap.Pop(pq).(common.RankedStudent),
	}
	count, rank := 0, students[0].Rank()
	for ; students[count].Rank() == rank; students = append(students, heap.Pop(pq).(common.RankedStudent)) {
		count++
	}
	for i := 0; i <= count; i++ {
		heap.Push(pq, students[i])
	}
	return count
}
