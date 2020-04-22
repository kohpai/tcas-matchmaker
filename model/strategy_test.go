package model

import (
	"testing"
)

func TestApply_AllowAll_AdmitAll(t *testing.T) {
	strategy := NewApplyStrategy(Conditions().AllowAll, 0)
	jointCourse := NewJointCourse("1234", NewAvailableSpots(4, 0, 0, 0, 0, 0, 0), strategy)
	ranking := Ranking{
		"1352": 1,
		"1351": 1,
		"1350": 2,
		"1349": 3,
		"1348": 3,
		"1347": 3,
	}
	course := NewCourse("1234", jointCourse, ranking)

	genders := Genders()
	ss := []*Student{
		NewStudent("1347", genders.Male),
		NewStudent("1348", genders.Male),
		NewStudent("1349", genders.Male),
		NewStudent("1350", genders.Male),
		NewStudent("1351", genders.Male),
		NewStudent("1352", genders.Male),
	}

	for _, s := range ss {
		course.Apply(s)
	}

	if students := jointCourse.Students().Students(); len(students) != len(ss) {
		t.Error("Not all students are admitted", students)
	}
}

func TestApply_AllowAllNoReplicas_AdmitNone(t *testing.T) {
	strategy := NewApplyStrategy(Conditions().AllowAll, 0)
	jointCourse := NewJointCourse("1234", NewAvailableSpots(5, 0, 0, 0, 0, 0, 0), strategy)
	ranking := Ranking{
		"1352": 1,
		"1351": 2,
		"1350": 2,
		"1349": 3,
		"1348": 4,
		"1347": 5,
	}
	course := NewCourse("1234", jointCourse, ranking)

	genders := Genders()
	ss := []*Student{
		NewStudent("1347", genders.Male),
		NewStudent("1348", genders.Male),
		NewStudent("1349", genders.Male),
		NewStudent("1350", genders.Male),
		NewStudent("1351", genders.Male),
		NewStudent("1352", genders.Male),
	}

	for _, s := range ss {
		course.Apply(s)
	}

	if students := jointCourse.Students().Students(); len(students) != 5 {
		t.Error("Not all students are admitted", students)
	}
}

func TestApply_NoCondition_DuplicatedStudentsAreNotAdmitted(t *testing.T) {
	strategy := NewApplyStrategy("", 0)
	jointCourse := NewJointCourse("1234", NewAvailableSpots(3, 0, 0, 0, 0, 0, 0), strategy)
	ranking := Ranking{
		"1352": 1,
		"1351": 1,
		"1350": 2,
		"1349": 3,
		"1348": 3,
		"1347": 3,
	}
	course := NewCourse("1234", jointCourse, ranking)

	genders := Genders()
	ss := []*Student{
		NewStudent("1347", genders.Male),
		NewStudent("1348", genders.Male),
		NewStudent("1349", genders.Male),
		NewStudent("1350", genders.Male),
		NewStudent("1351", genders.Male),
		NewStudent("1352", genders.Male),
	}

	for _, s := range ss {
		course.Apply(s)
	}

	if students := jointCourse.Students().Students(); len(students) != 3 {
		t.Error("Wrong number of students admitted", students)
	}
}

func TestApply_DenyAll1_NoDuplicatedStudentsAreAdmitted(t *testing.T) {
	strategy := NewApplyStrategy(Conditions().DenyAll, 0)
	jointCourse := NewJointCourse("1234", NewAvailableSpots(4, 0, 0, 0, 0, 0, 0), strategy)
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

	genders := Genders()
	ss := []*Student{
		NewStudent("1347", genders.Male),
		NewStudent("1348", genders.Male),
		NewStudent("1349", genders.Male),
		NewStudent("1350", genders.Male),
		NewStudent("1351", genders.Male),
		NewStudent("1352", genders.Male),
		NewStudent("1353", genders.Male),
		NewStudent("1354", genders.Male),
	}

	for _, s := range ss {
		course.Apply(s)
	}

	if students := jointCourse.Students().Students(); len(students) != 2 {
		t.Error("Wrong number of students", students)
	}
}

