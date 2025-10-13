def kadane_algorithm(arr):
    max_current = max(arr[0], 0)
    max_global = max_current
    for i in arr[1:]:
        max_current = max(i, max_current + i)
        if max_current > max_global:
            max_global = max_current
    return max_global   

if __name__ == "__main__":
    import sys
    # Example: python3 src/kadanealgorithm.py 1 -2 3 4 -1 2 1 -5 4
    arr = list(map(int, sys.argv[1:]))
    result = kadane_algorithm(arr)
    print(result)