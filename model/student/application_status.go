package student

import (
	"sync"
)

type ApplicationStatus string

type applicationStatus struct {
	Accepted ApplicationStatus
	Rejected ApplicationStatus
	Pending  ApplicationStatus
}

var once sync.Once
var _applicationStatus applicationStatus

// TransactionTypes returns the types of a transaction
func ApplicationStatuses() applicationStatus {
	once.Do(func() {
		_applicationStatus.Accepted = "ACCEPTED"
		_applicationStatus.Rejected = "REJECTED"
		_applicationStatus.Pending = "PENDING"
	})
	return _applicationStatus
}
