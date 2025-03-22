package src

type radixNode struct {
	children  map[byte]*radixNode
	isEnd     bool
	edgeLabel string
}

type CompressedRadixTrie struct {
	root *radixNode
}

func NewCompressedRadixTrie() *CompressedRadixTrie {
	return &CompressedRadixTrie{root: &radixNode{children: make(map[byte]*radixNode)}}
}

func (crt *CompressedRadixTrie) Insert(data string) {
	current := crt.root
	if len(current.children) == 0 {
		current.children[data[0]] = &radixNode{children: make(map[byte]*radixNode), edgeLabel: data, isEnd: true}
		return
	}

	i := 0
	for i < len(data) {
		if child, exists := current.children[data[i]]; exists {
			j := 0
			for ; j < len(child.edgeLabel) && i+j < len(data) && child.edgeLabel[j] == data[i+j]; j++ {
			}
			if j == len(child.edgeLabel) {
				current = child
				i += j
				if i == len(data) {
					child.isEnd = true
					return
				}
			} else {
				// Split the edge
				common := child.edgeLabel[:j]
				remainingOld := child.edgeLabel[j:]
				remainingNew := data[i+j:]

				newChild := &radixNode{children: make(map[byte]*radixNode), edgeLabel: remainingNew, isEnd: true}
				oldChild := &radixNode{children: child.children, edgeLabel: remainingOld, isEnd: child.isEnd}
				current.children[data[i]] = &radixNode{
					children:  map[byte]*radixNode{remainingOld[0]: oldChild, remainingNew[0]: newChild},
					edgeLabel: common,
					isEnd:     j == len(data)-i,
				}
				return
			}
		} else {
			current.children[data[i]] = &radixNode{children: make(map[byte]*radixNode), edgeLabel: data[i:], isEnd: true}
			return
		}
	}
}

func (crt *CompressedRadixTrie) Search(data string) bool {
	if data == "" {
		return crt.root.isEnd // Only true if empty string was inserted
	}
	current := crt.root
	i := 0
	for i < len(data) {
		if child, exists := current.children[data[i]]; exists {
			for j := 0; j < len(child.edgeLabel); j++ {
				if i+j >= len(data) || child.edgeLabel[j] != data[i+j] {
					return false
				}
			}
			i += len(child.edgeLabel)
			current = child
			if i == len(data) {
				return child.isEnd
			}
		} else {
			return false
		}
	}
	return false
}

func (crt *CompressedRadixTrie) Delete(data string) bool {
	if data == "" {
		if crt.root.isEnd {
			crt.root.isEnd = false
			return true
		}
		return false
	}

	current := crt.root
	var parent *radixNode // Track parent for deletion and merging
	var parentKey byte    // Key to the current node in parent's children
	i := 0

	for i < len(data) {
		if child, exists := current.children[data[i]]; exists {
			parent = current
			parentKey = data[i]
			current = child

			// Traverse the edge label
			j := 0
			for ; j < len(child.edgeLabel) && i+j < len(data); j++ {
				if child.edgeLabel[j] != data[i+j] {
					return false // Mismatch, string not in trie
				}
			}

			// Check if we reached the end of the input
			if i+j == len(data) {
				if !child.isEnd {
					return false // String exists as prefix but not as a word
				}
				if j < len(child.edgeLabel)-1 {
					// Case 1: partial match of edge label but full match of input. "mad" and "madness"
					return false
				}
				child.isEnd = false // Mark as no longer a word

				// Case 1: Leaf node - remove it from parent
				if len(child.children) == 0 && parent != nil {
					delete(parent.children, parentKey)
					// Check if parent needs merging after deletion
					if len(parent.children) == 1 && !parent.isEnd && parent != crt.root {
						for key, grandChild := range parent.children {
							newEdgeLabel := parent.edgeLabel + grandChild.edgeLabel
							parent.children[key] = &radixNode{
								children:  grandChild.children,
								edgeLabel: newEdgeLabel,
								isEnd:     grandChild.isEnd,
							}
							break // Only one child, so break after first iteration
						}
					}
					return true
				}

				// Case 2: Node has one child - merge with it
				if len(child.children) == 1 && !child.isEnd {
					for _, grandChild := range child.children {
						newEdgeLabel := child.edgeLabel + grandChild.edgeLabel
						if parent != nil {
							delete(parent.children, parentKey) // Remove old reference
							parent.children[child.edgeLabel[0]] = &radixNode{
								children:  grandChild.children,
								edgeLabel: newEdgeLabel,
								isEnd:     grandChild.isEnd,
							}
						} else {
							// Root case: replace root directly
							crt.root = &radixNode{
								children:  map[byte]*radixNode{newEdgeLabel[0]: grandChild},
								edgeLabel: newEdgeLabel,
								isEnd:     grandChild.isEnd,
							}
						}
						break
					}
					return true
				}
				// Case 3: Node has multiple children or is still an end node - no further action
				return true
			}

			// Move to next node
			i += len(child.edgeLabel)
		} else {
			return false // No matching child
		}
	}
	return false // String not fully matched
}

func CompressedRadixTrieRun() {
	crt := NewCompressedRadixTrie()
	crt.Insert("Mad")
	crt.Insert("Mam")
	crt.Insert("Rad")
	crt.Delete("M")
	crt.Delete("Mad")
	crt.Search("Ram")
	crt.Search("Rad")
}
