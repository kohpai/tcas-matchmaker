package model

import (
	"container/heap"
)

type AllowSomeStrategy struct {
	BaseStrategy
	leastReplicatedRank uint16
	exceedLimit         uint16
}

func (strategy *AllowSomeStrategy) countBeingRemovedReplicas(pq *PriorityQueue) int {
	length, limit := pq.Len(), int(pq.Limit())
	delta := length - limit
	count := strategy.countEdgeReplicas(pq)

	if delta > int(strategy.exceedLimit) || count <= delta {
		return count
	}

	return 0
}

func (strategy *AllowSomeStrategy) Apply(rankedStudent *RankedStudent) bool {
	jc := strategy.jointCourse
	pq := jc.Students()
	rank := rankedStudent.Rank()

	if !pq.IsFull() {
		lrr := strategy.leastReplicatedRank
		if lrr == 0 || rank < lrr {
			heap.Push(pq, rankedStudent)
			return true
		}

		return false
	}

	tmp := heap.Pop(pq).(*RankedStudent)
	heap.Push(pq, tmp)
	lastRank := tmp.Rank()
	if rank > lastRank {
		return false
	}

	heap.Push(pq, rankedStudent)
	count := strategy.countBeingRemovedReplicas(pq)

	if count > 0 {
		strategy.leastReplicatedRank = lastRank
	}

	studentsBeingRemoved := make([]*Student, 0)
	for i := 0; i < count; i++ {
		rs := heap.Pop(pq).(*RankedStudent)
		student := rs.Student()
		student.ClearCourse()
		studentsBeingRemoved = append(studentsBeingRemoved, student)
	}
	strategy.findAndRemoveFromOthers(pq, studentsBeingRemoved)

	return rank < lastRank || count < 1
}
