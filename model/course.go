package model

import (
	"log"
)

type Course struct {
	id          string
	isFull      bool
	jointCourse *JointCourse
	students    []*Student
}

func NewCourse(id string, jointCourse *JointCourse) *Course {
	students := make([]*Student, 0)
	course := &Course{
		id,
		false,
		jointCourse,
		students,
	}

	if jointCourse.availableSpots == 0 {
		course.isFull = true
	}

	jointCourse.courses = append(jointCourse.courses, course)

	return course
}

func (course *Course) Apply(s *Student) bool {
	if course.isFull {
		return false
	}

	// @ASSERTION, this shouldn't happen
	if !course.jointCourse.Apply() {
		log.Println("Apply returns false")
	}

	course.students = append(course.students, s)
	return true
}
