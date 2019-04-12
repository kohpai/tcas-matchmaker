package model

import (
	"fmt"
)

type ClearingHouse struct {
	students         []*Student
	acceptedStudents []*Student
	rejectedStudents []*Student
}

func NewClearingHouse(students []*Student) *ClearingHouse {
	as, rs := make([]*Student, 0), make([]*Student, 0)
	return &ClearingHouse{
		students, as, rs,
	}
}

func (ch *ClearingHouse) Students() []*Student {
	return ch.students
}

func (ch *ClearingHouse) AcceptedStudents() []*Student {
	return ch.acceptedStudents
}

func (ch *ClearingHouse) RejectedStudents() []*Student {
	return ch.rejectedStudents
}

func (ch *ClearingHouse) Execute() {
	ch.executePending()

	statuses := ApplicationStatuses()
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
	statuses := ApplicationStatuses()
	isPending := true
	for i := 0; isPending; i++ {
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
