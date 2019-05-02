package model

import (
	"fmt"

	st "github.com/kohpai/tcas-3rd-round-resolver/model/student"
)

type ClearingHouse struct {
	students         []*st.Student
	acceptedStudents []*st.Student
	rejectedStudents []*st.Student
}

func NewClearingHouse(students []*st.Student) *ClearingHouse {
	as, rs := make([]*st.Student, 0), make([]*st.Student, 0)
	return &ClearingHouse{
		students, as, rs,
	}
}

func (ch *ClearingHouse) Students() []*st.Student {
	return ch.students
}

func (ch *ClearingHouse) AcceptedStudents() []*st.Student {
	return ch.acceptedStudents
}

func (ch *ClearingHouse) RejectedStudents() []*st.Student {
	return ch.rejectedStudents
}

func (ch *ClearingHouse) Execute() {
	ch.executePending()

	statuses := st.ApplicationStatuses()
	for _, student := range ch.students {
		switch student.ApplicationStatus() {
		case statuses.Accepted:
			ch.acceptedStudents = append(ch.acceptedStudents, student)
		case statuses.Rejected:
			ch.rejectedStudents = append(ch.rejectedStudents, student)
		}
	}
}

func (ch *ClearingHouse) executePending() {
	statuses := st.ApplicationStatuses()
	isPending := true
	for isPending {
		isPending = false
		for _, student := range ch.students {
			if student.ApplicationStatus() == statuses.Pending {
				isPending = true
				student.Propose()
			}
		}
	}
}

func (ch *ClearingHouse) String() string {
	return fmt.Sprintf(
		"{students: %v, acceptedStudents: %v, rejectedStudents: %v}",
		ch.students,
		ch.acceptedStudents,
		ch.rejectedStudents,
	)
}
