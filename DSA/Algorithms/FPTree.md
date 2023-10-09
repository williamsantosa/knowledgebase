# Frequent Pattern Growth (Algorithm) Tree

Reminder: Duplicates items in the same transaction (by default) do not occur. ITems either occur zero or once. If seen twice, it only counts once per transaction.

Apriori drawbacks:

1. Candidate sets are built at each step.
2. Algorithm has to repeatedly scan the datbase to build candidate sets.

## Steps

Steps to do the algorithm.

### 1. Initialize

| TID   | Items           |
| ----- | --------------- |
| $T_1$ | \{E,K,M,N,O,Y\} |
| $T_2$ | \{D,E,K,N,O,Y\} |
| $T_3$ | \{A,E,K,M\}     |
| $T_4$ | \{C,K,M,U,Y\}   |
| $T_5$ | \{C,E,I,K,O,O\} |

min_support = 3

confidence = 60%

## 2. Construct ordered-item sets

First step of frequent pattern growth tree is the same as the first step in Apriori Algorithm (create set of 1-itemsets and their support count). The set is then sorted in descending order based off support count.

| itemset | sup_count |
| ------- | --------- |
| A       | 1         |
| C       | 2         |
| D       | 1         |
| E       | 4         |
| I       | 1         |
| K       | 5         |
| M       | 3         |
| N       | 2         |
| O       | 3         |
| U       | 1         |
| Y       | 3         |

Prune itemsets with sup_count < min_support and sort in descending order. Set L = \{
  \{K:5\}, 
  \{E:4\},
  \{M:3\}, 
  \{O:3\}, 
  \{Y:3\}
\}.

In the original dataset, create the respective ordered-item set by iterating over frequent pattern set and hcecking if current item is in transaction. If contained, item is inserted in ordered-item set for current transaction.

| TID   | Items           | Ordered-Item Set |
| ----- | --------------- | ---------------- |
| $T_1$ | \{E,K,M,N,O,Y\} | \{K,E,M,O,Y\}    |
| $T_2$ | \{D,E,K,N,O,Y\} | \{K,E,O,Y\}      |
| $T_3$ | \{A,E,K,M\}     | \{K,E,M\}        |
| $T_4$ | \{C,K,M,U,Y\}   | \{K,M,Y\}        |
| $T_5$ | \{C,E,I,K,O,O\} | \{K,E,O\}        |

## 3. Construct Trie (Look at link)

Create a Trie starting with null and with item : count for each transaction's ordered-item set. If none, then create new node with support count of 1. There can be multiple nodes with the same item depending on the linkage.

## 4. Mine the FP Tree by creating conditional pattern base.

Create in ascending order. The conditional pattern base the set of every path starting from null till the item, the keys being the path set (up till item) and value being the value of the node.

| Items | Conditional Pattern Base                  |
| ----- | ----------------------------------------- |
| Y     | \{\{K,E,M,O:1\}, \{K,E,O:1\}, \{K,M:1\}\} |
| O     | \{\{K,E,M:1\}, \{K,E:2\}\}                |
| M     | \{\{K,E:2\}, \{K:1\}\}                    |
| E     | \{K:4\}                                   |
| K     |                                           |

## 5. Create conditional frequent pattern tree

Create conditional frequent pattern tree by taking set of all elements common in all paths and calculating the support count by summing the support counts of all paths in conditional pattern base.

| Items | Conditional Pattern Base                  | Conditional Frequent Pattern Tree |
| ----- | ----------------------------------------- | --------------------------------- |
| Y     | \{\{K,E,M,O:1\}, \{K,E,O:1\}, \{K,M:1\}\} | \{K:3\}                           |
| O     | \{\{K,E,M:1\}, \{K,E:2\}\}                | \{K,E:3\}                         |
| M     | \{\{K,E:2\}, \{K:1\}\}                    | \{K:3\}                           |
| E     | \{K:4\}                                   | \{K:4\}                           |
| K     |                                           |                                   |

## 6. Create frequent pattern rules

Create frequeunt pattern rules by pairing conditional frequent pattern tree's set to corresponding to item as given in below table.

| Items | Frequent Pattern Generated                       |
| ----- | ------------------------------------------------ |
| Y     | \{\<K,Y:3\>\}                                    |
| O     | \{\{\<K,O:3\>\}, \{\<K,E:3\>\}\{\<K,E,O:3\>\} \} |
| M     | \{\<K,M:3\>\}                                    |
| E     | \{\<E,K:4\>\}                                    |
| K     |                                                  |

The possible association rules are inferred the same way as Apriori for each frequent pattern. That is, set $A,B,C$ exist such that $C =$ frequent pattern, $A,B \subset C$, $A \cup B = C$, and $A \cap B = \emptyset$.

For example, for item $Y$, the possible association rules are:

- $Y \rightarrow K$
- $K \rightarrow Y$

Calculate the confidence of both.

- $\text{Confidence}(Y \rightarrow K) = \frac{\text{Sup}(Y \cup K)}{\text{Sup}(Y)} \cdot 100\% = \frac{3}{3} \cdot 100\% = 100\%$
- $\text{Confidence}(K \rightarrow Y) = \frac{\text{Sup}(Y \cup K)}{\text{Sup}(K)} \cdot 100\% = \frac{3}{5} \cdot 100\% = 60\%$