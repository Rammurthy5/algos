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

o

(m)min
//{"m": NodePointer for "min"}
min

func (crt *CompressedRadixTrie) Insert(data string) {
	root := crt.root	
	if _, ok := len(root.children)>0; !ok {
		root.children[data[0]] = &radixNode{children: make(map[byte]*radixNode), edgeLabel: data, isEnd: true}
		return
	}
	// iterate over the key, start from the root and the first character of the key. 
	// If the character is not in the children of the root, add it to the children of the root.
	// If the character is in the children of the root, move to the next node and the next character of the key.
	// Repeat the process until the end of the key is reached.
	for i := 0; i < len(data); {
		for k, val := range root.children {
			if k == data[i] {
				for j:=0; j<len(val.edgeLabel); j++ {
					if val.edgeLabel[j] != data[i+j] {
						// split the edge
						splitNode_old := &radixNode{children: make(map[byte]*radixNode), edgeLabel: val.edgeLabel[j:], isEnd: val.isEnd}
						splitNode_new := &radixNode{children: make(map[byte]*radixNode), edgeLabel: val.edgeLabel[j:], isEnd: val.isEnd}
						old_key := val.edgeLabel[j]
						new_key := data[i+j]
						val.edgeLabel = val.edgeLabel[:j]
						val.children[old_key] = splitNode_old
						val.children[new_key] = splitNode_new
						val.isEnd = false
						// TODO: we split the edge into two. we need to add logic to handle split edge that has children already.

				root = val
				
				i++
				break
			}
		}
		root.children[data[i]] = &radixNode{children: make(map[byte]*radixNode), edgeLabel: data[i:], isEnd: true}
	}

		
