package course

type Course struct {
	id          string
	isFull      bool
	jointCourse *JointCourse
}

func NewCourse(id string, jointCourse *JointCourse) *Course {
	return &Course{
		id,
		false,
		jointCourse,
	}
}

func (course *Course) IsFull() bool {
	return course.isFull
}
