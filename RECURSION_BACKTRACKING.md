# RECURSION
method calls itself until the stack overflows. to stop that, we add a constraint.Typical stack sizes Linux & Mac: usually 8 MB per thread.
Windows: 1 MB.
You can also set a custom stack size at compile or runtime. ulimit -s    # show stack size (in KB)
ulimit -s 16384   (set it to 16 MB)

# patterns 
all multiple call recursion will be based on two scenarios - take and not-take. 

## time complexity
depends on how many recursive calls are made and the amount of work done at each step. To analyze it systematically, we usually use recursion tree. 

# Backtracking
Backtracking is advanced form of recursion. It explores all possible solutions *incrementally* and backtracks (by pruning) as soon as it fails to meet constraints that can't lead to solution.