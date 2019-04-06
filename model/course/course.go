package course

type Course struct {
	id          string
	isFull      bool
	jointCourse *JointCourse
}

func NewCourse(id string, jointCourse *JointCourse) *Course {
	course := &Course{
		id,
		false,
		jointCourse,
	}

	if jointCourse.availableSpots == 0 {
		course.isFull = true
	}

	jointCourse.courses = append(jointCourse.courses, course)

	return course
}
