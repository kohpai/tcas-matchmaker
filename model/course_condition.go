package model

import "sync"

type Condition int

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
		_condition.AllowAll = 1
		_condition.DenyAll = 2
		_condition.AllowSome = 3
	})
	return _condition
}
