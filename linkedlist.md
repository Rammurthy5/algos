# what 
Linkedin List comprised of nodes (data + pointer ref to next node).
the advantage of linked list pointer ref in nodes are not to worry about scattered memory location of the peer nodes. this makes linked list preferable over an array.
linked list is foundation for stacks, queues, graphs, and trees.

why do sometimes we have two pointers in a node? (doubly-linked list).
so that reverse iteration is also faster.

## Tail and Head in Linked list structure
Head present and points to 1st node is common in linked list implementation.
Head and Tail might be present at times for Insert time optimisation.
No head and no tail is uncommon but possible. e.g circular linked list where the last node connects back to the first. 

## circular linked list
can be singly or doubly LL. common usecases are: CPU Scheduling (Round-Robin Algorithm); Multiplayer Games (Turn-Based Systems); Memory Management (Circular Buffers a.k.a Ring Buffers use a circular LL to manage memory efficiently.It’s commonly used in audio/video streaming, networking, and real-time systems.
e.g.
Networking: Storing packets in a fixed-size buffer.Audio processing: Handling real-time sound streams.); Task Scheduling in Real-Time Systems; Implementing Undo-Redo Features;Traffic Light Systems;Data Streaming (Media players use circular linked lists to handle continuous playback where after the last song/video, it loops back to the first one).

## Complexity
Operation	                Singly Linked List	                                 Doubly Linked List
Create (Insertion) at Head	O(1)	                                                 O(1)
Create (Insertion) at Tail	O(n) (without tail pointer) / O(1) (with tail pointer)	 O(1)
Insert at Middle	        O(n)	                                                 O(n)
Read/Search	                O(n)	                                                 O(n)
Update	                    O(n)	                                                 O(n)
Delete at Head	            O(1)	                                                 O(1)
Delete at Tail	            O(n) (without tail pointer) / O(1) (with tail pointer)	 O(1)
Delete from Middle	        O(n)	                                                 O(n)