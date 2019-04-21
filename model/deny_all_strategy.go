package model

type DenyAllStrategy struct {
	BaseStrategy
}

func (strategy *DenyAllStrategy) InitRanking(ranking Ranking) Ranking {
	rankCount := make(RankCount)
	for _, rank := range ranking {
		rankCount[rank] += 1
	}

	for key, rank := range ranking {
		if rankCount[rank] > 1 {
			delete(ranking, key)
		}
	}

	return ranking
}
