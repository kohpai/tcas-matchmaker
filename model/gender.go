package model

import (
	"sync"
)

type Gender string

type gender struct {
	once   sync.Once
	Male   Gender
	Female Gender
}

var _gender gender

func Genders() gender {
	_gender.once.Do(func() {
		_gender.Male = "MALE"
		_gender.Female = "FEMALE"
	})
	return _gender
}
