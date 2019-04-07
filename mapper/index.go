package mapper

type Course struct {
	id        string
	jointId   string
	limit     uint16
	condition uint8
	addLimit  uint16
}

type Student struct {
	citizenId        string
	appliedCourseIds [6]string
}

type Ranking struct {
	courseId  string
	citizenId string
	rank      uint8
}
