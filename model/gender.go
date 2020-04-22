package model

import (
	"sync"
)

type Gender uint8

type gender struct {
	once   sync.Once
	Male   Gender
	Female Gender
}

var _gender gender

func Genders() gender {
	_gender.once.Do(func() {
		_gender.Male = 1
		_gender.Female = 2
	})
	return _gender
}
