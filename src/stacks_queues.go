package src

import "errors"

type Stack interface {
	Push(int)
	Pop() (int, error)
	Peek() (int, error)
	IsEmpty() bool
}

// Queue interface (FIFO: First In, First Out)
type Queue interface {
	Enqueue(int)
	Dequeue() (int, error)
	Peek() (int, error)
	IsEmpty() bool
}

// Stack implementation using LinkedList
type StackImpl struct {
	list *LinkedList
}

func NewStack() Stack {
	return &StackImpl{list: NewLinkedList()}
}

func (s *StackImpl) Push(data int) {
	s.list.Insert(data) // Insert at the end of the list
}

func (s *StackImpl) Pop() (int, error) {
	s.list.mu.Lock()
	defer s.list.mu.Unlock()

	if s.list.head == nil {
		return 0, errors.New("stack is empty")
	}

	// If only one node
	if s.list.head == s.list.tail {
		data := s.list.head.Data.(int)
		s.list.head = nil
		s.list.tail = nil
		return data, nil
	}

	// Traverse to the second-to-last node
	temp := s.list.head
	for temp.Next != s.list.tail {
		temp = temp.Next
	}
	data := s.list.tail.Data.(int)
	temp.Next = nil
	s.list.tail = temp
	return data, nil
}

func (s *StackImpl) Peek() (int, error) {
	s.list.mu.RLock()
	defer s.list.mu.RUnlock()

	if s.list.head == nil {
		return 0, errors.New("stack is empty")
	}
	return s.list.tail.Data.(int), nil
}

func (s *StackImpl) IsEmpty() bool {
	s.list.mu.RLock()
	defer s.list.mu.RUnlock()
	return s.list.head == nil
}

type QueueImpl struct {
	list *LinkedList
}

func NewQueue() Queue {
	return &QueueImpl{list: NewLinkedList()}
}

func (q *QueueImpl) Enqueue(data int) {
	q.list.Insert(data) // Inserts at tail
}

func (q *QueueImpl) Dequeue() (int, error) {
	q.list.mu.Lock()
	defer q.list.mu.Unlock()

	if q.list.head == nil {
		return 0, errors.New("queue is empty")
	}

	data := q.list.head.Data.(int)
	q.list.head = q.list.head.Next
	if q.list.head == nil {
		q.list.tail = nil // If list becomes empty
	}
	return data, nil
}

func (q *QueueImpl) Peek() (int, error) {
	q.list.mu.RLock()
	defer q.list.mu.RUnlock()

	if q.list.head == nil {
		return 0, errors.New("queue is empty")
	}
	return q.list.head.Data.(int), nil
}

func (q *QueueImpl) IsEmpty() bool {
	q.list.mu.RLock()
	defer q.list.mu.RUnlock()
	return q.list.head == nil
}

func StacksQueuesRun() {
	// Stack
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	for !stack.IsEmpty() {
		data, _ := stack.Pop()
		println(data)
	}

	// Queue
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)

	for !queue.IsEmpty() {
		data, _ := queue.Dequeue()
		println(data)
	}
}
