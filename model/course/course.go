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

	jointCourse.courses = append(jointCourse.courses, course)

	return course
}

func (course *Course) IsFull() bool {
	return course.isFull
}
