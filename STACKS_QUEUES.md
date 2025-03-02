stacks can be implemented with arrays or LL following LIFO. 

## time complexity
O(1) for Pop, Peek, Push, isEmpty.

## real-time usecases
recursion, undo / redo, maze-solving, pathfinding, backtracking algorithms

# Queue
queues are similarly based on arrays or LL following FIFO. if queues are to be implemented with arrays, it can be circular array with 2 pointers or a fixed array.
there are couple of variants with the queue: priority queue and deque (double-ended queue). Deque allows adding/removing from both ends, blending stack and queue behaviors. priority queue is based on priority, not just order.

## time complexity
O(1) for enqueue, peek, isEmpty.
O(n) for dequeue if its fixed array. 
O(1) for circular array or LL.

## real-time usecases
message queues, task scheduling.