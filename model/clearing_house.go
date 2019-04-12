package model

import (
	"fmt"
)

type ClearingHouse struct {
	pendingStudents  []*Student
	acceptedStudents []*Student
	rejectedStudents []*Student
}

func NewClearingHouse(students []*Student) *ClearingHouse {
	as, rs := make([]*Student, 0), make([]*Student, 0)
	return &ClearingHouse{
		students, as, rs,
	}
}

func (ch *ClearingHouse) PendingStudents() []*Student {
	return ch.pendingStudents
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
	for _, student := range ch.pendingStudents {
		switch student.ApplicationStatus() {
		case statuses.Accepted:
			ch.acceptedStudents = append(ch.acceptedStudents, student)
		case statuses.Rejected:
			ch.rejectedStudents = append(ch.rejectedStudents, student)
		}
	}

	ch.pendingStudents = []*Student{}
}

func (ch *ClearingHouse) executePending() {
	statuses := ApplicationStatuses()
	isPending := false
	for _, student := range ch.pendingStudents {
		if student.ApplicationStatus() == statuses.Pending {
			isPending = true
			student.Propose()
		}
	}

	if isPending {
		ch.executePending()
	}
}

func (ch *ClearingHouse) String() string {
	return fmt.Sprintf(
		"{pendingStudents: %v, acceptedStudents: %v, rejectedStudents: %v}",
		ch.pendingStudents,
		ch.acceptedStudents,
		ch.rejectedStudents,
	)
}
