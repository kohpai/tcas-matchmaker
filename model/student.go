package model

import (
	"errors"
	"fmt"
)

type Student struct {
	citizenId         string
	applicationStatus ApplicationStatus
	preferredCourses  [6]*Course
	appIds            [6]string
	courseIndex       int
}

func NewStudent(citizenId string) *Student {
	return &Student{
		citizenId:         citizenId,
		applicationStatus: ApplicationStatuses().Pending,
		courseIndex:       -1,
	}
}

func (student *Student) SetCourse(course *Course) {
	for i, c := range student.preferredCourses {
		if c == course {
			student.courseIndex = i
			break
		}
	}
	student.applicationStatus = ApplicationStatuses().Accepted
}

func (student *Student) ClearCourse() {
	student.courseIndex = -1
	student.applicationStatus = ApplicationStatuses().Pending
}

func (student *Student) CitizenId() string {
	return student.citizenId
}

func (student *Student) ApplicationStatus() ApplicationStatus {
	return student.applicationStatus
}

func (student *Student) PreferredCourse(priority uint8) (*Course, error) {
	if priority < 1 || 6 < priority {
		return nil, errors.New("priority out of range")
	}

	return student.preferredCourses[priority-1], nil
}

func (student *Student) AppId(priority uint8) (string, error) {
	if priority < 1 || 6 < priority {
		return "", errors.New("priority out of range")
	}

	return student.appIds[priority-1], nil
}

func (student *Student) CourseIndex() int {
	return student.courseIndex
}

func (student *Student) SetPreferredCourse(priority uint8, course *Course, appId string) error {
	if priority < 1 || 6 < priority {
		return errors.New("priority out of range")
	}

	student.preferredCourses[priority-1] = course
	student.appIds[priority-1] = appId
	return nil
}

func (student *Student) Propose() ApplicationStatus {
	statuses := ApplicationStatuses()
	if student.applicationStatus != statuses.Pending {
		return student.applicationStatus
	}

	isAccepted := false
	for _, course := range student.preferredCourses {
		if course == nil {
			continue
		}

		if isAccepted = course.Apply(student); isAccepted {
			break
		}
	}

	if !isAccepted {
		student.applicationStatus = statuses.Rejected
	}

	return student.applicationStatus
}

func (student *Student) String() string {
	return fmt.Sprintf(
		"{citizenId: %s, applicationStatus: %s}",
		student.citizenId,
		student.applicationStatus,
	)
}