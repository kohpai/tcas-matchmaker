package student

import "fmt"

type Student struct {
	citizenId         string
	applicationStatus ApplicationStatus
}

func (student *Student) SetStatus(status ApplicationStatus) {
	student.applicationStatus = status
}

func (student *Student) Propose() {
	// for each course in perferredCourses
}

func NewStudent(citizenId string) *Student {
	return &Student{
		citizenId:         citizenId,
		applicationStatus: ApplicationStatuses().Pending,
	}
}

func (student *Student) String() string {
	return fmt.Sprintf(
		"{\n\tCitizen ID: %s,\n\t Application Status: %s\n}",
		student.citizenId,
		student.applicationStatus,
	)
}
