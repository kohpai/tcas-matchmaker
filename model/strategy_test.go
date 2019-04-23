package model

import (
	"testing"
)

func TestApply_AllowAll_AllStudentsAreAdmitted(t *testing.T) {
	strategy := NewApplyStrategy(Conditions().AllowAll)
	jointCourse := NewJointCourse("1234", 4, strategy)
	ranking := Ranking{
		"1352": 1,
		"1351": 1,
		"1350": 2,
		"1349": 3,
		"1348": 3,
		"1347": 3,
	}
	course := NewCourse("1234", jointCourse, ranking)

	ss := []*Student{
		NewStudent("1347"),
		NewStudent("1348"),
		NewStudent("1349"),
		NewStudent("1350"),
		NewStudent("1351"),
		NewStudent("1352"),
	}

	for _, s := range ss {
		course.Apply(s)
	}

	if students := jointCourse.Students().Students(); len(students) != len(ss) {
		t.Error("Not all students are admitted", students)
	}
}

func TestApply_NoCondition_DuplicatedStudentsAreNotAdmitted(t *testing.T) {
	strategy := NewApplyStrategy(0)
	jointCourse := NewJointCourse("1234", 3, strategy)
	ranking := Ranking{
		"1352": 1,
		"1351": 1,
		"1350": 2,
		"1349": 3,
		"1348": 3,
		"1347": 3,
	}
	course := NewCourse("1234", jointCourse, ranking)

	ss := []*Student{
		NewStudent("1347"),
		NewStudent("1348"),
		NewStudent("1349"),
		NewStudent("1350"),
		NewStudent("1351"),
		NewStudent("1352"),
	}

	for _, s := range ss {
		course.Apply(s)
	}

	if students := jointCourse.Students().Students(); len(students) != 3 {
		t.Error("Wrong number of students admitted", students)
	}
}

func TestApply_DenyAll_NoDuplicatedStudentsAreAdmitted(t *testing.T) {
	strategy := NewApplyStrategy(Conditions().DenyAll)
	jointCourse := NewJointCourse("1234", 4, strategy)
	ranking := Ranking{
		"1354": 1,
		"1353": 1,
		"1352": 2,
		"1351": 2,
		"1350": 2,
		"1349": 3,
		"1348": 3,
		"1347": 3,
	}
	course := NewCourse("1234", jointCourse, ranking)

	ss := []*Student{
		NewStudent("1347"),
		NewStudent("1348"),
		NewStudent("1349"),
		NewStudent("1350"),
		NewStudent("1351"),
		NewStudent("1352"),
		NewStudent("1353"),
		NewStudent("1354"),
	}

	for _, s := range ss {
		course.Apply(s)
	}

	if students := jointCourse.Students().Students(); len(students) != 2 {
		t.Error("Wrong number of students", students)
	}
}
