package model

import "testing"

func TestApply_AllowAll_AllStudentsAreAdmitted(t *testing.T) {
	jointCourse := NewJointCourse("1234", 3)
	ranking := Ranking{
		"1352": 1,
		"1351": 1,
		"1350": 2,
		"1349": 3,
		"1348": 3,
		"1347": 3,
	}
	course := NewCourse("1234", Conditions().AllowAll, jointCourse, ranking)

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

	if len(course.Students().Students()) != len(ss) {
		t.Error("Not all students are admitted", course)
	}
}

func TestApply_NoCondition_DuplicatedStudentsAreNotAdmitted(t *testing.T) {
	jointCourse := NewJointCourse("1234", 3)
	ranking := Ranking{
		"1352": 1,
		"1351": 1,
		"1350": 2,
		"1349": 3,
		"1348": 3,
		"1347": 3,
	}
	course := NewCourse("1234", 0, jointCourse, ranking)

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

	if len(course.Students().Students()) != 3 {
		t.Error("Wrong number of students admitted", course)
	}
}

func TestApply_DenyAll_NoDuplicatedStudentsAreAdmitted(t *testing.T) {
	jointCourse := NewJointCourse("1234", 3)
	ranking := Ranking{
		"1352": 1,
		"1351": 1,
		"1350": 2,
		"1349": 3,
		"1348": 3,
		"1347": 3,
	}
	course := NewCourse("1234", Conditions().DenyAll, jointCourse, ranking)

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

	if len(course.Students().Students()) != 1 {
		t.Error("Students are replicated", course)
	}
}
