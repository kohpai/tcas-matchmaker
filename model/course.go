package model

import (
	"container/heap"
	"fmt"
)

type Ranking map[string]uint16
type RankCount map[uint16]uint16

type Course struct {
	id          string
	isFull      bool
	ranking     Ranking
	strategy    ApplyStrategy
	jointCourse *JointCourse
	students    *PriorityQueue
}

func NewCourse(
	id string,
	condition Condition,
	jointCourse *JointCourse,
	ranking Ranking,
) *Course {
	strategy := NewApplyStrategy(condition)
	ranking = strategy.InitRanking(ranking)

	isFull := jointCourse.AvailableSpots() == 0
	course := &Course{
		id,
		isFull,
		ranking,
		strategy,
		jointCourse,
		&PriorityQueue{
			[]*RankedStudent{},
		},
	}

	strategy.SetCourse(course)
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

func (course *Course) Students() *PriorityQueue {
	return course.students
}

func (course *Course) Apply(s *Student) bool {
	rank := course.ranking[s.CitizenId()]
	if rank == 0 {
		return false
	}

	strategy := course.strategy
	strategy.IncRankCount(rank)
	rankedStudent := &RankedStudent{
		s, rank, 0,
	}

	heap.Push(course.students, rankedStudent)
	rankedStudent.student.SetCourse(course)

	return strategy.Apply(rankedStudent)
}

func (course *Course) String() string {
	return fmt.Sprintf(
		"{id: %v, isFull: %v, ranking: %v, students: %v}",
		course.id,
		course.isFull,
		course.ranking,
		course.students.students,
	)
}
