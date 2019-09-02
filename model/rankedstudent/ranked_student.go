package rankedstudent

import (
	"fmt"

	"github.com/kohpai/tcas-3rd-round-resolver/model/common"
)

type RankedStudent struct {
	student common.Student
	rank    int
	index   int
}

func NewRankedStudent(student common.Student, rank, index int) *RankedStudent {
	return &RankedStudent{
		student,
		rank,
		index,
	}
}

func (rs *RankedStudent) SetIndex(index int) {
	rs.index = index
}

func (rs *RankedStudent) Student() common.Student {
	return rs.student
}

func (rs *RankedStudent) Rank() int {
	return rs.rank
}

func (rs *RankedStudent) String() string {
	return fmt.Sprintf("{student: %v, rank: %v}", rs.student, rs.rank)
}
