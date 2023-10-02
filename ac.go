package main

type TrieNode struct {
	children map[rune]*TrieNode
	fail     *TrieNode
	isEnd    bool
	pattern  string
}

type ACMatch struct {
	root *TrieNode
}

func NewACMatch() *ACMatch {
	return &ACMatch{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (ac *ACMatch) AddPattern(pattern string) {
	node := ac.root
	for _, char := range pattern {
		if node.children == nil {
			node.children = make(map[rune]*TrieNode)
		}
		if _, exists := node.children[char]; !exists {
			node.children[char] = &TrieNode{}
		}
		node = node.children[char]
	}
	node.isEnd = true
	node.pattern = pattern
}

func (ac *ACMatch) Build() {
	queue := make([]*TrieNode, 0)
	ac.root.fail = nil
	queue = append(queue, ac.root)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for char, child := range current.children {
			queue = append(queue, child)

			failNode := current.fail
			for failNode != nil {
				if _, exists := failNode.children[char]; exists {
					child.fail = failNode.children[char]
					break
				}
				failNode = failNode.fail
			}

			if failNode == nil {
				child.fail = ac.root
			}
		}
	}
}

func (ac *ACMatch) Search(text string) map[int][]string {
	result := make(map[int][]string)
	node := ac.root

	for i, char := range text {
		for node != nil && node.children[char] == nil {
			node = node.fail
		}
		//root fail may be nil
		if node == nil {
			node = ac.root
			continue
		}

		node = node.children[char]
		matchNode := node
		for matchNode != nil {
			if matchNode.isEnd {
				result[i-len(matchNode.pattern)+1] = append(result[i-len(matchNode.pattern)+1], matchNode.pattern)
			}
			matchNode = matchNode.fail
		}
	}
	return result
}

func AC(text string, addresses []string) map[int][]string {
	ac := NewACMatch()

	for _, pattern := range addresses {
		ac.AddPattern(pattern)
	}

	ac.Build()
	return ac.Search(text)
}
