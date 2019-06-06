package model

import "container/heap"

type DenyAllStrategy struct {
	BaseStrategy
	leastReplicatedRank float64
	rankCount           RankCount
}

func (strategy *DenyAllStrategy) countBeingRemovedReplicas() uint16 {
	return strategy.countEdgeReplicas()
}

func (strategy *DenyAllStrategy) Apply(rankedStudent *RankedStudent) bool {
	jc := strategy.jointCourse
	pq := jc.Students()
	rank := rankedStudent.Rank()

	if !jc.IsFull() {
		lrr := strategy.leastReplicatedRank
		if lrr == -1 || rank > lrr {
			heap.Push(pq, rankedStudent)
			jc.DecSpots()
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
	if rank < lastRank {
		return false
	}

	heap.Push(pq, rankedStudent)
	count := strategy.countBeingRemovedReplicas()
	strategy.rankCount[lastRank], strategy.leastReplicatedRank = count, lastRank
	if count > 0 {
		rs := heap.Pop(pq).(*RankedStudent)
		rs.Student().ClearCourse()
	}
	for i := uint16(1); i < count; i++ {
		rs := heap.Pop(pq).(*RankedStudent)
		rs.Student().ClearCourse()
		jc.IncSpots()
	}

	return rank > lastRank
}
