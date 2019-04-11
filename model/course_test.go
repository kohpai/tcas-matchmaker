package model

import (
	"testing"
)

func TestNewCourse_Always_ReturnsCourse(t *testing.T) {
	jointCourse := NewJointCourse("1234", 10)
	course := NewCourse("1234", jointCourse, nil)

	if course.Id() != "1234" {
		t.Error("Course ID not matched", course)
	}

	if course.JointCourse() != jointCourse {
		t.Error("Joint course is incorrect", course)
	}
}

func TestIsFull_AvailableSpotsGreaterThanZero_ReturnsFalse(t *testing.T) {
	jointCourse := NewJointCourse("1234", 1)
	course := NewCourse("1234", jointCourse, nil)

	if course.IsFull() {
		t.Error("Course is full", course)
	}
}

func TestIsFull_AvailableSpotsIsReachingZero_ReturnsTrue(t *testing.T) {
	jointCourse := NewJointCourse("1234", 1)
	course := NewCourse("1234", jointCourse, nil)

	jointCourse.Apply()

	if !course.IsFull() {
		t.Error("Course is NOT full", course)
	}
}

func TestIsFull_AvailableSpotsIsAlreadyZero_ReturnsTrue(t *testing.T) {
	jointCourse := NewJointCourse("1234", 0)
	course := NewCourse("1234", jointCourse, nil)

	if !course.IsFull() {
		t.Error("Course is NOT full", course)
	}
}

func TestApply_CourseIsNotFull_ReturnsTrue(t *testing.T) {
	jointCourse := NewJointCourse("1234", 1)
	ranking := map[string]uint16{
		"1349": 1,
	}
	course := NewCourse("1234", jointCourse, ranking)

	s := NewStudent("1349")

	if !course.Apply(s) {
		t.Error("Apply returns false", course)
	}
}

func TestApply_OneSpotLeft_CourseIsFull(t *testing.T) {
	jointCourse := NewJointCourse("1234", 1)
	ranking := map[string]uint16{
		"1349": 1,
	}
	course := NewCourse("1234", jointCourse, ranking)

	s := NewStudent("1349")
	course.Apply(s)

	if !course.IsFull() {
		t.Error("Course is NOT full", course)
	}
}

func TestApply_MoreSpotsLeft_StudentsAreEnrolled(t *testing.T) {
	jointCourse := NewJointCourse("1234", 3)
	ranking := map[string]uint16{
		"1351": 1,
		"1350": 2,
		"1349": 3,
	}
	course := NewCourse("1234", jointCourse, ranking)

	ss := []*Student{
		NewStudent("1349"),
		NewStudent("1350"),
		NewStudent("1351"),
	}

	for _, s := range ss {
		course.Apply(s)
	}

	regStudents := course.Students()

	if s := regStudents.Pop(); ss[2] != s.student {
		t.Error("Student is not matched,", s)
	}
	if s := regStudents.Pop(); ss[1] != s.student {
		t.Error("Student is not matched,", s)
	}
	if s := regStudents.Pop(); ss[0] != s.student {
		t.Error("Student is not matched,", s)
	}
}
