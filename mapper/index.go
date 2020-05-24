package mapper

import (
	"sync"
)

type AdmitStatus int8

type admitStatus struct {
	Admitted AdmitStatus
	Full     AdmitStatus
	Late     AdmitStatus
}

type Application struct {
	ApplicationId    string      `csv:"_id"`
	CitizenId        string      `csv:"register_id"`
	Gender           uint8       `csv:"title"`
	SchoolProgram    uint8       `csv:"school_program"`
	FormalApplicable uint8       `csv:"school_niets_formal"`
	CourseId         string      `csv:"round_id"`
	Priority         uint8       `csv:"priority"`
	Ranking          float64     `csv:"score"`
	Status           AdmitStatus `csv:"status"`
}

type Course struct {
	CourseId            string `csv:"_id"`
	UniversityId        string `csv:"university_id"`
	JointId             string `csv:"join_id"`
	ReceiveAmount       uint16 `csv:"receive_student_number"`
	ExceedAllowedAmount string `csv:"receive_add_limit"`
	// Gender based
	MaleReceiveAmount   uint16 `csv:"gender_male_number"`
	FemaleReceiveAmount uint16 `csv:"gender_female_number"`
	// Program based
	FormalReceiveAmount    uint16 `csv:"receive_student_number_formal"`
	InterReceiveAmount     uint16 `csv:"receive_student_number_international"`
	VocatReceiveAmount     uint16 `csv:"receive_student_number_vocational"`
	NonformalReceiveAmount uint16 `csv:"receive_student_number_nonformal"`
}

var once sync.Once
var _admitStatus admitStatus

// TransactionTypes returns the types of a transaction
func AdmitStatuses() admitStatus {
	once.Do(func() {
		_admitStatus.Admitted = 2
		_admitStatus.Full = 8 // rejected by the course
		_admitStatus.Late = 9 // already admitted to some other course
	})
	return _admitStatus
}
