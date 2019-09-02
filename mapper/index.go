package mapper

import (
	"sync"

	"github.com/kohpai/tcas-3rd-round-resolver/model/course"
)

type Course struct {
	Id        string           `json:"course_id"`
	JointId   string           `json:"round_3_join_id"`
	Limit     int              `json:"round_3_receive"`
	Condition course.Condition `json:"round_3_condition"`
	AddLimit  int              `json:"round_3_add_limit"`
}

type Student struct {
	CitizenId string `json:"citizen_id"`
	CourseId  string `json:"course_id"`
	Priority  int    `json:"priority"`
}

type Ranking struct {
	CourseId  string `csv:"course_id"`
	CitizenId string `csv:"citizen_id"`
	Rank      int    `csv:"ranking"`
}

type AdmitStatus int8

type admitStatus struct {
	Admitted AdmitStatus
	Full     AdmitStatus
	Late     AdmitStatus
}

type Output struct {
	UniversityId      string      `csv:"university_id"`
	UniversityName    string      `csv:"university_name"`
	CourseId          string      `csv:"course_id"`
	FacultyName       string      `csv:"faculty_name"`
	CourseName        string      `csv:"course_name"`
	ProjectName       string      `csv:"project_name"`
	ApplicationId     string      `csv:"application_id"`
	CitizenId         string      `csv:"citizen_id"`
	Title             string      `csv:"title"`
	FirstName         string      `csv:"first_name"`
	LastName          string      `csv:"last_name"`
	PhoneNumber       string      `csv:"phone_number"`
	Email             string      `csv:"email"`
	ApplicationDate   string      `csv:"application_date"`
	InterviewLocation string      `csv:"interview_location"`
	InterviewDate     string      `csv:"interview_date"`
	InterviewTime     string      `csv:"interview_time"`
	Ranking           int         `csv:"ranking"`
	Round             string      `csv:"round"`
	AdmitStatus       AdmitStatus `csv:"admit_status"`
}

var once sync.Once
var _admitStatus admitStatus

// TransactionTypes returns the types of a transaction
func AdmitStatuses() admitStatus {
	once.Do(func() {
		_admitStatus.Admitted = 9
		_admitStatus.Full = -2
		_admitStatus.Late = -3
	})
	return _admitStatus
}
