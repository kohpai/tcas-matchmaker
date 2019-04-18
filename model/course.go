package model

import (
	"container/heap"
	"log"
)

type Ranking map[string]uint16
type RankCount map[uint16]uint16

type Course struct {
	id          string
	isFull      bool
	condition   Condition
	jointCourse *JointCourse
	ranking     Ranking
	rankCount   RankCount
	students    PriorityQueue
}

func NewCourse(
	id string,
	condition Condition,
	jointCourse *JointCourse,
	ranking Ranking,
) *Course {
	isFull := jointCourse.AvailableSpots() == 0
	course := &Course{
		id,
		isFull,
		condition,
		jointCourse,
		ranking,
		make(RankCount),
		PriorityQueue{},
	}

	jointCourse.RegisterCourse(course)
	return course
}

func (course *Course) Id() string {
	return course.id
}

func (course *Course) IsFull() bool {
	return course.isFull
}

func (course *Course) JointCourse() *JointCourse {
	return course.jointCourse
}

func (course *Course) Ranking() Ranking {
	return course.ranking
}

func (course *Course) Students() PriorityQueue {
	return course.students
}

func (course *Course) Apply(s *Student) bool {
	rank := course.ranking[s.CitizenId()]
	if rank == 0 {
		return false
	}

	course.rankCount[rank] += 1
	rankedStudent := &RankedStudent{
		s, rank, 0,
	}

	heap.Push(&course.students, rankedStudent)

	rankedStudent.student.SetCourse(course)

	conditions := Conditions()
	switch course.condition {
	case conditions.AllowAll:
		return course.applyAll(rankedStudent)
	case conditions.AllowSome:
		return course.applySome(rankedStudent)
	}

	return true
}

func (course *Course) applyAll(rankedStudent *RankedStudent) bool {
	rank := rankedStudent.rank
	if course.rankCount[rank] > 1 {
		return true
	}

	if course.isFull {
		rs := heap.Pop(&course.students).(*RankedStudent)
		rs.student.ClearCourse()
		return rankedStudent != rs
	}

	// @ASSERTION, this shouldn't happen
	if !course.jointCourse.Apply() {
		log.Println("Apply returns false")
	}

	return true
}

func (course *Course) applySome(rankedStudent *RankedStudent) bool {
	if course.isFull {
		rs := heap.Pop(&course.students).(*RankedStudent)
		rs.student.ClearCourse()
		return rankedStudent != rs
	}

	// @ASSERTION, this shouldn't happen
	if !course.jointCourse.Apply() {
		log.Println("Apply returns false")
	}

	return true
}
