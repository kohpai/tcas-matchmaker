package model

import (
	"sync"
)

type Program string

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
		_program.Formal = "FORMAL"
		_program.Inter = "INTER"
		_program.Vocat = "VOCAT"
		_program.NonFormal = "NONFORMAL"
	})
	return _program
}
