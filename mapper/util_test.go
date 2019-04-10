package mapper

import (
	"testing"
)

func TestCreateRankingMap_Always_ReturnRankingMap(t *testing.T) {
	rankings := []Ranking{
		{"1234", "13499", 1},
		{"1234", "13501", 2},
		{"1234", "13502", 3},
		{"1235", "13500", 2},
		{"1236", "13499", 1},
		{"1237", "13500", 1},
		{"1237", "13499", 2},
	}

	rankingMap := createRankingMap(rankings)

	if rankingMap["1234"]["13499"] != 1 {
		t.Error("Rank is incorrect", rankings[0])
	}
	if rankingMap["1234"]["13501"] != 2 {
		t.Error("Rank is incorrect", rankings[1])
	}
	if rankingMap["1234"]["13502"] != 3 {
		t.Error("Rank is incorrect", rankings[2])
	}
	if rankingMap["1235"]["13500"] != 2 {
		t.Error("Rank is incorrect", rankings[3])
	}
	if rankingMap["1236"]["13499"] != 1 {
		t.Error("Rank is incorrect", rankings[4])
	}
	if rankingMap["1237"]["13500"] != 1 {
		t.Error("Rank is incorrect", rankings[5])
	}
	if rankingMap["1237"]["13499"] != 2 {
		t.Error("Rank is incorrect", rankings[6])
	}
}

func TestCreateJointCourseMap_Always_ReturnsJointCourseMap(t *testing.T) {
	courses := []Course{
		{"1234", "", 10, 1, 0},
		{"1235", "", 11, 1, 0},
		{"1236", "123", 12, 1, 0},
		{"1237", "123", 12, 1, 0},
	}

	jointCourseMap := createJointCourseMap(courses)

	if jc := jointCourseMap["1234"]; jc.AvailableSpots() != 10 {
		t.Error("Available spots is incorrect", jc)
	}

	if jc := jointCourseMap["1235"]; jc.AvailableSpots() != 11 {
		t.Error("Available spots is incorrect", jc)
	}

	if jc := jointCourseMap["123"]; jc.AvailableSpots() != 12 {
		t.Error("Available spots is incorrect", jc)
	}
}

func TestCreateCourseMap_Always_ReturnsCourseMap(t *testing.T) {
	rankings := []Ranking{
		{"1234", "13499", 1},
		{"1234", "13501", 2},
		{"1234", "13502", 3},
		{"1235", "13500", 2},
		{"1236", "13499", 1},
		{"1237", "13500", 1},
		{"1237", "13499", 2},
	}
	courses := []Course{
		{"1234", "", 10, 1, 0},
		{"1235", "", 11, 1, 0},
		{"1236", "123", 12, 1, 0},
		{"1237", "123", 12, 1, 0},
	}

	courseMap := CreateCourseMap(courses, rankings)

	if course := courseMap["1234"]; course.JointCourse().AvailableSpots() != 10 {
		t.Error("Available spots is incorrect", course)
	}

	if course := courseMap["1235"]; course.JointCourse().AvailableSpots() != 11 {
		t.Error("Available spots is incorrect", course)
	}

	if course := courseMap["1236"]; course.JointCourse().AvailableSpots() != 12 {
		t.Error("Available spots is incorrect", course)
	}

	if course := courseMap["1237"]; course.JointCourse().AvailableSpots() != 12 {
		t.Error("Available spots is incorrect", course)
	}

	jc1, jc2 := courseMap["1236"].JointCourse(), courseMap["1237"].JointCourse()

	if jc1 != jc2 {
		t.Error("Incorrect joint course", jc1, jc2)
	}

	if len(jc1.Courses()) != 2 {
		t.Error("Length of courses is incorrect", jc1)
	}
}

func TestCreateStudentMap_Always_ReturnsStudentMap(t *testing.T) {
	rankings := []Ranking{
		{"1234", "13499", 1},
		{"1234", "13501", 2},
		{"1234", "13502", 3},
		{"1235", "13500", 2},
		{"1236", "13499", 1},
		{"1237", "13500", 1},
		{"1237", "13499", 2},
	}
	courses := []Course{
		{"1234", "", 10, 1, 0},
		{"1235", "", 11, 1, 0},
		{"1236", "123", 12, 1, 0},
		{"1237", "123", 12, 1, 0},
	}
	students := []Student{
		{"13499", "1234", 1},
		{"13499", "1236", 2},
		{"13500", "1237", 3},
	}

	courseMap := CreateCourseMap(courses, rankings)
	studentMap := CreateStudentMap(students, courseMap)

	student1, student2 := studentMap["13499"], studentMap["13500"]
	if course, err := student1.PreferredCourse(1); err != nil || course.Id() != "1234" {
		t.Error("Course is incorrect", course)
	}

	if course, err := student1.PreferredCourse(2); err != nil || course.Id() != "1236" {
		t.Error("Course is incorrect", course)
	}

	if course, err := student2.PreferredCourse(3); err != nil || course.Id() != "1237" {
		t.Error("Course is incorrect", course)
	}
}
