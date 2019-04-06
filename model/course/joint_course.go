package course

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

func (jointCourse *JointCourse) Apply() bool {
	if jointCourse.availableSpots == 0 {
		return false
	}

	jointCourse.availableSpots -= 1

	return true
}

func (jointCourse *JointCourse) String() string {
	return fmt.Sprintf(
		"{\n\tCourse ID: %s, \n\t Available Spots: %d,\n}",
		jointCourse.id,
		jointCourse.availableSpots,
	)
}
