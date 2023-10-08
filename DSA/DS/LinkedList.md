# Linked Lists

Most of the LL algorithms will employ the use of pointers. Most use either 1 or 2. Singly linked is only one direction, doubly both directions.

## Insert

1. New node point to next
2. Change current node to new node
3. If doubly LL, change new node prev to curr and next node prev to new node

## Remove

1. Create 2 pointers, trav1 and trav2 located at start and start + 1
2. Advance trav2 until find node you want to remove
3. Assign temp to node we want to remove
4. Advance trav2 by 1
5. Make trav1 point to trav2
6. If doubly, make trav2 point to trav1
7. Delete temp