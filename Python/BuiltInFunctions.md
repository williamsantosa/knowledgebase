# General Methods

## sort

Changes the values in the list_values. Has key and reverse parameter. Key is the thing it'll sort by (takes input of function), `reverse=True` means descending.

```python
list_values.sort()
```

## sorted

Same as `sort()` but instead of changing a list in-place, it'll return the sorted version of it instead. Has same parameters *after* the initial list_values one.

```python
sorted_list_values = sorted(list_values)
```