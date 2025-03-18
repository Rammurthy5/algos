space-optimized variant of std trie. stores strings or sequences such as in dict, routing tables, autocomplete systems.
Nodes represent substrings, not single char like std trie.

## typical components of a radix trie node
node is of balanced tree or hashmap or array.
balanced tree will use balanced memory and lookup speed. 
array suits smaller data set. 
hashmap is best for fast lookups but higher memory.
each node has a edge, children pointers, terminal flag.

## Hashmap or Array or Balanced tree
Depends on the workload  - read-heavy or write-heavy or mixed.


## distributed radix trie implmentation
use hashmap.
distribute across two nodes.
Do insert, search, delete.

## time complexity
Space Efficiency | Lookup Speed  | Complexity  | Use Case                          |
|----------------|---------------|-------------|-----------------------------------|
  Good           | O(L)          | Moderate    | IP routing, space-sensitive       |