# Sliding Window Technique

Check [here](https://www.youtube.com/watch?v=jM2dhDPYMQM&ab_channel=QuanticDev).

1. Used for finding subarrays in an array that satisfy given conditions
2. Maintain a subset of items as your window and resize and move that window within the larger list until find a solution
3. Subset of DP (solving problem via dividing into sub problems)
4. Time complexity: O(n) = linear time
5. Space complexity: O(1) = constant space

Find subarrays that add up to 9.

Create a window that contains elements. Move right until >= 9. If == 9, found already. If > 9, then move left by 1. Repeat.

Subarrays are contiguous by definition, so elements are adjacent.
Input size could be anything.

ALWAYS ask questions to interviewer.

Given an array of integers, find maximum sum subarray of required size.

[-1, 2, 3, 1, -3, 2] Subarray size: 2

### 1 Brute Force

Time complexity is O(n) for hash table creation and O(1) for every lookup afterward. The space complexity is O(n).

### 2 Sliding Window

Same time complexity but space complexity is O(1).