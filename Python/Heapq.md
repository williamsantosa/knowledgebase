# Heapq

Documentation [here](https://docs.python.org/3/library/heapq.html).

A heap is a tree-like data structure stored as an array. The children of the parent are stored in `pos * 2` and `pos * 2 + 1` of the array. A heap is either min or max, meaning (if min) that the parent is minimal compared to its children. This property recursively holds for all children.

## heapify

Changes a list into a min-heap. If want to make a max-heap, just multiply all elements of a list into its negation and when popping revert it again by multiplying by -1 again. The heapify will alter the list variable passed in.

```python
heap = [...]
heapq.heapify(heap)
```

## heappush

Pushes a value into the heap.

```python
heap = [...]
heapq.heapify(heap)
heapq.heappush(heap, val)
```

## heappop

Pops the smallest value from the heap, which is `heap[0]`.

```python
heap = [...]
heapq.heapify(heap)
popped_val = heapq.heappop(heap)
```

## heapreplace

Pop and return the smallest value from the heap and push the new value to the heap.

```python
heap = [...]
heapq.heapify(heap)
popped_val = heapq.replace(heap, val)
```

## merge

Merge multiple sorted inputs into a new single sorted output.

```python
heapq.merge(*iterables, key=None, reverse=False)
```

## nlargest

Return a list with the n largest elements from the dataset defined by iterable. key, if provided, specifies a function of one argument that is used to extract a comparison key from each element in iterable (for example, `key=str.lower`). Equivalent to: `sorted(iterable, key=key, reverse=True)[:n]`.

## nsmallest

Return a list with the n smallest elements from the dataset defined by iterable. key, if provided, specifies a function of one argument that is used to extract a comparison key from each element in iterable (for example, `key=str.lower`). Equivalent to: `sorted(iterable, key=key)[:n]`.