package model

import "testing"

func TestNewStudent_Always_ReturnsStudent(t *testing.T) {
	student := NewStudent("1349900696510")

	if student.CitizenId() != "1349900696510" {
		t.Error("Citizen ID not matched", student)
	}

	if student.ApplicationStatus() != ApplicationStatuses().Pending {
		t.Error("Application status is not PENDING", student)
	}

	for i := 1; i < 7; i++ {
		course, err := student.PreferredCourse(uint8(i))
		if course != nil || err != nil {
			t.Error("Preferred Courses is not empty", student, err)
			break
		}
	}

	if student.Course() != nil {
		t.Error("Course is not nil", student)
	}
}

func TestSetPreferredCourse_PriorityWithinOneToSix_ReturnsNil(t *testing.T) {
	jointCourse := NewJointCourse("1234", 1)
	course := NewCourse("1234", jointCourse, nil)
	student := NewStudent("1349900696510")

	if err := student.SetPreferredCourse(2, course); err != nil {
		t.Error("Cannot set preferred course", err)
	}

	if c, err := student.PreferredCourse(2); c != course || err != nil {
		t.Error("Course does not matched", student, err)
	}
}

func TestSetPreferredCourse_PriorityOutOfRange_ReturnsError(t *testing.T) {
	jointCourse := NewJointCourse("1234", 1)
	course := NewCourse("1234", jointCourse, nil)
	student := NewStudent("1349900696510")

	if err := student.SetPreferredCourse(7, course); err == nil {
		t.Error("Set preferred course without error")
	}

	for i := 1; i < 7; i++ {
		course, err := student.PreferredCourse(uint8(i))
		if course != nil || err != nil {
			t.Error("Preferred Courses is not empty", student, err)
			break
		}
	}
}
