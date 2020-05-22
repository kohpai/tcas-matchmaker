package model

import (
	"testing"
)

func TestApply_AllowAll_AdmitAll(t *testing.T) {
	strategy := NewApplyStrategy(Conditions().AllowAll, 0)
	jointCourse := NewJointCourse("1234", NewAvailableSpots(4, 0, 0, 0, 0, 0, 0), strategy)
	ranking := Ranking{
		"1352": -123.4,
		"1351": -123.4,
		"1350": -99.8,
		"1349": -50,
		"1348": -50,
		"1347": -50,
	}
	course := NewCourse("1234", jointCourse, ranking)

	programs := Programs()
	genders := Genders()
	ss := []*Student{
		NewStudent("1347", genders.Male, programs.Formal, false),
		NewStudent("1348", genders.Male, programs.Formal, false),
		NewStudent("1349", genders.Male, programs.Formal, false),
		NewStudent("1350", genders.Male, programs.Formal, false),
		NewStudent("1351", genders.Male, programs.Formal, false),
		NewStudent("1352", genders.Male, programs.Formal, false),
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
		"1352": -123.4,
		"1351": -99.8,
		"1350": -99.8,
		"1349": -50,
		"1348": -9,
		"1347": -0.2,
	}
	course := NewCourse("1234", jointCourse, ranking)

	programs := Programs()
	genders := Genders()
	ss := []*Student{
		NewStudent("1347", genders.Male, programs.Formal, false),
		NewStudent("1348", genders.Male, programs.Formal, false),
		NewStudent("1349", genders.Male, programs.Formal, false),
		NewStudent("1350", genders.Male, programs.Formal, false),
		NewStudent("1351", genders.Male, programs.Formal, false),
		NewStudent("1352", genders.Male, programs.Formal, false),
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
		"1352": -123.4,
		"1351": -123.4,
		"1350": -99.8,
		"1349": -50,
		"1348": -50,
		"1347": -50,
	}
	course := NewCourse("1234", jointCourse, ranking)

	programs := Programs()
	genders := Genders()
	ss := []*Student{
		NewStudent("1347", genders.Male, programs.Formal, false),
		NewStudent("1348", genders.Male, programs.Formal, false),
		NewStudent("1349", genders.Male, programs.Formal, false),
		NewStudent("1350", genders.Male, programs.Formal, false),
		NewStudent("1351", genders.Male, programs.Formal, false),
		NewStudent("1352", genders.Male, programs.Formal, false),
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
		"1354": -123.4,
		"1353": -123.4,
		"1352": -99.8,
		"1351": -99.8,
		"1350": -99.8,
		"1349": -50,
		"1348": -50,
		"1347": -50,
	}
	course := NewCourse("1234", jointCourse, ranking)

	programs := Programs()
	genders := Genders()
	ss := []*Student{
		NewStudent("1347", genders.Male, programs.Formal, false),
		NewStudent("1348", genders.Male, programs.Formal, false),
		NewStudent("1349", genders.Male, programs.Formal, false),
		NewStudent("1350", genders.Male, programs.Formal, false),
		NewStudent("1351", genders.Male, programs.Formal, false),
		NewStudent("1352", genders.Male, programs.Formal, false),
		NewStudent("1353", genders.Male, programs.Formal, false),
		NewStudent("1354", genders.Male, programs.Formal, false),
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
		"1354": -123.4,
		"1353": -99.8,
		"1352": -99.8,
		"1351": -99.8,
	}
	course := NewCourse("1234", jointCourse, ranking)

	programs := Programs()
	genders := Genders()
	ss := []*Student{
		NewStudent("1351", genders.Male, programs.Formal, false),
		NewStudent("1352", genders.Male, programs.Formal, false),
		NewStudent("1353", genders.Male, programs.Formal, false),
		NewStudent("1354", genders.Male, programs.Formal, false),
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
		"1351": -99.8,
		"1352": -99.8,
		"1353": -99.8,
		"1354": -123.4,
	}
	course := NewCourse("1234", jointCourse, ranking)

	programs := Programs()
	genders := Genders()
	ss := []*Student{
		NewStudent("1351", genders.Male, programs.Formal, false),
		NewStudent("1352", genders.Male, programs.Formal, false),
		NewStudent("1353", genders.Male, programs.Formal, false),
		NewStudent("1354", genders.Male, programs.Formal, false),
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
		"1354": -123.4,
		"1353": -123.4,
		"1352": -99.8,
		"1351": -50,
		"1350": -50,
		"1349": -50,
	}
	course := NewCourse("1234", jointCourse, ranking)

	programs := Programs()
	genders := Genders()
	ss := []*Student{
		NewStudent("1349", genders.Male, programs.Formal, false),
		NewStudent("1350", genders.Male, programs.Formal, false),
		NewStudent("1351", genders.Male, programs.Formal, false),
		NewStudent("1352", genders.Male, programs.Formal, false),
		NewStudent("1353", genders.Male, programs.Formal, false),
		NewStudent("1354", genders.Male, programs.Formal, false),
	}

	for _, s := range ss {
		course.Apply(s)
	}

	if students := jointCourse.Students().Students(); len(students) != 3 {
		t.Error("Wrong number of students", students)
	}
}

