package model

import (
	"container/heap"
	"log"
)

type Ranking map[string]uint16

type Course struct {
	id          string
	isFull      bool
	jointCourse *JointCourse
	ranking     Ranking
	students    PriorityQueue
}

func NewCourse(id string, jointCourse *JointCourse, ranking Ranking) *Course {
	students := PriorityQueue{}
	course := &Course{
		id,
		false,
		jointCourse,
		ranking,
		students,
	}

	if jointCourse.AvailableSpots() == 0 {
		course.isFull = true
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
	rankedStudent := &RankedStudent{
		s, course.ranking[s.CitizenId()], 0,
	}

	heap.Push(&course.students, rankedStudent)

	rankedStudent.student.SetCourse(course)

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
