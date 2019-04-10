package mapper

import (
	"testing"
)

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
	courses := []Course{
		{"1234", "", 10, 1, 0},
		{"1235", "", 11, 1, 0},
		{"1236", "123", 12, 1, 0},
		{"1237", "123", 12, 1, 0},
	}

	courseMap := CreateCourseMap(courses)

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

	courseMap := CreateCourseMap(courses)
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
