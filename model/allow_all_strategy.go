package model

import "container/heap"

type AllowAllStrategy struct {
	BaseStrategy
}

func (strategy *AllowAllStrategy) countBeingRemovedReplicas() uint16 {
	jc := strategy.jointCourse
	students := jc.Students().Students()
	length := uint16(len(students))
	limit := jc.Limit()

	count, rank := uint16(1), students[0].Rank()
	for _, s := range students[1:] {
		if s.Rank() != rank {
			break
		}
		count++
	}

	if delta := length - limit; count <= delta {
		return count
	}
	return 0
}

func (strategy *AllowAllStrategy) Apply(rankedStudent *RankedStudent) bool {
	jc := strategy.jointCourse
	pq := jc.Students()

	if !jc.IsFull() {
		heap.Push(pq, rankedStudent)
		jc.DecSpots()
		return true
	}

	rank := rankedStudent.Rank()
	lastRank := pq.Students()[0].Rank()

	switch {
	case rank == lastRank:
		heap.Push(pq, rankedStudent)
		return true
	case rank > lastRank:
		return false
	}

	heap.Push(pq, rankedStudent)
	count := strategy.countBeingRemovedReplicas()

	for ; count > 0; count-- {
		rs := heap.Pop(pq).(*RankedStudent)
		rs.Student().ClearCourse()
	}

	return true
}
