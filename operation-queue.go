package main

import (
	"sync"
)

type QueueNode struct {
	fn func()
	t string
}

type Queue struct {
    m sync.Mutex
    q []QueueNode
}

func NewQueue() *Queue {
    return &Queue{
        q: []QueueNode{},
    }
}

func (q *Queue) Enqueue(f func(), t string) {
    q.m.Lock()
    defer q.m.Unlock()
    q.q = append(q.q, QueueNode{fn: f, t: t})
}

func (q *Queue) Dequeue() (QueueNode, error) {
    q.m.Lock()
    defer q.m.Unlock()
    if len(q.q) == 0 {
        return QueueNode{}, nil
    }
    n := q.q[0]
    q.q = q.q[1:]
    return n, nil
}