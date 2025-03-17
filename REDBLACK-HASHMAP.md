# Hash Maps
uses a normal array (bucket list) to store key-value pair. each element in the bucket list is a linked list (bucket). it does auto-resizing when the size of the bucket list hits 75% full.
stores as key-value pair, relies on hashing function to determine where to store each value in an underlying array. challenges are handling collisions (multiple keys produce same index hash) in hashing. Chaining (Store multiple values in a linked list at the same index), open-addressing (Find another empty slot using a probing strategy) are techniques to handle collision. Java uses chaining, python uses open-addressing specialized probing, golang uses open-addressing quadratic-probing.
a normal hashmap isnt thread-safe, not designed for concurrency. it needs external help to handle that. concurrentHashMaps can be used and they're designed for handling scale.
## Resizing
 When the load factor exceeds a certain threshold (75%), hash maps need to resize, which involves rehashing all entries. This operation can take O(n) time, although it's amortized over insertions to maintain an average case of O(1) for insertions. 
Standard initial capacity of HashMaps in Java8 is 16, Python & Golang offer 1 bucket (8).
## Insertion, Deletion, and Lookup
 hash maps provide average-case time complexity of O(1) for insertion, deletion, and lookup operations, assuming a good hash function with minimal collisions. This is because the key is hashed to determine its storage location directly. However, in the worst case, due to collisions, these operations can degrade to O(n) where n is the number of elements in the hash map. Techniques like chaining or open addressing are used to resolve collisions.


# Red-Black Trees (Self-Balancing BST)
## Insertion, Deletion, and Lookup
 being a type of self-balancing binary search tree, guarantee a worst-case time complexity of O(log n) for these operations. This is due to their balanced nature which ensures that the height of the tree remains relatively small compared to the number of nodes.
## Rebalancing
 After insertion or deletion, rebalancing might be necessary to maintain the tree's properties, which includes color changes and rotations. These operations are efficient, requiring at most O(log n) time in the worst case for rebalancing, but typically less due to the efficiency of the rebalancing process.

For most practical applications, hash maps are preferred for their average-case performance when dealing with random data and when order does not matter. However, red-black trees are more suitable in scenarios where order is important or when you need guaranteed logarithmic time complexity for operations, especially in cases where the worst-case scenario must be minimized, like in real-time systems or file systems. 

Remember, the choice between the two often depends on the specific requirements of the application, such as the need for ordered data, the frequency of operations, and the acceptable trade-offs between speed and space.

# Real-time usecase for Red-black trees Distributed File System (e.g., Google File System, HDFS)
### Order Preservation
 Red-black trees maintain a sorted order of elements, which is crucial in scenarios where data needs to be accessed or processed in a specific order. Hash maps do not inherently maintain any order.
### Guaranteed Performance
 While hash maps offer O(1) average case, red-black trees provide a guaranteed O(log n) time complexity for all operations (insert, delete, search) in the worst case. This predictability is vital in real-time systems where performance must be consistent.
### Range Queries
 Red-black trees are efficient for range queries or finding the next/previous element in a sequence, which is less straightforward with hash maps.
### Balanced Structure
 The self-balancing nature of red-black trees ensures that the tree remains efficient even after many insertions and deletions, avoiding the potential for degradation in performance that can occur with hash maps due to poor hash function distribution or high collision rates.

 # Real-time usecase for HashMaps (Distributed Caching System (e.g., Memcached, Redis))
 ### Scenario
  In large-scale web applications, caching frequently accessed data can significantly reduce database load and improve response times. 
### Use
 A hash map (or a similar structure like a distributed hash table) is used to store key-value pairs where keys might be user session IDs, cache keys for database results, or API response data. 
### Why Hash Map
 The O(1) average case for insertion and retrieval is crucial for high throughput and low latency in caching scenarios. For instance, when a user logs in, their session data can be quickly retrieved or updated. Systems like Redis use a combination of data structures, but hash maps are fundamental for key-value storage.
### Citation
 This use case aligns with the general understanding of how caching works in distributed systems, but for specific references, look at documentation or articles on Memcached or Redis.
