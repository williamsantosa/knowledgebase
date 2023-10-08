# Apriori Algorithm

https://www.geeksforgeeks.org/apriori-algorithm/#

Every non-empty subset of frequent itemset must be frequent. Key concept of Apriori algorithm is anti-monotonicity of support measure. Apriori assumes:
- All subsets of frequent itemset must be frequent
- If an itemset is infrequent, its supersets are infrequent

## Steps

Steps to do the algorithm.

### 1. Initialize

1. Dataset (table{TID, items})
2. Minimum support count (int)
3. Minimum confidence (%)

| TID | Items                |
| --- | -------------------- |
| T1  | $l_1, l_2, l_5$      |
| T2  | $l_2, l_4$           |
| T3  | $l_2, l_3$           |
| T4  | $l_1, l_2, l_4$      |
| T5  | $l_1, l_3$           |
| T6  | $l_2, l_3$           |
| T7  | $l_1, l_3$           |
| T8  | $l_1, l_2, l_3, l_5$ |
| T9  | $l_1, l_2, l_3$      |

min_support = 2

min_confidence = 50%

### 2. Compute Algorithm

#### K = 1

Create candidate set $C_1$ from original itemset.

| Itemset | sup_count |
| ------- | --------- |
| $l_1$   | 6         |
| $l_2$   | 7         |
| $l_3$   | 6         |
| $l_4$   | 2         |
| $l_5$   | 2         |

Create itemset $L_1$ by trimming itemsets with sup_count < min_support.

| Itemset | sup_count |
| ------- | --------- |
| $l_1$   | 6         |
| $l_2$   | 7         |
| $l_3$   | 6         |
| $l_4$   | 2         |
| $l_5$   | 2         |

#### K = 2

Create candidate set $C_2$ from $L_1$ self join by checking that it has (K-2) elements in common. Therefore, no elements should match because 2-2 = 0, 0 elements must match between itemsets.

| Itemset   | sup_count |
| --------- | --------- |
| $l_1,l_2$ | 4         |
| $l_1,l_3$ | 4         |
| $l_1,l_4$ | 1         |
| $l_1,l_5$ | 2         |
| $l_2,l_3$ | 4         |
| $l_2,l_4$ | 2         |
| $l_2,l_5$ | 2         |
| $l_3,l_4$ | 0         |
| $l_3,l_5$ | 1         |
| $l_4,l_5$ | 0         |

Create itemset $L_2$ by trimming itemsets with sup_count < min_support.

| Itemset   | sup_count |
| --------- | --------- |
| $l_1,l_2$ | 4         |
| $l_1,l_3$ | 4         |
| $l_1,l_5$ | 2         |
| $l_2,l_3$ | 4         |
| $l_2,l_4$ | 2         |
| $l_2,l_5$ | 2         |

#### K = 3

Create candidate set $C_3$ from $L_2$ self join by checking that it has (K-2) elements in common. Therefore, first element should match because 3-2 = 1, 1 element must match between itemsets.

| Itemset       | sup_count |
| ------------- | --------- |
| $l_1,l_2,l_3$ | 2         |
| $l_1,l_2,l_5$ | 2         |
| $l_1,l_3,l_5$ | 0         |
| $l_2,l_3,l_4$ | 0         |
| $l_2,l_3,l_5$ | 1         |
| $l_2,l_4,l_5$ | 0         |

Create itemset $L_3$ by trimmming itemsets with sup_count < min_support.

| Itemset       | sup_count |
| ------------- | --------- |
| $l_1,l_2,l_3$ | 2         |
| $l_1,l_2,l_5$ | 2         |

#### K = 4

Create candidate set $C_4$ from $L_3$ self join by checking that it has (K-2) elements in common. Therefore, first two elements should match because 4-2 = 2, 2 element must match between itemsets.

| Itemset           | sup_count |
| ----------------- | --------- |
| $l_1,l_2,l_3,l_5$ | 1         |

Create itemset $L_4$ by trimming itemsets with sup_count < min_support.

| Itemset | sup_count |
| ------- | --------- |
|         |           |

There exists no frequent itemset in $L_4$. Therefore, are no frequent itemsets found further.

### 3. Identify association rules

Confidence of 60% means 60% of customers who purchased milk and bread also bought butter.

The following rule holds:

$$\text{Confidence}(A \rightarrow B)=\frac{\text{Sup}(A \cup B)}{\text{Sup}(A)}$$

We can show the rule generation by taking the itemset $C$ and computing all possible rules from its subset where $A, B \subset C$, $A \cap B = \emptyset$ and $A \cup B = C$.

For the last populated itemset $L_3$, we take the example of frequent itemset $\{l_1, l_2, l_3\}$.

Possible rules and its confidence:

- $\text{Confidence}(l_1 \cup l_2 \rightarrow l_3) = \text{Sup}(\frac{l_1 \cup l_2 \cup l_3}{l_1 \cup l_2}) = \frac{2}{4} \cdot 100 = 50\%$
- $\text{Confidence}(l_1 \cup l_3 \rightarrow l_2) = \text{Sup}(\frac{l_1 \cup l_2 \cup l_3}{l_1 \cup l_3}) = \frac{2}{4} \cdot 100 = 50\%$
- $\text{Confidence}(l_2 \cup l_3 \rightarrow l_1) = \text{Sup}(\frac{l_1 \cup l_2 \cup l_3}{l_2 \cup l_3}) = \frac{2}{4} \cdot 100 = 50\%$
- $\text{Confidence}(l_1 \rightarrow l_2 \cup l_3) = \text{Sup}(\frac{l_1 \cup l_2 \cup l_3}{l_1}) = \frac{2}{6} \cdot 100 = 33\%$
- $\text{Confidence}(l_2 \rightarrow l_1 \cup l_3) = \text{Sup}(\frac{l_1 \cup l_2 \cup l_3}{l_2}) = \frac{2}{7} \cdot 100 = 28\%$
- $\text{Confidence}(l_3 \rightarrow l_1 \cup l_2) = \text{Sup}(\frac{l_1 \cup l_2 \cup l_3}{l_3}) = \frac{2}{6} \cdot 100 = 33\%$

Strong association rules (50% confidence rate):

- $l_1 \cup l_2 \rightarrow l_3$
- $l_1 \cup l_3 \rightarrow l_2$
- $l_2 \cup l_3 \rightarrow l_1$

## Pros and Cons

Pros:
- Scalability
- Simplicity
- Robustness
- Versatility

Cons:
- Time-consuming due to scanning dataset when creating candidate sets
- Memory-intensive due to having to hold many candidate/level sets
- Low-performance due to exponential time complexity $O(2^D)$ where $D$ is the horizontal width
- Low minimum support or large itemsets reduce efficiency, generating many more candidates which need to be tested and accumulated.
- Detecting a frequent pattern of size 100 $(v_1, v_2, ..., v_{100})$, requires it to generate $2^{100}$ candidate itemsets.