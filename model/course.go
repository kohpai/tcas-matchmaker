package model

import (
	"log"
)

type Ranking map[string]uint16

type Course struct {
	id          string
	isFull      bool
	jointCourse *JointCourse
	ranking     Ranking
	students    *PriorityQueue
}

func NewCourse(id string, jointCourse *JointCourse, ranking Ranking) *Course {
	students := &PriorityQueue{}
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

func (course *Course) Students() *PriorityQueue {
	return course.students
}

func (course *Course) Apply(s *Student) bool {
	if course.isFull {
		return false
	}

	// @ASSERTION, this shouldn't happen
	if !course.jointCourse.Apply() {
		log.Println("Apply returns false")
	}

	rankedStudent := &RankedStudent{
		s, course.ranking[s.CitizenId()], nil, nil,
	}

	course.students.Push(rankedStudent)
	return true
}
