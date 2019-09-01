package applystrategy

import (
	"container/heap"

	"github.com/kohpai/tcas-3rd-round-resolver/model/common"
	"github.com/kohpai/tcas-3rd-round-resolver/model/course"
)

type DenyAllStrategy struct {
	BaseStrategy
	leastReplicatedRank int
	rankCount           course.RankCount
}

func (strategy *DenyAllStrategy) countBeingRemovedReplicas() int {
	return strategy.countEdgeReplicas()
}

func (strategy *DenyAllStrategy) Apply(rankedStudent common.RankedStudent) bool {
	jc := strategy.jointCourse
	pq := jc.Students()
	rank := rankedStudent.Rank()

	if !jc.IsFull() {
		lrr := strategy.leastReplicatedRank
		if lrr == 0 || rank < lrr {
			heap.Push(pq, rankedStudent)
			jc.DecSpots()
			return true
		}

		if rank == lrr {
			strategy.rankCount[rank] += 1
		}
		return false
	}

	tmp := heap.Pop(pq).(common.RankedStudent)
	heap.Push(pq, tmp)
	lastRank := tmp.Rank()
	if rank > lastRank {
		return false
	}

	heap.Push(pq, rankedStudent)
	count := strategy.countBeingRemovedReplicas()
	strategy.rankCount[lastRank], strategy.leastReplicatedRank = count, lastRank
	for ; count > 0; count-- {
		rs := heap.Pop(pq).(common.RankedStudent)
		rs.Student().ClearCourse()
		jc.IncSpots()
	}

	jc.DecSpots()

	return rank < lastRank
}
