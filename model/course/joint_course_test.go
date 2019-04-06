package course

import "testing"

func TestNewJointCourse_Always_ReturnsJointCourse(t *testing.T) {
	jointCourse := NewJointCourse("1234", 100)

	if jointCourse.id != "1234" {
		t.Error("Joint course ID not matched", jointCourse)
	}

	if jointCourse.availableSpots != 100 {
		t.Error("Joint course available spots is incorrect", jointCourse)
	}

	if len(jointCourse.courses) != 0 {
		t.Error("Courses is not empty", jointCourse)
	}
}

func TestApply_AvailableSpotsGreaterThanZero_DecreasesByOne(t *testing.T) {
	jointCourse := NewJointCourse("1234", 100)

	if !jointCourse.Apply() {
		t.Error("Apply returns false", jointCourse)
	}

	if jointCourse.availableSpots != 99 {
		t.Error("Joint course available spots is incorrect", jointCourse)
	}
}

func TestApply_AvailableSpotsIsZero_ReturnsFalse(t *testing.T) {
	jointCourse := NewJointCourse("1234", 0)

	if jointCourse.Apply() {
		t.Error("Apply returns true", jointCourse)
	}

	if jointCourse.availableSpots != 0 {
		t.Error("Joint course available spots is incorrect", jointCourse)
	}
}
