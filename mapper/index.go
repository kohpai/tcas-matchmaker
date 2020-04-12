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
	ApplicationId    string      `csv:id`
	CitizenId        string      `csv:"citizen_id"`
	Gender           uint8       `csv:gender`
	SchoolProgram    uint8       `csv:school_program`
	FormalApplicable uint8       `csv:formal_niets`
	CourseId         string      `csv:round_id`
	Priority         uint8       `csv:priority`
	Ranking          uint16      `csv:ranking`
	Status           AdmitStatus `csv:status`
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
		_admitStatus.Admitted = 9
		_admitStatus.Full = -2
		_admitStatus.Late = -3
	})
	return _admitStatus
}
