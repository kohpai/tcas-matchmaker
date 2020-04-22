package model

import (
	"container/heap"
	"log"
)

type ApplyStrategy interface {
	SetJointCourse(*JointCourse)
	Apply(*RankedStudent) bool
}

type Metadata struct {
	leastReplicatedRank uint16
}

type BaseStrategy struct {
	maleMetadata      *Metadata
	femaleMetadata    *Metadata
	formalMetadata    *Metadata
	interMetadata     *Metadata
	vocatMetadata     *Metadata
	nonFormalMetadata *Metadata
	jointCourse       *JointCourse
}

func NewApplyStrategy(condition Condition, exceedLimit uint16) ApplyStrategy {
	base := BaseStrategy{
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
	}

	conditions := Conditions()

	switch condition {
	case conditions.AllowAll:
		return &AllowAllStrategy{
			base,
		}
	case conditions.DenyAll:
		return &DenyAllStrategy{
			base,
			0,
		}
	case conditions.AllowSome:
		return &AllowSomeStrategy{
			base,
			0,
			exceedLimit,
		}
	}

	return &base
}

func (strategy *BaseStrategy) SetJointCourse(jc *JointCourse) {
	strategy.jointCourse = jc

	if jc.MaleQ() != nil {
		strategy.maleMetadata = &Metadata{0}
	}

	if jc.FemaleQ() != nil {
		strategy.femaleMetadata = &Metadata{0}
	}

	if jc.FormalQ() != nil {
		strategy.formalMetadata = &Metadata{0}
	}

	if jc.InterQ() != nil {
		strategy.interMetadata = &Metadata{0}
	}

	if jc.VocatQ() != nil {
		strategy.vocatMetadata = &Metadata{0}
	}

	if jc.NonFormalQ() != nil {
		strategy.nonFormalMetadata = &Metadata{0}
	}
}

func (strategy *BaseStrategy) Apply(rankedStudent *RankedStudent) bool {
	jc := strategy.jointCourse
	pq := jc.Students()

	if !pq.IsFull() {
		// rejected, admitted, or nothing
		if !strategy.applySublist(rankedStudent) {
			return false
		}
		heap.Push(pq, rankedStudent)
		pq.DecSpots()
		return true
	}

	heap.Push(pq, rankedStudent)
	rs := heap.Pop(pq).(*RankedStudent)
	rs.Student().ClearCourse()
	return rankedStudent != rs
}

func (strategy *BaseStrategy) countEdgeReplicas(pq *PriorityQueue) int {
	lastStudent := heap.Pop(pq).(*RankedStudent)
	if pq.Len() < 1 {
		heap.Push(pq, lastStudent)
		return 1
	}

	nextStudent := heap.Pop(pq).(*RankedStudent)
	students := []*RankedStudent{lastStudent}
	for ; lastStudent.Rank() == nextStudent.Rank(); nextStudent = heap.Pop(pq).(*RankedStudent) {
		students = append(students, nextStudent)
		if pq.Len() < 1 {
			break
		}
	}

	for _, rs := range students {
		heap.Push(pq, rs)
	}
	if students[len(students)-1] != nextStudent {
		heap.Push(pq, nextStudent)
	}

	return len(students)
}

func (strategy *BaseStrategy) applySublist(rankedStudent *RankedStudent) bool {
	jc := strategy.jointCourse
	student := rankedStudent.Student()

	genders := Genders()
	gender := student.Gender()

	if pq := jc.MaleQ(); pq != nil && gender == genders.Male {
		if !strategy.applyDenyAll(pq, strategy.maleMetadata, rankedStudent) {
			return false
		}
	}

	if pq := jc.FemaleQ(); pq != nil && gender == genders.Female {
		if !strategy.applyDenyAll(pq, strategy.femaleMetadata, rankedStudent) {
			return false
		}
	}

	programs := Programs()
	program := student.Program()

	if pq := jc.FormalQ(); pq != nil && program == programs.Formal {
		if !strategy.applyAllowAll(pq, strategy.formalMetadata, rankedStudent) {
			return false
		}
	}

	if pq := jc.InterQ(); pq != nil && program == programs.Inter {
		if !strategy.applyAllowAll(pq, strategy.interMetadata, rankedStudent) {
			return false
		}
	}

	if pq := jc.VocatQ(); pq != nil && program == programs.Vocat {
		if !strategy.applyAllowAll(pq, strategy.vocatMetadata, rankedStudent) {
			return false
		}
	}

	if pq := jc.NonFormalQ(); pq != nil && program == programs.NonFormal {
		if !strategy.applyAllowAll(pq, strategy.nonFormalMetadata, rankedStudent) {
			return false
		}
	}

	return true
}

