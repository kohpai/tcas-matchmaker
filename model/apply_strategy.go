package model

import (
	"container/heap"
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
	student := rs.Student()
	student.ClearCourse()
	strategy.findAndRemoveFromOthers(pq, []*Student{student})
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

	var pq *PriorityQueue
	var metadata *Metadata

	switch gender {
	case genders.Male:
		pq, metadata = jc.MaleQ(), strategy.maleMetadata
	case genders.Female:
		pq, metadata = jc.FemaleQ(), strategy.femaleMetadata
	}

	if pq != nil && !strategy.applyDenyAll(pq, metadata, rankedStudent) {
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

	if pq != nil && !strategy.applyAllowAll(pq, metadata, rankedStudent) {
		return false
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
	strategy.findAndRemoveFromOthers(pq, studentsBeingRemoved)

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
	strategy.findAndRemoveFromOthers(pq, studentsBeingRemoved)

	return true
}

func (strategy *BaseStrategy) findAndRemoveFromOthers(pq *PriorityQueue, students []*Student) {
	jc := strategy.jointCourse

	if q := jc.Students(); q != nil && pq != q {
		strategy.findAndRemoveFromList(q, students)
	}

	if q := jc.MaleQ(); q != nil && pq != q {
		strategy.findAndRemoveFromList(q, students)
	}

	if q := jc.FemaleQ(); q != nil && pq != q {
		strategy.findAndRemoveFromList(q, students)
	}

	if q := jc.FormalQ(); q != nil && pq != q {
		strategy.findAndRemoveFromList(q, students)
	}

	if q := jc.InterQ(); q != nil && pq != q {
		strategy.findAndRemoveFromList(q, students)
	}

	if q := jc.VocatQ(); q != nil && pq != q {
		strategy.findAndRemoveFromList(q, students)
	}

	if q := jc.NonFormalQ(); q != nil && pq != q {
		strategy.findAndRemoveFromList(q, students)
	}
}

func (strategy *BaseStrategy) findAndRemoveFromList(pq *PriorityQueue, students []*Student) {
	beingRemovedStudents := make([]*RankedStudent, 0)
	mainList := pq.Students()
	for _, student := range students {
		for _, rs := range mainList {
			if student == rs.Student() {
				beingRemovedStudents = append(beingRemovedStudents, rs)
			}
		}
	}

	for _, student := range beingRemovedStudents {
		heap.Remove(pq, student.index)
		pq.IncSpots()
	}
}
