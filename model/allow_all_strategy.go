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

func (strategy *AllowAllStrategy) Apply(rankedStudent *RankedStudent) bool {
	jc := strategy.jointCourse
	pq := jc.Students()

	if !pq.IsFull() {
		// rejected, admitted, or nothing
		// if !strategy.applySublist(rankedStudent) {
		// 	return false
		// }
		heap.Push(pq, rankedStudent)
		return true
	}

	rank := rankedStudent.Rank()
	tmp := heap.Pop(pq).(*RankedStudent)
	heap.Push(pq, tmp)
	lastRank := tmp.Rank()

	switch {
	case rank == lastRank:
		// if !strategy.applySublist(rankedStudent) {
		// 	return false
		// }
		heap.Push(pq, rankedStudent)
		return true
	case rank > lastRank:
		return false
	}

	// if !strategy.applySublist(rankedStudent) {
	// 	return false
	// }

	heap.Push(pq, rankedStudent)
	count := strategy.countBeingRemovedReplicas(pq)

	// studentsBeingRemoved := make([]*Student, 0)
	for ; count > 0; count-- {
		rs := heap.Pop(pq).(*RankedStudent)
		student := rs.Student()
		student.ClearCourse()
		// studentsBeingRemoved = append(studentsBeingRemoved, student)
	}
	// strategy.findAndRemoveFromOthers(pq, studentsBeingRemoved)

	return true
}
