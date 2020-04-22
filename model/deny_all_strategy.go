package model

import "container/heap"

type DenyAllStrategy struct {
	BaseStrategy
	leastReplicatedRank uint16
}

func (strategy *DenyAllStrategy) Apply(rankedStudent *RankedStudent) bool {
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

	count := strategy.countEdgeReplicas(pq)
	for i := 0; i < count; i++ {
		rs := heap.Pop(pq).(*RankedStudent)
		rs.Student().ClearCourse()
		pq.IncSpots()
	}

	if lrr := strategy.leastReplicatedRank; lrr < 1 || lastRank < lrr {
		strategy.leastReplicatedRank = lastRank
	}

	if rank < lastRank {
		heap.Push(pq, rankedStudent)
		pq.DecSpots()
	}

	return rank < lastRank
}
