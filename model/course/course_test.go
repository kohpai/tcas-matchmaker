package course

import "testing"

func TestNewCourse_Always_ReturnsCourse(t *testing.T) {
	jointCourse := NewJointCourse("1234", 10)
	course := NewCourse("1234", jointCourse)

	if course.id != "1234" {
		t.Error("Course ID not matched", course)
	}

	if course.jointCourse != jointCourse {
		t.Error("Joint course is incorrect", course)
	}
}

func TestIsFull_AvailableSpotsGreaterThanZero_ReturnsFalse(t *testing.T) {
	jointCourse := NewJointCourse("1234", 1)
	course := NewCourse("1234", jointCourse)

	if course.IsFull() {
		t.Error("Course is full", course)
	}
}
