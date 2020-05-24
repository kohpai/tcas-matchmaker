package model

import "container/heap"

type DenyAllStrategy struct {
	BaseStrategy
	leastReplicatedRank float32
}

func (strategy *DenyAllStrategy) apply(
	pq *PriorityQueue,
	metadata *Metadata,
	rankedStudent *RankedStudent,
) bool {
	// @TODO the caller is responsible for creating the copy
	// copiedRs := &RankedStudent{
	// 	rankedStudent.Student(),
	// 	rank,
	// 	0,
	// }

	rank := rankedStudent.Rank()

	if !pq.IsFull() {
		lrr := strategy.leastReplicatedRank
		if lrr < 1 || rank < lrr && strategy.applySublist(rankedStudent) {
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

	count := strategy.countEdgeReplicas(pq)
	studentsBeingRemoved := make([]*Student, 0)
	for i := 0; i < count; i++ {
		rs := heap.Pop(pq).(*RankedStudent)
		student := rs.Student()
		student.ClearCourse()
		studentsBeingRemoved = append(studentsBeingRemoved, student)
	}
	strategy.findAndRemoveFromOthers(pq, studentsBeingRemoved)

	if lrr := strategy.leastReplicatedRank; lrr < 1 || lastRank < lrr {
		strategy.leastReplicatedRank = lastRank
	}

	if rank < lastRank && strategy.applySublist(rankedStudent) {
		heap.Push(pq, rankedStudent)
		return true
	}

	return false
}

func (strategy *DenyAllStrategy) Apply(rankedStudent *RankedStudent) bool {
	return strategy.apply(strategy.jointCourse.Students(), nil, rankedStudent)
}
