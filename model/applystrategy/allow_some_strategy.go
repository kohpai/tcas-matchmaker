package applystrategy

import (
	"container/heap"

	"github.com/kohpai/tcas-3rd-round-resolver/model/common"
	"github.com/kohpai/tcas-3rd-round-resolver/model/course"
)

type AllowSomeStrategy struct {
	BaseStrategy
	leastReplicatedRank int
	rankCount           course.RankCount
	exceedLimit         int
}

func (strategy *AllowSomeStrategy) countBeingRemovedReplicas() (int, int) {
	jc := strategy.jointCourse
	pq := jc.Students()
	length, limit := pq.Len(), jc.Limit()
	delta := length - limit
	count := strategy.countEdgeReplicas()

	switch {
	case delta > strategy.exceedLimit:
		return count, count - delta
	case count <= delta:
		return count, 0
	}

	return 0, 0
}

func (strategy *AllowSomeStrategy) Apply(rankedStudent common.RankedStudent) bool {
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
	count, inc := strategy.countBeingRemovedReplicas()

	if count > 0 {
		strategy.rankCount[lastRank] = count
		strategy.leastReplicatedRank = lastRank
	}

	for i := 0; i < count; i++ {
		rs := heap.Pop(pq).(common.RankedStudent)
		rs.Student().ClearCourse()
	}
	for i := 0; i < inc; i++ {
		jc.IncSpots()
	}

	if rank < lastRank {
		return true
	}
	if count > 0 {
		return false
	}
	return true
}
