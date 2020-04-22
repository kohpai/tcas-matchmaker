package model

import (
	"container/heap"
	"fmt"
)

type JointCourse struct {
	id        string
	courses   []*Course
	main      *PriorityQueue
	male      *PriorityQueue
	female    *PriorityQueue
	formal    *PriorityQueue
	inter     *PriorityQueue
	vocat     *PriorityQueue
	nonFormal *PriorityQueue
	strategy  ApplyStrategy
}

func NewJointCourse(
	id string,
	availableSpots *AvailableSpots,
	strategy ApplyStrategy,
) *JointCourse {
	courses := make([]*Course, 0)
	var main *PriorityQueue
	var male *PriorityQueue
	var female *PriorityQueue
	var formal *PriorityQueue
	var inter *PriorityQueue
	var vocat *PriorityQueue
	var nonFormal *PriorityQueue

	if num := availableSpots.main; num != 0 {
		main = &PriorityQueue{num, []*RankedStudent{}}
		heap.Init(main)
	}
	if num := availableSpots.male; num != 0 {
		male = &PriorityQueue{num, []*RankedStudent{}}
		heap.Init(male)
	}
	if num := availableSpots.female; num != 0 {
		female = &PriorityQueue{num, []*RankedStudent{}}
		heap.Init(female)
	}
	if num := availableSpots.formal; num != 0 {
		formal = &PriorityQueue{num, []*RankedStudent{}}
		heap.Init(formal)
	}
	if num := availableSpots.inter; num != 0 {
		inter = &PriorityQueue{num, []*RankedStudent{}}
		heap.Init(inter)
	}
	if num := availableSpots.vocat; num != 0 {
		vocat = &PriorityQueue{num, []*RankedStudent{}}
		heap.Init(vocat)
	}
	if num := availableSpots.nonFormal; num != 0 {
		nonFormal = &PriorityQueue{num, []*RankedStudent{}}
		heap.Init(nonFormal)
	}

	jointCourse := &JointCourse{
		id,
		courses,
		main,
		male,
		female,
		formal,
		inter,
		vocat,
		nonFormal,
		strategy,
	}

	strategy.SetJointCourse(jointCourse)
	return jointCourse
}

func (jointCourse *JointCourse) Id() string {
	return jointCourse.id
}

func (jointCourse *JointCourse) Courses() []*Course {
	return jointCourse.courses
}

func (jointCourse *JointCourse) Students() *PriorityQueue {
	return jointCourse.main
}

func (jointCourse *JointCourse) MaleQ() *PriorityQueue {
	return jointCourse.male
}

func (jointCourse *JointCourse) FemaleQ() *PriorityQueue {
	return jointCourse.female
}

func (jointCourse *JointCourse) FormalQ() *PriorityQueue {
	return jointCourse.formal
}

func (jointCourse *JointCourse) InterQ() *PriorityQueue {
	return jointCourse.inter
}

func (jointCourse *JointCourse) VocatQ() *PriorityQueue {
	return jointCourse.vocat
}

func (jointCourse *JointCourse) NonFormalQ() *PriorityQueue {
	return jointCourse.nonFormal
}

func (jointCourse *JointCourse) RegisterCourse(course *Course) {
	jointCourse.courses = append(jointCourse.courses, course)
}

func (jointCourse *JointCourse) Apply(rankedStudent *RankedStudent) bool {
	if jointCourse.main.Limit() == 0 {
		return false
	}
	return jointCourse.strategy.Apply(rankedStudent)
}

func (jointCourse *JointCourse) String() string {
	return fmt.Sprintf(
		"{id: %s, courses: %v}",
		jointCourse.id,
		jointCourse.courses,
	)
}
