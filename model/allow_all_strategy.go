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

func (strategy *AllowAllStrategy) applySublist(rankedStudent *RankedStudent) bool {
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

func (strategy *AllowAllStrategy) apply(
	pq *PriorityQueue,
	metadata *Metadata,
	rankedStudent *RankedStudent,
	isSublist bool,
) bool {
	if !isSublist && !strategy.applySublist(rankedStudent) {
		return false
	}

	if !pq.IsFull() {
		// rejected, admitted, or nothing
		heap.Push(pq, rankedStudent)
		return true
	}

	rank := rankedStudent.Rank()
	tmp := heap.Pop(pq).(*RankedStudent)
	heap.Push(pq, tmp)
	lastRank := tmp.Rank()

	switch {
	case rank == lastRank:
		heap.Push(pq, rankedStudent)
		return true
	case rank > lastRank:
		if !isSublist {
			strategy.findAndRemoveFromOthers(pq, []*Student{rankedStudent.Student()})
		}
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
