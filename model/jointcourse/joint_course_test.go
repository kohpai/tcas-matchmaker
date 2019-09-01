package model

import "testing"

func TestNewJointCourse_Always_ReturnsJointCourse(t *testing.T) {
	strategy := NewApplyStrategy(Conditions().AllowAll(), 0)
	jointCourse := NewJointCourse("1234", 100, strategy)

	if jointCourse.Id() != "1234" {
		t.Error("Joint course ID not matched", jointCourse)
	}

	if jointCourse.AvailableSpots() != 100 {
		t.Error("Joint course available spots is incorrect", jointCourse)
	}

	if len(jointCourse.Courses()) != 0 {
		t.Error("Courses is not empty", jointCourse)
	}
}

func TestDecSpots_AvailableSpotsGreaterThanZero_DecreasesByOne(t *testing.T) {
	strategy := NewApplyStrategy(Conditions().AllowAll(), 0)
	jointCourse := NewJointCourse("1234", 100, strategy)

	if jointCourse.DecSpots(); jointCourse.AvailableSpots() != 99 {
		t.Error("Joint course available spots is incorrect", jointCourse)
	}
}

func TestDecSpots_AvailableSpotsIsZero_ReturnsFalse(t *testing.T) {
	strategy := NewApplyStrategy(Conditions().AllowAll(), 0)
	jointCourse := NewJointCourse("1234", 0, strategy)

	if jointCourse.DecSpots(); jointCourse.AvailableSpots() != 0 {
		t.Error("Joint course available spots is incorrect", jointCourse)
	}
}

func TestRegisterCourse_ByDefault_RegistersCourse(t *testing.T) {
	strategy := NewApplyStrategy(Conditions().AllowAll(), 0)
	jointCourse := NewJointCourse("1234", 10, strategy)
	courses := []*Course{
		NewCourse("1234", jointCourse, nil),
		NewCourse("1235", jointCourse, nil),
		NewCourse("1236", jointCourse, nil),
	}

	regCourses := jointCourse.Courses()
	if len(regCourses) != 3 {
		t.Error("Courses length is incorrect", jointCourse)
	}

	for i := 0; i < 3; i++ {
		course := regCourses[i]
		if course != courses[i] {
			t.Errorf("Course %d is incorrect", i)
		}
	}
}