func TestApply_Male_NoDuplicatedStudentsAreAdmitted(t *testing.T) {
	strategy := NewApplyStrategy("", 0)
	availableSpots := NewAvailableSpots(6, 3, 0, 0, 0, 0, 0)
	jointCourse := NewJointCourse("1234", availableSpots, strategy)
	ranking := Ranking{
		"1352": -50,
		"1351": -50,
		"1350": -99.8,
		"1349": -99.8,
		"1348": -99.8,
		"1347": -123.4,
	}
	course := NewCourse("1234", jointCourse, ranking)

	programs := Programs()
	genders := Genders()
	ss := []*Student{
		NewStudent("1347", genders.Male, programs.Formal, false),
		NewStudent("1348", genders.Male, programs.Formal, false),
		NewStudent("1349", genders.Male, programs.Formal, false),
		NewStudent("1350", genders.Male, programs.Formal, false),
		NewStudent("1351", genders.Female, programs.Formal, false),
		NewStudent("1352", genders.Female, programs.Formal, false),
	}

	for _, s := range ss {
		course.Apply(s)
	}

	if pq := jointCourse.Students(); pq.Len() != 3 {
		t.Error("Wrong number of students", pq.Students())
	}
}

func TestApply_Female_NoDuplicatedStudentsAreAdmitted(t *testing.T) {
	strategy := NewApplyStrategy("", 0)
	availableSpots := NewAvailableSpots(6, 0, 3, 0, 0, 0, 0)
	jointCourse := NewJointCourse("1234", availableSpots, strategy)
	ranking := Ranking{
		"1352": -50,
		"1351": -50,
		"1350": -99.8,
		"1349": -99.8,
		"1348": -99.8,
		"1347": -123.4,
	}
	course := NewCourse("1234", jointCourse, ranking)

	programs := Programs()
	genders := Genders()
	ss := []*Student{
		NewStudent("1347", genders.Female, programs.Formal, false),
		NewStudent("1348", genders.Female, programs.Formal, false),
		NewStudent("1349", genders.Female, programs.Formal, false),
		NewStudent("1350", genders.Female, programs.Formal, false),
		NewStudent("1351", genders.Male, programs.Formal, false),
		NewStudent("1352", genders.Male, programs.Formal, false),
	}

	for _, s := range ss {
		course.Apply(s)
	}

	if pq := jointCourse.Students(); pq.Len() != 3 {
		t.Error("Wrong number of students", pq.Students())
	}
}

func TestApply_FormalJustInPlace_AdmitAllReplicas(t *testing.T) {
	strategy := NewApplyStrategy("", 0)
	availableSpots := NewAvailableSpots(6, 0, 0, 3, 0, 0, 0)
	jointCourse := NewJointCourse("1234", availableSpots, strategy)
	ranking := Ranking{
		"1347": -99.8,
		"1348": -99.8,
		"1349": -50,
		"1350": -50,
		"1351": -123.4,
		"1352": -123.4,
	}
	course := NewCourse("1234", jointCourse, ranking)

	programs := Programs()
	genders := Genders()
	ss := []*Student{
		NewStudent("1347", genders.Female, programs.Formal, false),
		NewStudent("1348", genders.Female, programs.Formal, false),
		NewStudent("1349", genders.Female, programs.Formal, false),
		NewStudent("1350", genders.Female, programs.Formal, false),
		NewStudent("1351", genders.Male, programs.Vocat, false),
		NewStudent("1352", genders.Male, programs.Vocat, false),
	}

	for _, s := range ss {
		course.Apply(s)
	}

	if pq := jointCourse.Students(); pq.Len() != 6 {
		t.Error("Wrong number of students", pq.Students())
	}
}

func TestApply_VocatExceeds_RejectReplicas(t *testing.T) {
	strategy := NewApplyStrategy("", 0)
	availableSpots := NewAvailableSpots(6, 0, 0, 0, 0, 2, 0)
	jointCourse := NewJointCourse("1234", availableSpots, strategy)
	ranking := Ranking{
		"1347": -50,
		"1348": -50,
		"1349": -99.8,
		"1350": -99.8,
		"1351": -123.4,
		"1352": -123.4,
	}
	course := NewCourse("1234", jointCourse, ranking)

	programs := Programs()
	genders := Genders()
	ss := []*Student{
		NewStudent("1347", genders.Female, programs.Vocat, false),
		NewStudent("1348", genders.Female, programs.Vocat, false),
		NewStudent("1349", genders.Female, programs.Vocat, false),
		NewStudent("1350", genders.Female, programs.Vocat, false),
		NewStudent("1351", genders.Male, programs.Formal, false),
		NewStudent("1352", genders.Male, programs.Formal, false),
	}

	for _, s := range ss {
		course.Apply(s)
	}

	if pq := jointCourse.Students(); pq.Len() != 4 {
		t.Error("Wrong number of students", pq.Students())
	}
}

func TestApply_FormalExceedsAsVocat_RejectReplicas(t *testing.T) {
	strategy := NewApplyStrategy("", 0)
	availableSpots := NewAvailableSpots(6, 0, 0, 2, 0, 0, 0)
	jointCourse := NewJointCourse("1234", availableSpots, strategy)
	ranking := Ranking{
		"1347": -50,
		"1348": -50,
		"1349": -99.8,
		"1350": -99.8,
		"1351": -123.4,
		"1352": -123.4,
	}
	course := NewCourse("1234", jointCourse, ranking)

	programs := Programs()
	genders := Genders()
	ss := []*Student{
		NewStudent("1347", genders.Female, programs.Vocat, true),
		NewStudent("1348", genders.Female, programs.Vocat, true),
		NewStudent("1349", genders.Female, programs.Vocat, true),
		NewStudent("1350", genders.Female, programs.Vocat, true),
		NewStudent("1351", genders.Male, programs.Vocat, false),
		NewStudent("1352", genders.Male, programs.Vocat, false),
	}

	for _, s := range ss {
		course.Apply(s)
	}

	if pq := jointCourse.Students(); pq.Len() != 4 {
		t.Error("Wrong number of students", pq.Students())
	}
}
