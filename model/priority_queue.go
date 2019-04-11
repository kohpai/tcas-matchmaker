package model

import "fmt"

type RankedStudent struct {
	student *Student
	rank    uint16
	next    *RankedStudent
	prev    *RankedStudent
}

type PriorityQueue struct {
	head *RankedStudent
	tail *RankedStudent
	len  uint16
}

func (pq *PriorityQueue) Head() *RankedStudent {
	return pq.head
}

func (pq *PriorityQueue) Tail() *RankedStudent {
	return pq.tail
}

func (pq *PriorityQueue) Len() uint16 {
	return pq.len
}

// - returns the highest-rank RankedStudent where its rank is less than or
//   equal to the given rank.
// - If no such RankedStudent found, returns the nearest
//   higher-rank RankedStudent
// - returns nil if the PriorityQueue is empty
func (pq *PriorityQueue) Find(rank uint16) *RankedStudent {
	for current := pq.head; current != nil; current = current.next {
		if current.rank <= rank {
			continue
		}

		if current.prev == nil {
			return current
		}
		return current.prev
	}
	return nil
}

func (pq *PriorityQueue) Push(rs *RankedStudent) {
	current := pq.Find(rs.rank)

	switch {
	// pq is empty
	case current == nil:
		pq.head = rs
		pq.tail = rs
		rs.next = nil
		rs.prev = nil
	// rs is head
	case current.rank > rs.rank:
		rs.next = current
		rs.prev = nil
		current.prev = rs
		pq.head = rs
	// rs is in between or tail
	case current.rank <= rs.rank:
		rs.next = current.next
		rs.prev = current
		current.next = rs
		if rs.next == nil {
			pq.tail = rs
		}
	}

	pq.len += 1
}

func (pq *PriorityQueue) Pop() *RankedStudent {
	if pq.head == nil {
		return nil
	}

	poping := pq.head
	pq.head = poping.next
	pq.len -= 1

	return poping
}

func (rs *RankedStudent) String() string {
	return fmt.Sprintf(
		"{student: %s, rank: %d}",
		rs.student,
		rs.rank,
	)
}
