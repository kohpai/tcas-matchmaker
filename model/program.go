package model

import (
	"sync"
)

type Program uint8

type program struct {
	once      sync.Once
	Formal    Program
	Inter     Program
	Vocat     Program
	NonFormal Program
}

var _program program

func Programs() program {
	_program.once.Do(func() {
		_program.Formal = 1
		_program.Inter = 2
		_program.Vocat = 3
		_program.NonFormal = 4
	})
	return _program
}
