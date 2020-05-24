package model

import (
	"fmt"
)

type Ranking map[string]float64
type RankCount map[uint16]uint16

type Course struct {
	id          string
	ranking     Ranking
	jointCourse *JointCourse
}

func NewCourse(
	id string,
	jointCourse *JointCourse,
	ranking Ranking,
) *Course {
	course := &Course{
		id,
		ranking,
		jointCourse,
	}

	jointCourse.RegisterCourse(course)
	return course
}

func (course *Course) Id() string {
	return course.id
}

func (course *Course) JointCourse() *JointCourse {
	return course.jointCourse
}

func (course *Course) Ranking() Ranking {
	return course.ranking
}

func (course *Course) Apply(student *Student) bool {
	rank := course.ranking[student.CitizenId()]
	if rank == 0 {
		return false
	}

	rankedStudent := &RankedStudent{
		student,
		rank,
		0,
	}

	if !course.jointCourse.Apply(rankedStudent) {
		return false
	}

	student.SetCourse(course)
	return true
}

func (course *Course) String() string {
	return fmt.Sprintf(
		"{id: %v, ranking: %v}",
		course.id,
		course.ranking,
	)
}
