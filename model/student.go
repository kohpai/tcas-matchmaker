package model

import (
	"errors"
	"fmt"
)

type Student struct {
	citizenId         string
	gender            Gender
	applicationStatus ApplicationStatus
	apps              [6]*Application
	preferredCourses  [6]*Course
	appIds            [6]string
	appIndex          int
}

func NewStudent(citizenId string, gender Gender) *Student {
	return &Student{
		citizenId:         citizenId,
		gender:            gender,
		applicationStatus: ApplicationStatuses().Pending,
		appIndex:          -1,
	}
}

func (student *Student) SetCourse(course *Course) {
	for i, a := range student.apps {
		if a != nil && a.Course() == course {
			student.appIndex = i
			break
		}
	}
	student.applicationStatus = ApplicationStatuses().Accepted
}

func (student *Student) ClearCourse() {
	student.appIndex = -1
	student.applicationStatus = ApplicationStatuses().Pending
}

func (student *Student) CitizenId() string {
	return student.citizenId
}

func (student *Student) ApplicationStatus() ApplicationStatus {
	return student.applicationStatus
}

func (student *Student) Application(priority uint8) (*Application, error) {
	if priority < 1 || 6 < priority {
		return nil, errors.New("priority out of range")
	}

	return student.apps[priority-1], nil
}

func (student *Student) AppIndex() int {
	return student.appIndex
}

func (student *Student) Gender() Gender {
	return student.gender
}

func (student *Student) SetPreferredApp(priority uint8, course *Course, appId string) error {
	if priority < 1 || 6 < priority {
		return errors.New("priority out of range")
	}

	app := &Application{
		course: course,
		id:     appId,
	}

	student.apps[priority-1] = app
	return nil
}

func (student *Student) Propose() ApplicationStatus {
	statuses := ApplicationStatuses()
	if student.applicationStatus != statuses.Pending {
		return student.applicationStatus
	}

	isAccepted := false
	for _, app := range student.apps {
		if app == nil {
			continue
		}

		if isAccepted = app.Course().Apply(student); isAccepted {
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
