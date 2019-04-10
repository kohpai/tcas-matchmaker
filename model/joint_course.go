package model

import "fmt"

type JointCourse struct {
	id             string
	availableSpots uint16
	courses        []*Course
}

func NewJointCourse(id string, availableSpots uint16) *JointCourse {
	courses := make([]*Course, 0)

	return &JointCourse{
		id,
		availableSpots,
		courses,
	}
}

func (jointCourse *JointCourse) Id() string {
	return jointCourse.id
}

func (jointCourse *JointCourse) AvailableSpots() uint16 {
	return jointCourse.availableSpots
}

func (jointCourse *JointCourse) Courses() []*Course {
	return jointCourse.courses
}

func (jointCourse *JointCourse) Apply() bool {
	if jointCourse.availableSpots == 0 {
		return false
	}

	jointCourse.availableSpots -= 1

	if jointCourse.availableSpots == 0 {
		for _, course := range jointCourse.courses {
			course.isFull = true
		}
	}

	return true
}

func (jointCourse *JointCourse) String() string {
	return fmt.Sprintf(
		"{Course ID: %s, Available Spots: %d}",
		jointCourse.id,
		jointCourse.availableSpots,
	)
}
