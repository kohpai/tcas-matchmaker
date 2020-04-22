package model

import (
	"container/heap"
)

type AllowSomeStrategy struct {
	BaseStrategy
	leastReplicatedRank uint16
	exceedLimit         uint16
}

func (strategy *AllowSomeStrategy) countBeingRemovedReplicas(pq *PriorityQueue) (int, int) {
	length, limit := pq.Len(), int(pq.Limit())
	delta := length - limit
	count := strategy.countEdgeReplicas(pq)

	switch {
	case delta > int(strategy.exceedLimit):
		return count, count - delta
	case count <= delta:
		return count, 0
	}

	return 0, 0
}

func (strategy *AllowSomeStrategy) Apply(rankedStudent *RankedStudent) bool {
	jc := strategy.jointCourse
	pq := jc.Students()
	rank := rankedStudent.Rank()

	if !pq.IsFull() {
		lrr := strategy.leastReplicatedRank
		if lrr == 0 || rank < lrr {
			heap.Push(pq, rankedStudent)
			pq.DecSpots()
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
	count, inc := strategy.countBeingRemovedReplicas(pq)

	if count > 0 {
		strategy.leastReplicatedRank = lastRank
	}

	for i := 0; i < count; i++ {
		rs := heap.Pop(pq).(*RankedStudent)
		rs.Student().ClearCourse()
	}
	for i := 0; i < inc; i++ {
		pq.IncSpots()
	}

	if rank < lastRank {
		return true
	}
	if count > 0 {
		return false
	}
	return true
}
