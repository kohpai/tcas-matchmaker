package model

type AllowAllStrategy struct {
	BaseStrategy
}

func (strategy *AllowAllStrategy) Apply(rankedStudent *RankedStudent) bool {
	rank := rankedStudent.rank
	rankCount := strategy.rankCount
	rankCount[rank] += 1
	if rankCount[rank] > 1 {
		return true
	}

	return strategy.BaseStrategy.Apply(rankedStudent)
}