func (strategy *BaseStrategy) applyDenyAll(
	pq *PriorityQueue,
	metadata *Metadata,
	rankedStudent *RankedStudent,
) bool {
	rank := rankedStudent.Rank()
	copiedRs := &RankedStudent{
		rankedStudent.Student(),
		rank,
		0,
	}

	if !pq.IsFull() {
		lrr := metadata.leastReplicatedRank
		if lrr < 1 || rank < lrr {
			heap.Push(pq, copiedRs)
			pq.DecSpots()
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

	studentsBeingRemoved := make([]*Student, 0)
	count := strategy.countEdgeReplicas(pq)
	for i := 0; i < count; i++ {
		rs := heap.Pop(pq).(*RankedStudent)
		student := rs.Student()
		student.ClearCourse()
		studentsBeingRemoved = append(studentsBeingRemoved, student)
		pq.IncSpots()
	}
	strategy.findAndRemoveFromList(strategy.jointCourse.Students(), studentsBeingRemoved)

	if lrr := metadata.leastReplicatedRank; lrr < 1 || lastRank < lrr {
		metadata.leastReplicatedRank = lastRank
	}

	if rank < lastRank {
		heap.Push(pq, copiedRs)
		pq.DecSpots()
	}

	return rank < lastRank
}

func (strategy *BaseStrategy) applyAllowAll(
	pq *PriorityQueue,
	metadata *Metadata,
	rankedStudent *RankedStudent,
) bool {
	rank := rankedStudent.Rank()
	copiedRs := &RankedStudent{
		rankedStudent.Student(),
		rank,
		0,
	}

	if !pq.IsFull() {
		heap.Push(pq, copiedRs)
		pq.DecSpots()
		return true
	}

	tmp := heap.Pop(pq).(*RankedStudent)
	heap.Push(pq, tmp)
	lastRank := tmp.Rank()

	switch {
	case rank == lastRank:
		heap.Push(pq, copiedRs)
		return true
	case rank > lastRank:
		return false
	}

	heap.Push(pq, copiedRs)
	count := strategy.countBeingRemovedReplicas(pq)
	studentsBeingRemoved := make([]*Student, 0)
	for ; count > 0; count-- {
		rs := heap.Pop(pq).(*RankedStudent)
		student := rs.Student()
		student.ClearCourse()
		studentsBeingRemoved = append(studentsBeingRemoved, student)
	}
	strategy.findAndRemoveFromList(strategy.jointCourse.Students(), studentsBeingRemoved)

	return true
}

func (strategy *BaseStrategy) findAndRemoveFromList(pq *PriorityQueue, students []*Student) {
	beingRemovedStudents := make([]*RankedStudent, 0)
	mainList := pq.Students()
	for i := 0; i < len(students); i++ {
		for j := 0; j < len(mainList); j++ {
			if students[i] == mainList[j].Student() {
				beingRemovedStudents = append(beingRemovedStudents, mainList[j])
				pq.IncSpots()
			}
		}
	}

	for _, student := range beingRemovedStudents {
		heap.Remove(pq, student.index)
	}

	// @ASSERTION, this shouldn't happen
	if len(beingRemovedStudents) != len(students) {
		log.Println("Couldn't find all students to be removed")
	}
}
