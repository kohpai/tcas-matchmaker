package model

import "container/heap"

type AllowSomeStrategy struct {
	BaseStrategy
	leastReplicatedRank uint16
	rankCount           RankCount
	exceedLimit         uint16
}

func (strategy *AllowSomeStrategy) countBeingRemovedReplicas() (uint16, uint16) {
	jc := strategy.jointCourse
	students := jc.Students().Students()
	count, rank := uint16(1), students[0].Rank()
	for _, s := range students[1:] {
		if s.Rank() != rank {
			break
		}
		count++
	}

	length := uint16(len(students))
	limit := jc.Limit()
	delta := length - limit

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

	heap.Push(pq, rankedStudent)
	count, inc := strategy.countBeingRemovedReplicas()

	if count > 0 {
		strategy.rankCount[lastRank] = count
		strategy.leastReplicatedRank = lastRank
	}

	for ; count > 0; count-- {
		rs := heap.Pop(pq).(*RankedStudent)
		rs.Student().ClearCourse()
	}
	for ; inc > 0; inc-- {
		jc.IncSpots()
	}

	if rank < lastRank {
		return true
	}
	if count > 0 {
		return false
	}
	return true
}
