package model

import "testing"

func TestNewStudent_Always_ReturnsStudent(t *testing.T) {
	student := NewStudent("1349900696510")

	if student.citizenId != "1349900696510" {
		t.Error("Citizen ID not matched", student)
	}

	if student.applicationStatus != ApplicationStatuses().Pending {
		t.Error("Application status is not PENDING", student)
	}

	isEmpty := true
	for _, course := range student.preferredCourses {
		if course != nil {
			isEmpty = false
			break
		}
	}

	if !isEmpty {
		t.Error("Preferred Courses is not empty", student)
	}

	if student.course != nil {
		t.Error("Course is not nil", student)
	}
}

func TestSetPreferredCourse_PriorityWithinOneToSix_ReturnsNil(t *testing.T) {
	jointCourse := NewJointCourse("1234", 1)
	course := NewCourse("1234", jointCourse)
	student := NewStudent("1349900696510")

	if err := student.SetPreferredCourse(2, course); err != nil {
		t.Error("Cannot set preferred course", err)
	}

	if student.preferredCourses[1] != course {
		t.Error("Course does not matched", student)
	}
}

func TestSetPreferredCourse_PriorityOutOfRange_ReturnsError(t *testing.T) {
	jointCourse := NewJointCourse("1234", 1)
	course := NewCourse("1234", jointCourse)
	student := NewStudent("1349900696510")

	if err := student.SetPreferredCourse(7, course); err == nil {
		t.Error("Set preferred course without error")
	}

	isEmpty := true
	for _, course := range student.preferredCourses {
		if course != nil {
			isEmpty = false
			break
		}
	}

	if !isEmpty {
		t.Error("Preferred Courses is not empty", student)
	}
}
