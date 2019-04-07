package mapper

type Course struct {
	id         string
	jointId    string
	name       string
	project    string
	faculty    string
	university string
	limit      uint16
}

type Student struct {
	applicationId    string
	citizenId        string
	title            string
	firstName        string
	lastName         string
	phone            string
	email            string
	appliedCourseIds [6]string
}

type Ranking struct {
	courseId  string
	citizenId string
	rank      uint8
	round     uint8
}
