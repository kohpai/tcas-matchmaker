package model

import (
	"container/heap"
	"testing"
)

func TestNewCourse_Always_ReturnsCourse(t *testing.T) {
	jointCourse := NewJointCourse("1234", 10)
	course := NewCourse("1234", Conditions().AllowAll, jointCourse, nil)

	if course.Id() != "1234" {
		t.Error("Course ID not matched", course)
	}

	if course.JointCourse() != jointCourse {
		t.Error("Joint course is incorrect", course)
	}
}

func TestIsFull_AvailableSpotsGreaterThanZero_ReturnsFalse(t *testing.T) {
	jointCourse := NewJointCourse("1234", 1)
	course := NewCourse("1234", Conditions().AllowAll, jointCourse, nil)

	if course.IsFull() {
		t.Error("Course is full", course)
	}
}

func TestIsFull_AvailableSpotsIsReachingZero_ReturnsTrue(t *testing.T) {
	jointCourse := NewJointCourse("1234", 1)
	course := NewCourse("1234", Conditions().AllowAll, jointCourse, nil)

	jointCourse.Apply()

	if !course.IsFull() {
		t.Error("Course is NOT full", course)
	}
}

func TestIsFull_AvailableSpotsIsAlreadyZero_ReturnsTrue(t *testing.T) {
	jointCourse := NewJointCourse("1234", 0)
	course := NewCourse("1234", Conditions().AllowAll, jointCourse, nil)

	if !course.IsFull() {
		t.Error("Course is NOT full", course)
	}
}

func TestApply_CourseIsNotFull_ReturnsTrue(t *testing.T) {
	jointCourse := NewJointCourse("1234", 1)
	ranking := Ranking{
		"1349": 1,
	}
	course := NewCourse("1234", Conditions().AllowAll, jointCourse, ranking)

	s := NewStudent("1349")

	if !course.Apply(s) {
		t.Error("Apply returns false", course)
	}
}

func TestApply_CourseIsFullAndStudentHasHigherRank_ReturnsTrue(t *testing.T) {
	jointCourse := NewJointCourse("1234", 1)
	ranking := Ranking{
		"1349": 2,
		"1350": 1,
	}
	course := NewCourse("1234", Conditions().AllowAll, jointCourse, ranking)

	ss := []*Student{
		NewStudent("1349"),
		NewStudent("1350"),
	}

	course.Apply(ss[0])
	if !course.Apply(ss[1]) {
		t.Error("Apply returns false", ss[1])
	}

	statuses := ApplicationStatuses()
	if ss[0].ApplicationStatus() != statuses.Pending {
		t.Error("Student has incorrect status", ss[0])
	}

	if ss[1].ApplicationStatus() != statuses.Accepted {
		t.Error("Student has incorrect status", ss[1])
	}
}

func TestApply_CourseIsFullAndStudentHasLowerRank_ReturnsFalse(t *testing.T) {
	jointCourse := NewJointCourse("1234", 1)
	ranking := Ranking{
		"1349": 1,
		"1350": 2,
	}
	course := NewCourse("1234", Conditions().AllowAll, jointCourse, ranking)

	ss := []*Student{
		NewStudent("1349"),
		NewStudent("1350"),
	}

	course.Apply(ss[0])
	if course.Apply(ss[1]) {
		t.Error("Apply returns true", ss[1])
	}

	statuses := ApplicationStatuses()
	if ss[0].ApplicationStatus() != statuses.Accepted {
		t.Error("Student has incorrect status", ss[0])
	}

	if ss[1].ApplicationStatus() != statuses.Pending {
		t.Error("Student has incorrect status", ss[1])
	}
}

func TestApply_OneSpotLeft_CourseIsFull(t *testing.T) {
	jointCourse := NewJointCourse("1234", 1)
	ranking := Ranking{
		"1349": 1,
	}
	course := NewCourse("1234", Conditions().AllowAll, jointCourse, ranking)

	s := NewStudent("1349")
	course.Apply(s)

	if !course.IsFull() {
		t.Error("Course is NOT full", course)
	}
}

func TestApply_MoreSpotsLeft_StudentsAreEnrolled(t *testing.T) {
	jointCourse := NewJointCourse("1234", 3)
	ranking := Ranking{
		"1351": 1,
		"1350": 2,
		"1349": 3,
	}
	course := NewCourse("1234", Conditions().AllowAll, jointCourse, ranking)

	ss := []*Student{
		NewStudent("1349"),
		NewStudent("1350"),
		NewStudent("1351"),
	}

	for _, s := range ss {
		course.Apply(s)
	}

	regStudents := course.Students()

	if s := heap.Pop(regStudents).(*RankedStudent); ss[0] != s.student {
		t.Error("Student is not matched,", s)
	}
	if s := heap.Pop(regStudents).(*RankedStudent); ss[1] != s.student {
		t.Error("Student is not matched,", s)
	}
	if s := heap.Pop(regStudents).(*RankedStudent); ss[2] != s.student {
		t.Error("Student is not matched,", s)
	}
}
