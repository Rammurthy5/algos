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