package course

import (
	"testing"

	"github.com/kohpai/tcas-3rd-round-resolver/model/student"
)

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

	if course.isFull {
		t.Error("Course is full", course)
	}
}

func TestIsFull_AvailableSpotsIsReachingZero_ReturnsTrue(t *testing.T) {
	jointCourse := NewJointCourse("1234", 1)
	course := NewCourse("1234", jointCourse)

	jointCourse.Apply()

	if !course.isFull {
		t.Error("Course is NOT full", course)
	}
}

func TestIsFull_AvailableSpotsIsAlreadyZero_ReturnsTrue(t *testing.T) {
	jointCourse := NewJointCourse("1234", 0)
	course := NewCourse("1234", jointCourse)

	if !course.isFull {
		t.Error("Course is NOT full", course)
	}
}

func TestApply_CourseIsNotFull_ReturnsTrue(t *testing.T) {
	jointCourse := NewJointCourse("1234", 1)
	course := NewCourse("1234", jointCourse)

	s := student.NewStudent("1349")

	if !course.Apply(s) {
		t.Error("Apply returns false", course)
	}
}

func TestApply_OneSpotLeft_CourseIsFull(t *testing.T) {
	jointCourse := NewJointCourse("1234", 1)
	course := NewCourse("1234", jointCourse)

	s := student.NewStudent("1349")
	course.Apply(s)

	if !course.isFull {
		t.Error("Course is NOT full", course)
	}
}

func TestApply_MoreSpotsLeft_StudentsAreEnrolled(t *testing.T) {
	jointCourse := NewJointCourse("1234", 3)
	course := NewCourse("1234", jointCourse)

	ss := []*student.Student{
		student.NewStudent("1349"),
		student.NewStudent("1350"),
		student.NewStudent("1351"),
	}

	for _, s := range ss {
		course.Apply(s)
	}

	for i, s := range course.students {
		if s != ss[i] {
			t.Errorf("Student is not enrolled, %d", i)
		}
	}
}
