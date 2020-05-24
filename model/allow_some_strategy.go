package model

import (
	"container/heap"
)

type AllowSomeStrategy struct {
	BaseStrategy
	Metadata
	exceedLimit uint16
}

func (strategy *AllowSomeStrategy) countBeingRemovedReplicas(pq *PriorityQueue) int {
	length, limit := pq.Len(), int(pq.Limit())
	delta := length - limit
	count := strategy.countEdgeReplicas(pq)

	if delta > int(strategy.exceedLimit) || count <= delta {
		return count
	}

	return 0
}

func (strategy *AllowSomeStrategy) apply(
	pq *PriorityQueue,
	//@TODO make use of this metadata
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

	heap.Push(pq, rankedStudent)
	count := strategy.countBeingRemovedReplicas(pq)

	if count > 0 {
		metadata.leastReplicatedRank = lastRank
	}

	studentsBeingRemoved := make([]*Student, 0)
	for i := 0; i < count; i++ {
		rs := heap.Pop(pq).(*RankedStudent)
		student := rs.Student()
		student.ClearCourse()
		studentsBeingRemoved = append(studentsBeingRemoved, student)
	}
	strategy.findAndRemoveFromOthers(pq, studentsBeingRemoved)

	admitted := rank < lastRank || count < 1
	if admitted && !isSublist && !strategy.applySublist(rankedStudent) {
		heap.Remove(pq, rankedStudent.index)
		return false
	}
	return admitted
}

func (strategy *AllowSomeStrategy) Apply(rankedStudent *RankedStudent) bool {
	return strategy.apply(strategy.jointCourse.Students(), &strategy.Metadata, rankedStudent, false)
}
