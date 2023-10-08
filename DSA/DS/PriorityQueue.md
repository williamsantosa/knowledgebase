# Priority Queues

ADT like queue but elements have a priority. Only supports comparative data. Must be ordered in some way (e.g ASC or DESC).

Instead of `dequeue` has `poll` which removes the element with highest priority in the priority queue.

Heap is a tree based data structure satisfying heap invariant property.

If node A is parent of node B, then A must be ordered with respect to B for all nodes A, B in heap.

Parent >= Child in max heap.
Child >= Parent in min heap.

Essentially, the relation for both parent and child follows the heap description for all parent and child.

Cannot have cycles. Each node has one parent.

Binary Heap construction: O(n)
Polling: O(log(n))
Peek: O(1)
Adding: O(log(n))