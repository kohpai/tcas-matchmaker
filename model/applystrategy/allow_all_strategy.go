package applystrategy

import (
	"container/heap"

	"github.com/kohpai/tcas-3rd-round-resolver/model/common"
)

type AllowAllStrategy struct {
	BaseStrategy
}

func (strategy *AllowAllStrategy) countBeingRemovedReplicas() int {
	jc := strategy.jointCourse
	students := jc.Students().Students()
	length, limit := len(students), jc.Limit()
	count := strategy.countEdgeReplicas()

	if delta := length - limit; count <= delta {
		return count
	}
	return 0
}

func (strategy *AllowAllStrategy) Apply(rankedStudent common.RankedStudent) bool {
	jc := strategy.jointCourse
	pq := jc.Students()

	if !jc.IsFull() {
		heap.Push(pq, rankedStudent)
		jc.DecSpots()
		return true
	}

	rank := rankedStudent.Rank()
	tmp := heap.Pop(pq).(common.RankedStudent)
	heap.Push(pq, tmp)
	lastRank := tmp.Rank()

	switch {
	case rank == lastRank:
		heap.Push(pq, rankedStudent)
		return true
	case rank > lastRank:
		return false
	}

	heap.Push(pq, rankedStudent)
	count := strategy.countBeingRemovedReplicas()

	for ; count > 0; count-- {
		rs := heap.Pop(pq).(common.RankedStudent)
		rs.Student().ClearCourse()
	}

	return true
}
