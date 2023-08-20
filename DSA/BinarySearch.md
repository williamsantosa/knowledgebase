# Binary Search

Left, right initialize to start and end of list. While left <= right or left < right (etc), check mid by integer division of left and right. Compare left to mid and/or right to mid or mid to target and do the appropriate manipulation based off that. If searchinhg for a value that exists, change right = mid - 1 if mid > target, etc.