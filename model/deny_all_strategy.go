package model

import "container/heap"

type DenyAllStrategy struct {
	BaseStrategy
	Metadata
}

func (strategy *DenyAllStrategy) apply(
	pq *PriorityQueue,
	metadata *Metadata,
	rankedStudent *RankedStudent,
	isSublist bool,
) bool {
	// @TODO the caller is responsible for creating the copy
	// copiedRs := &RankedStudent{
	// 	rankedStudent.Student(),
	// 	rank,
	// 	0,
	// }

	rank := rankedStudent.Rank()

	if !pq.IsFull() {
		lrr := metadata.leastReplicatedRank
		if lrr < 1 || rank < lrr {
			if !isSublist && !strategy.applySublist(rankedStudent) {
				return false
			}

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

	if lrr := metadata.leastReplicatedRank; lrr < 1 || lastRank < lrr {
		metadata.leastReplicatedRank = lastRank
	}

	if rank < lastRank {
		if !isSublist && !strategy.applySublist(rankedStudent) {
			return false
		}

		heap.Push(pq, rankedStudent)
		return true
	}

	return false
}

func (strategy *DenyAllStrategy) Apply(rankedStudent *RankedStudent) bool {
	return strategy.apply(strategy.jointCourse.Students(), &strategy.Metadata, rankedStudent, false)
}
