package model

import "container/heap"

type AllowAllStrategy struct {
	BaseStrategy
}

func (strategy *AllowAllStrategy) countRemovingReplicas() uint16 {
	jc := strategy.jointCourse
	students := jc.Students().Students()
	length := uint16(len(students))
	limit := jc.Limit()

	if length < limit {
		return 0
	}

	count, rank := uint16(1), students[0].Rank()
	for _, s := range students[1:] {
		if s.Rank() != rank {
			break
		}
		count++
	}

	if length == limit {
		if count > 1 {
			return 0
		}
		return 1
	}

	if delta := length - limit; count > delta+1 {
		return 0
	}
	return count
}

func (strategy *AllowAllStrategy) Apply(rankedStudent *RankedStudent) bool {
	jc := strategy.jointCourse
	pq := jc.Students()

	rank := rankedStudent.Rank()
	lastRank := pq.Students()[0].Rank()

	switch {
	case rank == lastRank:
		heap.Push(pq, rankedStudent)
		return true
	case rank > lastRank:
		return false
	default:
		count := strategy.countRemovingReplicas()
		if count > 0 {
			jc.IncSpots()
		}

		heap.Push(pq, rankedStudent)
		for i := uint16(0); i < count; i++ {
			rs := heap.Pop(pq).(*RankedStudent)
			rs.Student().ClearCourse()
		}

		return true
	}
}
