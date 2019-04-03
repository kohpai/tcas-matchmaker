package course

import "testing"

func TestNewJointCourse_Always_ReturnsJointCourse(t *testing.T) {
	jointCourse := NewJointCourse("1234", 100)

	if jointCourse.id != "1234" {
		t.Error("Joint course ID not matched", jointCourse)
	}

	if jointCourse.availableSpots != 100 {
		t.Error("Joint course ID not matched", jointCourse)
	}

	if len(jointCourse.courses) != 0 {
		t.Error("Courses is not empty", jointCourse)
	}
}
