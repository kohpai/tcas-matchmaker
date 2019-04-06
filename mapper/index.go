package mapper

type University struct {
	id   string
	name string
}

type Course struct {
	id      string
	name    string
	faculty string
	project string
	jointId string
	limit   uint16
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
