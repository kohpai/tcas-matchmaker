package model

import (
	"errors"
	"fmt"
)

type Student struct {
	citizenId         string
	applicationStatus ApplicationStatus
	preferredCourses  [6]*Course
	course            *Course
}

func (student *Student) SetStatus(status ApplicationStatus) {
	student.applicationStatus = status
}

func NewStudent(citizenId string) *Student {
	return &Student{
		citizenId:         citizenId,
		applicationStatus: ApplicationStatuses().Pending,
	}
}

func (student *Student) SetPreferredCourse(priority uint8, course *Course) error {
	if priority < 1 || 6 < priority {
		return errors.New("priority out of range")
	}

	student.preferredCourses[priority-1] = course
	return nil
}

func (student *Student) String() string {
	return fmt.Sprintf(
		"{\n\tCitizen ID: %s,\n\t Application Status: %s\n}",
		student.citizenId,
		student.applicationStatus,
	)
}
