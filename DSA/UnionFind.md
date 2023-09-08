# Union Find

Two operations, union and find.

We designate a "representative" in each of the objects within the same group. The "representative" element represents the entire group it belongs to. For example if (1,0,3,4) has representative 0, then find(4) = 0. Two representatives belong to the same group if and only if they are in the same group.

We can then arrange these elements as nodes within a tree, with each node containing a value and parent variable. 

Rearrange the groups such that the root is the "representative." We arrange it such that each node has a parent.

The parent of the root is itself.

## Union(x, y)

Unions the groups containing x and y.

Union((0,1), (2,3)) = (0,1,2,3).

To union two trees we set the root of one tree to be the root of the other.

``````
find(2) = 5 # root
find(1) = 0

find(2).parent = find(1)
``````
The above makes it so that node 2 changes its parent to be find(1) which is 0. Since it is its parent, it becomes part of the group.

## Find(x)

Finds the group x belongs to.

To find what group a node belongs to, we start the node then continue until we reach the root node.

## How to implement?

Initialize with `parent[i] = i`.

```python
def find(x):
  if parent[x] != x:
    return find(parent[x])
  else:
    return x
```

```python
def union(x,y)
  find(y).parent = find(x)
```

Time complexity: O(n).

Space complexity: O(n).

Can optimize to O(log(n)) by using *ranks*, where we place the smaller subtree under the bigger one.

### Implementation

```python
class UnionFind:
  def __init__(self, n):
    self.parents = [i for i in range(n)]
    self.rank = [1 for _ in range(n)]

  def find(self, x):
    while x != self.parents[x]:
      self.parents[x] = self.parents[self.parents[x]] # optimize search, immediately go to the following parent above
      x = self.parents[x]
    return x
  
  def union(self, x1, x2):
    p1, p2 = self.find(x1), self.find(x2)
    if p1 == p2:
      return False
    elif self.rank[p1] > self.rank[p2]:
      self.parents[p2] = p1
      self.rank[p1] += self.rank[p2]
    else:
      self.parents[p1] = p2
      self.rank[p2] += self.rank[p1]
    return True
```