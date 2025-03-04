A trie is a tree-like data structure designed for efficient storage and retrieval of strings, especially for tasks like prefix-based searching. Unlike binary trees or BSTs, which compare entire values,  tries break strings into individual characters and organize them in a way that exploits common prefixes.

## characteristics
Each node represents a single character of a string, root is typically empty. A path from the root to a node spells out a prefix or full word. Nodes often have a flag (e.g., isEndOfWord) to mark the end of a valid word. Each node can have multiple children—usually up to the size of the alphabet (e.g., 26 for lowercase English letters, or more if you include digits, symbols, etc.). A node could have map or an array
to store the children, and a isEndOfWord flag.

## operations
search O(L), where L is the length of the word. Each character is processed in O(1) with a hash map.
insert O(L), where L is the length of the word. Each character is processed in O(1) with a hash map.
prefix-search O(P + N), where P is the prefix length and N is the number of words with that prefix.

## real-time usecases
word predictor

