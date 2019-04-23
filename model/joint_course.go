package model

import (
	"fmt"
	"log"
)

type JointCourse struct {
	id             string
	limit          uint16
	availableSpots uint16
	courses        []*Course
	students       *PriorityQueue
	strategy       ApplyStrategy
}

func NewJointCourse(
	id string,
	availableSpots uint16,
	strategy ApplyStrategy,
) *JointCourse {
	courses := make([]*Course, 0)

	jointCourse := &JointCourse{
		id,
		availableSpots,
		availableSpots,
		courses,
		&PriorityQueue{
			[]*RankedStudent{},
		},
		strategy,
	}

	strategy.SetJointCourse(jointCourse)
	return jointCourse
}

func (jointCourse *JointCourse) Id() string {
	return jointCourse.id
}

func (jointCourse *JointCourse) Limit() uint16 {
	return jointCourse.limit
}

func (jointCourse *JointCourse) AvailableSpots() uint16 {
	return jointCourse.availableSpots
}

func (jointCourse *JointCourse) Courses() []*Course {
	return jointCourse.courses
}

func (jointCourse *JointCourse) Students() *PriorityQueue {
	return jointCourse.students
}

func (jointCourse *JointCourse) IsFull() bool {
	return jointCourse.availableSpots == 0
}

func (jointCourse *JointCourse) RegisterCourse(course *Course) {
	jointCourse.courses = append(jointCourse.courses, course)
}

func (jointCourse *JointCourse) IncSpots() {
	// @ASSERTION, this shouldn't happen
	if jointCourse.availableSpots >= jointCourse.limit {
		log.Println("available spots is more than limit")
		return
	}

	jointCourse.availableSpots += 1

	for _, course := range jointCourse.courses {
		course.SetIsFull(false)
	}
}

func (jointCourse *JointCourse) DecSpots() {
	if jointCourse.availableSpots == 0 {
		return
	}

	jointCourse.availableSpots -= 1

	if jointCourse.availableSpots == 0 {
		for _, course := range jointCourse.courses {
			course.SetIsFull(true)
		}
	}
}

func (jointCourse *JointCourse) Apply(rankedStudent *RankedStudent) bool {
	return jointCourse.strategy.Apply(rankedStudent)
}

func (jointCourse *JointCourse) String() string {
	return fmt.Sprintf(
		"{id: %s, availabelSpots: %d, courses: %v}",
		jointCourse.id,
		jointCourse.availableSpots,
		jointCourse.courses,
	)
}
