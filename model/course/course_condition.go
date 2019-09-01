package model

import "sync"

type Condition int

type condition struct {
	once      sync.Once
	allowAll  Condition
	denyAll   Condition
	allowSome Condition
}

var _condition condition

// TransactionTypes returns the types of a transaction
func Conditions() *condition {
	_condition.once.Do(func() {
		_condition.allowAll = 1
		_condition.denyAll = 2
		_condition.allowSome = 3
	})
	return &_condition
}

func (con *condition) AllowAll() Condition {
	return con.allowAll
}

func (con *condition) DenyAll() Condition {
	return con.denyAll
}

func (con *condition) AllowSome() Condition {
	return con.allowSome
}
