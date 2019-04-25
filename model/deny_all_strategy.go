package model

import "container/heap"

type DenyAllStrategy struct {
	BaseStrategy
	leastReplicatedRank uint16
	rankCount           RankCount
}

func (strategy *DenyAllStrategy) countBeingRemovedReplicas() uint16 {
	students := strategy.jointCourse.Students().Students()
	count, rank := uint16(1), students[0].Rank()
	for _, s := range students[1:] {
		if s.Rank() != rank {
			break
		}
		count++
	}
	return count
}

func (strategy *DenyAllStrategy) Apply(rankedStudent *RankedStudent) bool {
	jc := strategy.jointCourse
	pq := jc.Students()
	rank := rankedStudent.Rank()

	if !jc.IsFull() {
		lrr := strategy.leastReplicatedRank
		if lrr == 0 || rank < lrr {
			heap.Push(pq, rankedStudent)
			jc.DecSpots()
			return true
		}

		if rank == lrr {
			strategy.rankCount[rank] += 1
		}
		return false
	}

	lastRank := pq.Students()[0].Rank()
	if rank > lastRank {
		return false
	}

	if rank < lastRank {
		count := strategy.countBeingRemovedReplicas()
		strategy.rankCount[lastRank] = count
		strategy.leastReplicatedRank = lastRank
		for ; count > 0; count-- {
			rs := heap.Pop(pq).(*RankedStudent)
			rs.Student().ClearCourse()
			jc.IncSpots()
		}

		heap.Push(pq, rankedStudent)
		jc.DecSpots()

		return true
	}

	rs := heap.Pop(pq).(*RankedStudent)
	rs.Student().ClearCourse()
	strategy.rankCount[rank] = 2
	strategy.leastReplicatedRank = rank
	jc.IncSpots()

	return false
}
