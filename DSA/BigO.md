# Big O

How much time and space does it need?

## Upper Bound

| Type         | Worst Case Runtime |
| ------------ | ------------------ |
| Constant     | O(1)               |
| Logarithmic  | O(log(n))          |
| Linear       | O(n)               |
| Linearithmic | O(nlog(n))         |
| Quadric      | O(n^2)             |
| Cubic        | O(n^3)             |
| Exponential  | O(2^n)             |
| Factorial    | O(n!)              |

```python
for i in range(0, n, 1):
  for j in range(i, n, 1):
    # do something
```

The worst case for the above is O(n^2) because it represents the arithmetic series.