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
	Students() PriorityQueue
	Limit() int
	IsFull() bool
	DecSpots()
	IncSpots()
}

type PriorityQueue interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Push(x interface{})
	Pop() interface{}
	Students() []RankedStudent
}

type RankedStudent interface {
	Student() Student
	Rank() int
	SetIndex(int)
}

type ApplyStrategy interface {
	SetJointCourse(JointCourse)
	Apply(RankedStudent) bool
}
