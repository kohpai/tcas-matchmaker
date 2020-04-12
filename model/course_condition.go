package model

import "sync"

type Condition string

type condition struct {
	once      sync.Once
	AllowAll  Condition
	DenyAll   Condition
	AllowSome Condition
}

var _condition condition

// TransactionTypes returns the types of a transaction
func Conditions() condition {
	_condition.once.Do(func() {
		_condition.AllowAll = "A"
		_condition.DenyAll = "B"
		_condition.AllowSome = "C"
	})
	return _condition
}
