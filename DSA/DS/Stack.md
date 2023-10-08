# Stacks

Main operations are push and pop. A stack generally has those two plus a pointer to the top.

```
                |------|
top pointer --> | data | 
                |------|
                | data | 
                |------|
                | data | 
                |------|
                | data | 
                |------|
```

Last in, first out (LIFO).

Push, pop, peek, size: O(1)
Search: O(n)

Problem with multiple stacks is Towers of Hanoi.