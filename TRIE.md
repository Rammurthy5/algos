A trie is a tree-like data structure designed for efficient storage and retrieval of strings, especially for tasks like prefix-based searching. Unlike binary trees or BSTs, which compare entire values,  tries break strings into individual characters and organize them in a way that exploits common prefixes.
can be implemented as an array or hashmap for node.

## characteristics
Each node represents a single character of a string, root is typically empty. 
A path from the root to a node spells out a prefix or full word. Nodes often have a flag (e.g., isEndOfWord) to mark the end of a valid word.
 Each node can have multiple children—usually up to the size of the alphabet (e.g., 26 for lowercase English letters, or more if you include digits, symbols, etc.). 
 A node could have map or an array to store the children, and a isEndOfWord flag.

## operations
search O(L), where L is the length of the word. Each character is processed in O(1) with a hash map.
insert O(L), where L is the length of the word. Each character is processed in O(1) with a hash map.
prefix-search O(P + N), where P is the prefix length and N is the number of words with that prefix.

## real-time usecases
search on amazon, google. spellcheck, text predictions, code autocomplete on IDEs, IP packet routing, DNS lookups, NLP. 
the tries data are stored typically sharded in-memory datastore (redis, memcache). 
frequently searched items (hot data) are in in-memory datastore and cold data will be in s3, dynamo, cassandra, HFDS.

Stripe (Payment Gateway):
Compressed Trie: For indexing account numbers and transaction metadata.

Bitwise Trie: For validating credit card BINs and routing payments.

Visa (Fraud Detection):
Finite State Trie (Aho-Corasick): For matching transactions against fraud patterns or sanctions lists.

Bitwise Trie: For quick lookups of flagged transaction hashes.

Chase (Banking App):
Standard Trie: For autocomplete of recipient names or transaction memos.

Suffix Tree: For analyzing historical transaction data to detect spending patterns.

SWIFT (Cross-Border Transfers):
Compressed Trie: For routing tables to direct payments between banks.

Bitwise Trie: For matching binary routing codes efficiently.


## types

| Type                 | Space Efficiency | Lookup Speed  | Complexity  | Use Case                          |
|----------------------|-----------------|--------------|------------|-----------------------------------|
| Standard Trie       | Poor            | O(L)         | Simple     | Autocomplete, spell check        |
| Compressed Trie     | Good            | O(L)         | Moderate   | IP routing, space-sensitive      |
| Suffix Trie        | Very Poor       | O(L)         | Simple     | Substring search (basic)         |
| Suffix Tree        | Good            | O(L)         | High       | Bioinformatics, text indexing    |
| Ternary Search Trie | Moderate        | O(L + log k) | Moderate   | Dictionary with memory limits    |
| Bitwise Trie       | Good            | O(k) (bits)  | Simple     | IP routing, binary data          |
| Finite State Trie  | Moderate        | O(n + m)     | High       | Multi-pattern matching           |
| Succinct Trie      | Excellent       | O(L) + overhead | Very High | Massive dictionaries             |
| Burst Trie        | Good            | O(L) + overhead | Moderate  | Databases, file systems          |

Standard/Compressed Tries: Google Search autocomplete (compressed for space).
Bitwise Tries: AWS for IP routing.
Suffix Trees: Bioinformatics at companies like Illumina (via partnerships with tech).
Finite State Tries: Intrusion detection at Meta or Microsoft.

## compression techniques.
path compression (Compressed Tries), node compression (bitmaps), and bit-level compression (Succinct Tries) 
and few more but only these are important.

## trie types to focus
Standard trie (prefix trie)
Suffix Trie
Compressed Trie (RADIX / Patricia)
Bitwise trie
finite state trie (Aho-Corasick)