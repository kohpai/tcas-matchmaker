package model

import "container/heap"

type AllowAllStrategy struct {
	BaseStrategy
}

func (strategy *BaseStrategy) countBeingRemovedReplicas(pq *PriorityQueue) int {
	students := pq.Students()
	length, limit := len(students), int(pq.Limit())
	count := strategy.countEdgeReplicas(pq)

	if delta := length - limit; count <= delta {
		return count
	}
	return 0
}

func (strategy *AllowAllStrategy) apply(
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

	if !pq.IsFull() {
		// rejected, admitted, or nothing
		if !isSublist && !strategy.applySublist(rankedStudent) {
			return false
		}
		heap.Push(pq, rankedStudent)
		return true
	}

	rank := rankedStudent.Rank()
	tmp := heap.Pop(pq).(*RankedStudent)
	heap.Push(pq, tmp)
	lastRank := tmp.Rank()

	switch {
	case rank == lastRank:
		if !isSublist && !strategy.applySublist(rankedStudent) {
			return false
		}
		heap.Push(pq, rankedStudent)
		return true
	case rank > lastRank:
		return false
	}

	if !isSublist && !strategy.applySublist(rankedStudent) {
		return false
	}

	heap.Push(pq, rankedStudent)
	count := strategy.countBeingRemovedReplicas(pq)

	studentsBeingRemoved := make([]*Student, 0)
	for ; count > 0; count-- {
		rs := heap.Pop(pq).(*RankedStudent)
		student := rs.Student()
		student.ClearCourse()
		studentsBeingRemoved = append(studentsBeingRemoved, student)
	}
	strategy.findAndRemoveFromOthers(pq, studentsBeingRemoved)

	return true
}

func (strategy *AllowAllStrategy) Apply(rankedStudent *RankedStudent) bool {
	return strategy.apply(strategy.jointCourse.Students(), nil, rankedStudent, false)
}
