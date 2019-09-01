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
	SetIsFull(bool)
}

type JointCourse interface {
	AvailableSpots() int
	RegisterCourse(Course)
	Apply(RankedStudent) bool
}

type RankedStudent interface {
	Student() Student
	Rank() int
}

type ApplyStrategy interface {
	SetJointCourse(JointCourse)
	Apply(RankedStudent) bool
}
