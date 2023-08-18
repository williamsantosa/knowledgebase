# Set

## add

Adds an element to the set 

```python
set_var.add(val)
```

## clear

Removes all elements from the set.

```python
set_var.clear()
```

## copy

Return a copy of the set.

```python
copy_var = var.copy()
```

## difference

Returns a set containing the difference between two or more sets.
Basically, only in the first and not the second.

```python
diff = set1.difference(set2)
```

## difference_update

Removes the items in the set that are also included in another specified set.

Same as the `difference()` method but doesn't return, it removes it from the original set.

## discard

Removes the specified item.

```python
set1.remove(val)
```

## intersection

Returns the intersection of two sets.

```python
intersect = set1.intersection(set2)
```

## intersection_update

Replaces set1 with the intersection of set1 and set2.

```python
set1.intersection_update(set2)
```

## isdisjoint

Returns True if none of the items are present in both sets, otherwise False.

```python
res = set1.isdisjoint(set2)
```

## issubset

Returns whether another set contains this set or not.

If set1 is a subset of set2.

```python
res = set1.issubset(set2)
```

## issuperset

Returns whether anothe set is a superset of this set.

If set1 is a superset of set2 (if set2 is a subset of set1).

```python
res = set1.issuperset(set2)
```

## pop

Remove random item from set and return it.

```python
x = set1.pop()
```

## remove

Remove a specified element from set. Will raise error if specified item does not exist. Different from discard because it will raise error.

```python
set1.remove(element)
```

## symmetric_difference

Return a set containing all items from both sets, except items present in both sets.

```python
x = {"apple", "banana", "cherry"}
y = {"google", "microsoft", "apple"}

z = x.symmetric_difference(y)

print(z)
```

## symmetric_difference_update

Same as above by removing items present in both sets and inserting the other items.

## union

Returns a set containing all items from both sets.

```python
x = {"apple", "banana", "cherry"}
y = {"google", "microsoft", "apple"}

z = x.union(y)

print(z)
```

## update

Update the current set by adding items from another set or other iterable.

```python
x = {"apple", "banana", "cherry"}
y = {"google", "microsoft", "apple"}

x.update(y)

print(x)
```