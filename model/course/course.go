package course

import (
	"log"

	"github.com/kohpai/tcas-3rd-round-resolver/model/student"
)

type Course struct {
	id          string
	isFull      bool
	jointCourse *JointCourse
	students    []*student.Student
}

func NewCourse(id string, jointCourse *JointCourse) *Course {
	students := make([]*student.Student, 0)
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

func (course *Course) Apply(s *student.Student) bool {
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
