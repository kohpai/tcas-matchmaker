package student

import (
	"sync"
)

type ApplicationStatus string

type applicationStatus struct {
	once     sync.Once
	accepted ApplicationStatus
	rejected ApplicationStatus
	pending  ApplicationStatus
}

var _applicationStatus applicationStatus

// TransactionTypes returns the types of a transaction
func ApplicationStatuses() *applicationStatus {
	_applicationStatus.once.Do(func() {
		_applicationStatus.accepted = "ACCEPTED"
		_applicationStatus.rejected = "REJECTED"
		_applicationStatus.pending = "PENDING"
	})
	return &_applicationStatus
}

func (appStatus *applicationStatus) Accepted() ApplicationStatus {
	return appStatus.accepted
}

func (appStatus *applicationStatus) Rejected() ApplicationStatus {
	return appStatus.rejected
}

func (appStatus *applicationStatus) Pending() ApplicationStatus {
	return appStatus.pending
}
