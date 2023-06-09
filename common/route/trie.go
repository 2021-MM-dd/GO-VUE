package route

type node struct {
	path string
	part string
	sub  []*node
}

func (n *node) insert(path string, parts []string, height int) {
	if len(parts) == height {
		n.path = path
		return
	}
	//查找子节点，如果没有就插入
	part := parts[height]
	sub := n.matchSub(part)
	if sub == nil {
		sub = &node{part: part}
		n.sub = append(n.sub, sub)
		//可以用n.path是否为空来判断是否处于根节点
	}
	//一层一层查询插入
	sub.insert(path, parts, height+1)
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height {
		if n.path == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	subs := n.matchSubs(part)

	for _, sub := range subs {
		result := sub.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}

func (n *node) matchSub(part string) *node {
	for _, sub := range n.sub {
		if sub.part == part {
			return sub
		}
	}
	return nil
}

func (n *node) matchSubs(part string) []*node {
	nodes := make([]*node, 0)
	for _, sub := range n.sub {
		if sub.part == part {
			nodes = append(nodes, sub)
		}
	}
	return nodes
}

func (n *node) travel(list *([]*node)) {
	if n.path != "" {
		*list = append(*list, n)
	}
	for _, sub := range n.sub {
		sub.travel(list)
	}
}