func TestApply_DenyAll2_NoDuplicatedStudentsAreAdmitted(t *testing.T) {
	strategy := NewApplyStrategy(Conditions().DenyAll, 0)
	jointCourse := NewJointCourse("1234", NewAvailableSpots(3, 0, 0, 0, 0, 0, 0), strategy)
	ranking := Ranking{
		"1354": 1,
		"1353": 2,
		"1352": 2,
		"1351": 2,
	}
	course := NewCourse("1234", jointCourse, ranking)

	genders := Genders()
	ss := []*Student{
		NewStudent("1351", genders.Male),
		NewStudent("1352", genders.Male),
		NewStudent("1353", genders.Male),
		NewStudent("1354", genders.Male),
	}

	for _, s := range ss {
		course.Apply(s)
	}

	if students := jointCourse.Students().Students(); len(students) != 1 {
		t.Error("Wrong number of students", students)
	}
}

func TestApply_AllowSomeNotExceedLimit_AdmitAll(t *testing.T) {
	strategy := NewApplyStrategy(Conditions().AllowSome, 1)
	jointCourse := NewJointCourse("1234", NewAvailableSpots(3, 0, 0, 0, 0, 0, 0), strategy)
	ranking := Ranking{
		"1354": 1,
		"1353": 2,
		"1352": 2,
		"1351": 2,
	}
	course := NewCourse("1234", jointCourse, ranking)

	genders := Genders()
	ss := []*Student{
		NewStudent("1351", genders.Male),
		NewStudent("1352", genders.Male),
		NewStudent("1353", genders.Male),
		NewStudent("1354", genders.Male),
	}

	for _, s := range ss {
		course.Apply(s)
	}

	if students := jointCourse.Students().Students(); len(students) != 4 {
		t.Error("Wrong number of students", students)
	}
}

func TestApply_AllowSomeExceedLimit_AdmitNone(t *testing.T) {
	strategy := NewApplyStrategy(Conditions().AllowSome, 2)
	jointCourse := NewJointCourse("1234", NewAvailableSpots(3, 0, 0, 0, 0, 0, 0), strategy)
	ranking := Ranking{
		"1354": 1,
		"1353": 1,
		"1352": 2,
		"1351": 3,
		"1350": 3,
		"1349": 3,
	}
	course := NewCourse("1234", jointCourse, ranking)

	genders := Genders()
	ss := []*Student{
		NewStudent("1349", genders.Male),
		NewStudent("1350", genders.Male),
		NewStudent("1351", genders.Male),
		NewStudent("1352", genders.Male),
		NewStudent("1353", genders.Male),
		NewStudent("1354", genders.Male),
	}

	for _, s := range ss {
		course.Apply(s)
	}

	if students := jointCourse.Students().Students(); len(students) != 3 {
		t.Error("Wrong number of students", students)
	}
}

func TestApply_Gender1_NoDuplicatedStudentsAreAdmitted(t *testing.T) {
	strategy := NewApplyStrategy("", 0)
	availableSpots := NewAvailableSpots(6, 3, 0, 0, 0, 0, 0)
	jointCourse := NewJointCourse("1234", availableSpots, strategy)
	ranking := Ranking{
		"1352": 3,
		"1351": 3,
		"1350": 2,
		"1349": 2,
		"1348": 2,
		"1347": 1,
	}
	course := NewCourse("1234", jointCourse, ranking)

	genders := Genders()
	ss := []*Student{
		NewStudent("1347", genders.Male),
		NewStudent("1348", genders.Male),
		NewStudent("1349", genders.Male),
		NewStudent("1350", genders.Male),
		NewStudent("1351", genders.Female),
		NewStudent("1352", genders.Female),
	}

	for _, s := range ss {
		course.Apply(s)
	}

	if pq := jointCourse.Students(); pq.Len() != 3 {
		t.Error("Wrong number of students", pq.Students())
	}
}
