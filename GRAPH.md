A graph data structure is a non-linear collection of vertices (nodes) connected by edges (lines or arcs) that represent relationships between entities. 
Unlike trees or lists, a graph has no single root node or strict hierarchical order.


• Traversals:

	• Breadth-First Search (BFS): Explores neighbor nodes level by level (great for shortest path in unweighted graphs).
	• Depth-First Search (DFS): Explores as deep as possible down each branch before backtracking.

• Shortest Path:

	• Dijkstra's Algorithm or Bellman-Ford for weighted graph routes.

• Minimum Spanning Tree (MST):

	• Prim's or Kruskal's algorithms to connect all vertices with minimal total edge weight.


# Space Complexity	Best Use Case

Adjacency List	O(V + E)	Sparse graphs (few edges relative to nodes). Stores only existing connections.

Adjacency Matrix	O(V^2)	Dense graphs (almost all nodes are connected). Allocates memory for every possible pair of nodes.


# Graph Operations: Adjacency List vs Adjacency Matrix

| Category | Operation | Adjacency List | Adjacency Matrix | Explanation |
|---|---|---|---|---|
| **Create** | Add Vertex | O(1) | O(V²) | List adds a new key/bucket. Matrix must allocate and copy a larger (V+1) × (V+1) 2D array. |
| **Create** | Add Edge | O(1) | O(1) | List appends to a vertex's neighbor list. Matrix sets `grid[u][v] = weight`. |
| **Search (Read)** | Find Vertex (by value) | O(V + E) | O(V²) | Searching for a specific data value requires a full graph traversal (BFS/DFS). |
| **Search (Read)** | Check Edge (u, v) | O(deg(V)) | O(1) | Matrix checks `grid[u][v]` directly. List scans u's neighbors (or O(1) if using a Hash Set). |
| **Update** | Update Vertex Data | O(1) | O(1) | Assumes you already have a reference/index to the vertex. Otherwise, add Search time. |
| **Update** | Update Edge Weight | O(deg(V)) | O(1) | Matrix updates `grid[u][v]` instantly. List must first locate the edge in u's neighbor list. |
| **Delete** | Delete Vertex | O(V + E) | O(V²) | Must remove the vertex and scan the graph to delete all edges pointing to it. |
| **Delete** | Delete Edge (u, v) | O(deg(V)) | O(1) | Matrix sets `grid[u][v] = 0`. List must search u's neighbors to find and remove v. |


