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

func (strategy *AllowSomeStrategy) applySublist(rankedStudent *RankedStudent) bool {
	jc := strategy.jointCourse
	student := rankedStudent.Student()

	genders := Genders()
	gender := student.Gender()

	var pq *PriorityQueue
	var metadata *Metadata

	switch gender {
	case genders.Male:
		pq, metadata = jc.MaleQ(), strategy.maleMetadata
	case genders.Female:
		pq, metadata = jc.FemaleQ(), strategy.femaleMetadata
	}

	copiedRs := &RankedStudent{
		rankedStudent.Student(),
		rankedStudent.Rank(),
		0,
	}
	if pq != nil && !strategy.apply(pq, metadata, copiedRs, true) {
		return false
	}

	programs := Programs()
	program := student.Program()

	switch program {
	case programs.Formal:
		pq, metadata = jc.FormalQ(), strategy.formalMetadata
	case programs.Inter:
		pq, metadata = jc.InterQ(), strategy.interMetadata
	case programs.Vocat:
		pq, metadata = jc.VocatQ(), strategy.vocatMetadata
	case programs.NonFormal:
		pq, metadata = jc.NonFormalQ(), strategy.nonFormalMetadata
	}

	copiedRs = &RankedStudent{
		rankedStudent.Student(),
		rankedStudent.Rank(),
		0,
	}
	if pq != nil && !strategy.apply(pq, metadata, copiedRs, true) {
		return false
	}

	return true
}

func (strategy *AllowSomeStrategy) apply(
	pq *PriorityQueue,
	metadata *Metadata,
	rankedStudent *RankedStudent,
	isSublist bool,
) bool {
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
