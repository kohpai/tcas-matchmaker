package model

import "fmt"

type RankedStudent struct {
	student *Student
	rank    uint16
	index   int
}

type PriorityQueue []*RankedStudent

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].rank < pq[j].rank
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	rankedStudent := x.(*RankedStudent)
	rankedStudent.index = len(*pq)
	*pq = append(*pq, rankedStudent)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	l := len(old)
	rankedStudent := old[l-1]
	rankedStudent.index = -1
	*pq = old[0 : l-1]

	return rankedStudent
}

func (rs *RankedStudent) String() string {
	return fmt.Sprintf(
		"{student: %s, rank: %d, index: %d}",
		rs.student,
		rs.rank,
		rs.index,
	)
}
