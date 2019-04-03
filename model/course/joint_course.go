package course

type JointCourse struct {
	id             string
	availableSpots uint16
	courses        []Course
}

// func (jointCourse *JointCourse) Apply() bool {
// 	jointCourse.availableSpots -= 1
// 	return true
// }

func NewJointCourse(id string, availableSpots uint16) *JointCourse {
	return &JointCourse{
		id:             id,
		availableSpots: availableSpots,
		courses:        make([]Course, 0),
	}
}
