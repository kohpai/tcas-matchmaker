package model

import (
	"container/heap"
)

type AllowSomeStrategy struct {
	BaseStrategy
	leastReplicatedRank uint16
	rankCount           RankCount
	exceedLimit         uint16
}

func (strategy *AllowSomeStrategy) countBeingRemovedReplicas() (uint16, uint16) {
	jc := strategy.jointCourse
	pq := jc.Students()
	length, limit := uint16(pq.Len()), pq.Limit()
	delta := length - limit
	count := strategy.countEdgeReplicas()

	switch {
	case delta > strategy.exceedLimit:
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

		if rank == lrr {
			strategy.rankCount[rank] += 1
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
	count, inc := strategy.countBeingRemovedReplicas()

	if count > 0 {
		strategy.rankCount[lastRank] = count
		strategy.leastReplicatedRank = lastRank
	}

	for i := uint16(0); i < count; i++ {
		rs := heap.Pop(pq).(*RankedStudent)
		rs.Student().ClearCourse()
	}
	for i := uint16(0); i < inc; i++ {
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
