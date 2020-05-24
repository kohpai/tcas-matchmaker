package model

import "container/heap"

type DenyAllStrategy struct {
	BaseStrategy
	Metadata
}

func (strategy *DenyAllStrategy) applySublist(rankedStudent *RankedStudent) bool {
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

func (strategy *DenyAllStrategy) apply(
	pq *PriorityQueue,
	metadata *Metadata,
	rankedStudent *RankedStudent,
	isSublist bool,
) bool {
	if !isSublist && !strategy.applySublist(rankedStudent) {
		return false
	}

	rank := rankedStudent.Rank()

	if !pq.IsFull() {
		if lrr := metadata.leastReplicatedRank; metadata.noRank || rank < lrr {
			heap.Push(pq, rankedStudent)
			return true
		}

		return false
	}

	tmp := heap.Pop(pq).(*RankedStudent)
	heap.Push(pq, tmp)
	lastRank := tmp.Rank()
	if rank > lastRank {
		if !isSublist {
			strategy.findAndRemoveFromOthers(pq, []*Student{rankedStudent.Student()})
		}
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

	if lrr := metadata.leastReplicatedRank; metadata.noRank || lastRank < lrr {
		metadata.leastReplicatedRank = lastRank
		metadata.noRank = false
	}

	if rank < lastRank {
		heap.Push(pq, rankedStudent)
		return true
	}

	return false
}

func (strategy *DenyAllStrategy) Apply(rankedStudent *RankedStudent) bool {
	return strategy.apply(strategy.jointCourse.Students(), &strategy.Metadata, rankedStudent, false)
}
