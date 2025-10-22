# print all subsequences of a string
def print_subsequences(s, current="", index=0):
    if index == len(s):
        print(current)
        return
    # Include the current character
    print_subsequences(s, current + s[index], index + 1)
    # Exclude the current character
    print_subsequences(s, current, index + 1)

# print the sequnces that sum upto the target
def print_sequences_target(index, arr, t_arr, target, curr_sum):
    if index >= len(arr):
        if curr_sum==target:    
            print(t_arr)
        return

    t_arr.append(arr[index])
    print_sequences_target(index+1, arr, t_arr, target, curr_sum+arr[index])
    t_arr.pop()
    print_sequences_target(index+1, arr, t_arr, target, curr_sum)

# print at least one sequence that sums upto the target
def print_one_sequence_target(index, arr, t_arr, target, curr_sum):
    if index >= len(arr):
        if curr_sum==target:    
            print(t_arr)
            return True
        else:
            return False

    t_arr.append(arr[index])
    if print_one_sequence_target(index+1, arr, t_arr, target, curr_sum+arr[index]):
        return True
    t_arr.pop()
    if print_one_sequence_target(index+1, arr, t_arr, target, curr_sum):
        return True
    return False

# count the number of sequences that sum upto the target
def count_sequences_target(index, arr, target, curr_sum):
    if index >= len(arr):
        if curr_sum==target:    
            return 1
        else:
            return 0

    curr_sum+=arr[index]
    l = count_sequences_target(index+1, arr, target, curr_sum)
    curr_sum-=arr[index]
    r = count_sequences_target(index+1, arr, target, curr_sum)
    return l+r  

# combination sum I t.c O(2^t *k) s.c O(t*k)
def combination_sum(index, target, arr, t_arr):
    if target == 0:
        return [t_arr.copy()]   # Return a list with a copy of the valid combination
    if index >= len(arr):
        return []               # No combination found

    result = []

    if arr[index] <= target:
        t_arr.append(arr[index])
        result += combination_sum(index, target - arr[index], arr, t_arr)  # include current
        t_arr.pop()

    result += combination_sum(index + 1, target, arr, t_arr)  # exclude current

    return result

print(combination_sum(0, 7, [2, 3, 6, 7], []))

# combination sum II tc O(log n + (k⋅n))
def combination_sum_2(index, target, arr, t_arr):
    result = []
    
    if target == 0:
        return [t_arr.copy()]
    
    for i in range(index, len(arr)):
        # Skip duplicates (only for the same recursive level)
        if i > index and arr[i] == arr[i - 1]:
            continue
        if arr[i] > target:
            break  # early stopping due to sorted array
        t_arr.append(arr[i])
        result += combination_sum_2(i + 1, target - arr[i], arr, t_arr)
        t_arr.pop()
    return result



if __name__ == "__main__":
    import sys
    # Example: python3 src/recursion.py subsequences abc
    # Example: python3 src/recursion.py print_sequences_target 10 1 2 3 4 5
    # Example: python3 src/recursion.py print_one_sequence_target 10 1 2 3 4 5
    # Example: python3 src/recursion.py count_sequences_target 10 1 2 3 4 5
    command = sys.argv[1]
    if command == "subsequences":
        s = sys.argv[2]
        print_subsequences(s)
    elif command == "print_sequences_target":
        target = int(sys.argv[2])
        arr = list(map(int, sys.argv[3:]))
        print_sequences_target(0, arr, [], target, 0)
    elif command == "print_one_sequence_target":
        target = int(sys.argv[2])
        arr = list(map(int, sys.argv[3:]))
        print_one_sequence_target(0, arr, [], target, 0)
    elif command == "count_sequences_target":
        target = int(sys.argv[2])
        arr = list(map(int, sys.argv[3:]))
        result = count_sequences_target(0, arr, target, 0)
        print(result)   
