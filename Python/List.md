# List

Methods to manipulate a list. A list is an ordered container of items. Information is taken from [w3schools](https://www.w3schools.com/python/python_ref_list.asp).

## append()

Appends an item into the end of a list.

```python
animals = ['cat', 'dog', 'rabbit']
animals.append('monkey')
print(animals)

# ['cat', 'dog', 'rabbit', 'monkey']
```

### Syntax

```python
list.append(elmnt)
```

### Parameter Values

| Parameter | Description                                                   |
| --------- | ------------------------------------------------------------- |
| elmnt     | Required. An element of any type (string, number, object etc) |

## clear()

Removes all the elements from a list.

```python
animals = ['cat', 'dog', 'rabbit']
animals.clear()
print(animals)

# []
```

### Syntax

```python
list.clear()
```

## copy()

Return a copy of the specified list.

```python
animals = ['cat', 'dog', 'rabbit']
x = animals.copy()

print(x)
# ['cat', 'dog', 'rabbit']
```

### Syntax

```python
list.copy()
```

## count()

Returns the number of elements with the specified value.

```python
points = [1, 4, 2, 9, 7, 8, 9, 3, 1]
x = points.count(9)

print(x)
# 2
```

### Syntax

```python
list.count(value)
```

### Parameter Values

| Parameter | Description                                                                      |
| --------- | -------------------------------------------------------------------------------- |
| value     | Required. Any type (string, number, list, tuple, etc.). The value to search for. |

## extend()

Adds the specified list elements (or any iterable) to the end of the current list.

```python
fruits = ['apple', 'banana', 'cherry']
cars = ['Ford', 'BMW', 'Volvo']
fruits.extend(cars)

print(fruits)
# ['apple', 'banana', 'cherry', 'Ford', 'BMW', 'Volvo']
```

### Syntax

```python
list.extend(iterable)
```

### Parameter Values

| Parameter | Description                                     |
| --------- | ----------------------------------------------- |
| iterable  | Required. Any iterable (list, set, tuple, etc.) |

## index()

Returns the position of the first occurrence of the specified value, otherwise raise a value error.

```python
fruits = ['apple', 'banana', 'cherry']

x = fruits.index("cherry")
# 2
```

### Syntax

```python
list.index(elmnt)
```

### Parameter Values

| Parameter | Description                                                                |
| --------- | -------------------------------------------------------------------------- |
| elmnt     | Required. Any type (string, number, list, etc.). The element to search for |


## insert()

Inserts an element into the specified value at the specified position.

```python
fruits = ['apple', 'banana', 'cherry']
fruits.insert(1, "orange")

print(fruits)
# ['apple', 'orange', 'banana', 'cherry']
```

### Syntax

```python
list.insert(pos, elmnt)
```

### Parameter Values

| Parameter | Description                                                                |
| --------- | -------------------------------------------------------------------------- |
| pos       | Required. A number specifying in which position to insert the value        |
| elmnt     | Required. Any type (string, number, list, etc.). The element to search for |

## pop()

Removes the element at the specified position and returns it.

```python
fruits = ['apple', 'banana', 'cherry']
fruits.pop(1)

print(fruits)
# ['apple, 'cherry']
```

### Syntax

```python
list.pop(pos)
```

### Parameter Values

| Parameter | Description                                                                                                                    |
| --------- | ------------------------------------------------------------------------------------------------------------------------------ |
| pos       | Optional. A number specifying the position of the element you want to remove, default value is -1, which returns the last item |

## remove()

Removes the first occurrence of the specified element within the list. Otherwise raise value error.

```python
fruits = ['apple', 'banana', 'cherry']
fruits.remove("banana")

print(fruits)
# ['apple', 'cherry']
```

### Syntax

```python
list.remove(elmnt)
```

### Parameter Values

| Parameter | Description                                                                   |
| --------- | ----------------------------------------------------------------------------- |
| elmnt     | Required. Any type (string, number, list etc.) The element you want to remove |

## reverse()

Reverses the sorting order of the elements.

```python
fruits = ['apple', 'banana', 'cherry']
fruits.reverse()

print(fruits)
# ['cherry', 'banana', 'apple']
```

### Syntax

```python
list.reverse()
```

## sort()

Sorts the list by ascending by default. Can input a function to change sorting criteria(s).

```python
cars = ['Ford', 'BMW', 'Volvo']
cars.sort()

print(cars)
# ['BMW', 'Ford', 'Volvo']
```

### Syntax 

```python
list.sort(reverse=True|False, key=myFunc)
```

### Parameter Values

| Parameter | Description                                                                    |
| --------- | ------------------------------------------------------------------------------ |
| reverse   | Optional. reverse=True will sort the list descending. Default is reverse=False |
| key       | Optional. A function to specify the sorting criteria(s)                        |