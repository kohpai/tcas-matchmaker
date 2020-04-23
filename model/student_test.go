package model

import "testing"

func TestNewStudent_Always_ReturnsStudent(t *testing.T) {
	student := NewStudent("1349900696510", Genders().Male, Programs().Formal, false)

	if student.CitizenId() != "1349900696510" {
		t.Error("Citizen ID not matched", student)
	}

	if student.ApplicationStatus() != ApplicationStatuses().Pending {
		t.Error("Application status is not PENDING", student)
	}

	for i := 1; i < 7; i++ {
		app, err := student.Application(uint8(i))
		if app != nil || err != nil {
			t.Error("Preferred application is not empty", student, err)
			break
		}
	}

	if student.AppIndex() != -1 {
		t.Error("app index is not -1", student)
	}
}

func TestSetPreferredCourse_PriorityWithinOneToSix_ReturnsNil(t *testing.T) {
	strategy := NewApplyStrategy(Conditions().AllowAll, 0)
	jointCourse := NewJointCourse("1234", NewAvailableSpots(1, 0, 0, 0, 0, 0, 0), strategy)
	course := NewCourse("1234", jointCourse, nil)
	student := NewStudent("1349900696510", Genders().Male, Programs().Formal, false)

	if err := student.SetPreferredApp(2, course, ""); err != nil {
		t.Error("Cannot set preferred app", err)
	}

	if app, err := student.Application(2); app.Course() != course || err != nil {
		t.Error("Course does not matched", student, err)
	}
}

func TestSetPreferredCourse_PriorityOutOfRange_ReturnsError(t *testing.T) {
	strategy := NewApplyStrategy(Conditions().AllowAll, 0)
	jointCourse := NewJointCourse("1234", NewAvailableSpots(1, 0, 0, 0, 0, 0, 0), strategy)
	course := NewCourse("1234", jointCourse, nil)
	student := NewStudent("1349900696510", Genders().Male, Programs().Formal, false)

	if err := student.SetPreferredApp(7, course, ""); err == nil {
		t.Error("Set preferred app without error")
	}

	for i := 1; i < 7; i++ {
		app, err := student.Application(uint8(i))
		if app != nil || err != nil {
			t.Error("Application is not empty", student, err)
			break
		}
	}
}
