package gostructures

type LRUCache struct {
	capacity uint
	size     uint
	lookup   map[string]*lrucacheNode
	head     *lrucacheNode
	tail     *lrucacheNode
}

type lrucacheNode struct {
	next  *lrucacheNode
	prev  *lrucacheNode
	key   string
	value string
}

func NewLRUCache(capacity uint) LRUCache {
	return LRUCache{
		capacity: capacity,
		lookup:   make(map[string]*lrucacheNode),
	}
}

func (c *LRUCache) Add(key, value string) {
	if node, ok := c.lookup[key]; ok {
		node.value = value
		c.setNodeAsHead(node)
	} else {
		node := lrucacheNode{
			key:   key,
			value: value,
		}
		c.lookup[key] = &node
		c.size++
		c.setNodeAsHead(&node)

		// if we are over capacity, kick out the LRU key
		if c.size > c.capacity {
			c.deleteNode(c.tail)
		}
	}
}

func (c *LRUCache) HasKey(key string) bool {
	_, ok := c.lookup[key]
	return ok
}

func (c *LRUCache) Fetch(key string) (string, bool) {
	if node, ok := c.lookup[key]; ok {
		c.setNodeAsHead(node)
		return node.value, true
	}
	return "", false
}

func (c *LRUCache) Delete(key string) {
	if node, ok := c.lookup[key]; ok {
		c.deleteNode(node)
	}
}

func (c *LRUCache) setNodeAsHead(node *lrucacheNode) {
	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if c.tail == node && node.prev != nil {
		c.tail = node.prev
	}

	if c.tail == nil {
		c.tail = node
	}

	if c.head == nil {
		c.head = node
	}

	if c.head != node {
		c.head, node.next = node, c.head
		node.prev = nil
		node.next.prev = node
	}
}

func (c *LRUCache) deleteNode(node *lrucacheNode) {
	if node.prev != nil {
		node.prev.next = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	}

	if c.tail == node {
		c.tail = node.prev
	}

	if c.head == node {
		c.head = node.next
	}
	delete(c.lookup, node.key)
	c.size--
}
