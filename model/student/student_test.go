package student

import "testing"

func TestNewStudent_Always_ReturnsStudent(t *testing.T) {
	student := NewStudent("1349900696510")

	if student.citizenId != "1349900696510" {
		t.Error("Citizen ID not matched", student)
	}

	if student.applicationStatus != ApplicationStatuses().Pending {
		t.Error("Application status is not PENDING", student)
	}
}
