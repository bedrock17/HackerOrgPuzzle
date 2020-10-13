package main

import (
	"fmt"
	"time"
)

func (pq *posQueue) create(size int) {
	pq.q = make([]pos, size)
	pq.size = size
	pq.start = 0
	pq.end = 0
}

func (pq *posQueue) put(p pos) {
	if pq.end+1 >= pq.start {
		pq.q[pq.end%pq.size] = p
		pq.end++
	} else {
		fmt.Println("PQ OVERFLOW!", p, pq.start, pq.end, pq.size)
		time.Sleep(2000 * time.Millisecond)
	}
}

func (pq *posQueue) get() pos {
	var p = pos{-1, -1, -1}
	if pq.start+1 <= pq.end {
		p = pq.q[pq.start%pq.size]
		pq.start++
	} else {
		fmt.Println("PQ EMPTY")
	}
	return p
}
func (pq *posQueue) length() int {
	return pq.end - pq.start
}
