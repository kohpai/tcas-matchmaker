package model

import (
	"container/heap"
	"log"
)

type ApplyStrategy interface {
	SetCourse(*Course)
	IncRankCount(uint16)
	InitRanking(Ranking) Ranking
	Apply(*RankedStudent) bool
}

type BaseStrategy struct {
	course    *Course
	rankCount RankCount
}

func NewApplyStrategy(condition Condition) ApplyStrategy {
	base := BaseStrategy{
		nil,
		make(RankCount),
	}

	conditions := Conditions()

	switch condition {
	case conditions.AllowAll:
		return &AllowAllStrategy{
			base,
		}
	case conditions.DenyAll:
		return &DenyAllStrategy{
			base,
		}
	}

	return &base
}

func (strategy *BaseStrategy) SetCourse(course *Course) {
	strategy.course = course
}

func (strategy *BaseStrategy) IncRankCount(rank uint16) {
	strategy.rankCount[rank] += 1
}

func (strategy *BaseStrategy) InitRanking(ranking Ranking) Ranking {
	return ranking
}

func (strategy *BaseStrategy) Apply(rankedStudent *RankedStudent) bool {
	course := strategy.course
	if course.IsFull() {
		rs := heap.Pop(course.Students()).(*RankedStudent)
		rs.Student().ClearCourse()
		return rankedStudent != rs
	}

	// @ASSERTION, this shouldn't happen
	if !course.JointCourse().Apply() {
		log.Println("Apply returns false")
	}

	return true
}
