package common

type Ranking map[string]int

type Student interface {
	CitizenId() string
	ClearCourse()
	SetCourse(Course)
}

type Course interface {
	Id() string
	Apply(Student) bool
	Ranking() Ranking
}
