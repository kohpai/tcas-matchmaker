package model

import (
	"fmt"

	"github.com/kohpai/tcas-3rd-round-resolver/model/common"
)

type RankCount map[int]int

type Course struct {
	id          string
	isFull      bool
	ranking     common.Ranking
	jointCourse *JointCourse
}

func NewCourse(
	id string,
	jointCourse *JointCourse,
	ranking common.Ranking,
) *Course {
	isFull := jointCourse.AvailableSpots() == 0
	course := &Course{
		id,
		isFull,
		ranking,
		jointCourse,
	}

	jointCourse.RegisterCourse(course)
	return course
}

func (course *Course) Id() string {
	return course.id
}

func (course *Course) IsFull() bool {
	return course.isFull
}

func (course *Course) SetIsFull(isFull bool) {
	course.isFull = isFull
}

func (course *Course) JointCourse() *JointCourse {
	return course.jointCourse
}

func (course *Course) Ranking() common.Ranking {
	return course.ranking
}

func (course *Course) Apply(student common.Student) bool {
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
		"{id: %v, isFull: %v, ranking: %v}",
		course.id,
		course.isFull,
		course.ranking,
	)
}
