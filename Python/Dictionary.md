# Dictionary

# keys

Returns the keys in a list.

```python
dictionary_keys = dictionary.keys()
```

# values

Returns the values in a list.

```python
dictionary_values = dictionary.values()
```

# clear

Removes all elements from dictionary.

```python
dictionary_values.clear()
```

# copy

Returns a copy of dictionary.

```python
dictionary_copy = dictionary.copy()
```

# fromkeys

Returns a dictionary with specified keys and specified value.
```python
x = ('key1', 'key2', 'key3')
y = 5

thisdict = dict.fromkeys(x, y)

print(thisdict)

# {'key1': 5, 'key2': 5, 'key3': 5}
```

# get

Key name is the key it gets, value is the thing returned if doesn't exist. Value deefaults to None.

```python
dictionary.get(keyname, value)
```

# items

Returns a view object. Contains the key-value pairs of the dictionary as tuples in a list.

```python
car = {
  "brand": "Ford",
  "model": "Mustang",
  "year": 1964
}

x = car.items()

print(x)

# dict_items([('brand', 'Ford'), ('model', 'Mustang'), ('year', 1964)])
```
# pop

Removes the element with specified key. If value does not exist, return defaultValue. If not specified, error is raised.

```python
dictionary.pop(keyname, defaultvalue)
```

# popitem

Removes the item that was last inserted into the dictionary. Prior to 3.7, removes a random item.

# setdefault

Returns the value of item with specified key. If doesn't exist, insert key, with specified value. The keyname is the item you want to return the value from, the value is optional. If key exists, parameter has no effect. Otherwise, value becomes the key value.

```python
dictionary.setdefault(keyname, value)
```

# update

Inserts/updates the specified objects to the dictonary.

```python
dictionary.update(iterable)
dictionary.update({1:1, 2:2, 3:3})
```