package lru

// Item generic type of an item in the Cache
type Item interface{}

// Node is a single representation of a data structure in the Cache
type Node struct {
	key  string
	item Item
	pre  *Node
	next *Node
}

type Cache struct {
	first *Node
	last  *Node
	// hash map with keys = nodes, values = location in linked list.
	hash map[string]*Node
	n    int
	c    int
}

func New(capacity int) *Cache {
	return &Cache{
		c:    capacity,
		hash: make(map[string]*Node),
	}
}

// Access operation inserts the item onto the Cache if it’s not already present.
func (c *Cache) Access(key string, i Item) {
	item := c.get(key)
	if item == nil {
		c.set(key, i)
	}
}

// Remove operation deletes and returns the item that was least
// recently accessed
func (c *Cache) Remove() Item {
	// last position in the cache cointains the item least recently used
	next := c.last.next
	c.detach(c.last)
	delete(c.hash, c.last.key)
	c.last = next
	c.n--
	return c.last.item
}

// get operation shifts the item to the first position on Cache if it’s already present.
// Otherwise returns nil.
func (c *Cache) get(key string) Item {
	// if not present returns nil
	n, ok := c.hash[key]
	if !ok {
		return nil
	}
	if c.Size() == 1 || c.first == n {
		return n.item
	}
	// detach item from the linked list
	next := n.next
	c.detach(n)
	if c.last == n {
		c.last = next
	}
	// reinsert node at the beginning of the linked list
	c.attach(n)
	return n.item
}

// set operation inserts the item onto the Cache and removes the
// least recently used if cache is full. The least recently used item is the one
// in the last position.
func (c *Cache) set(key string, i Item) {
	n := &Node{
		key:  key,
		item: i,
	}
	if c.IsEmpty() {
		c.first = n
		c.last = n
		c.hash[key] = n
		c.n++
		return
	}
	if c.IsFull() {
		c.Remove()
	}
	c.attach(n)
	c.hash[key] = n
	c.n++
}

// attach operation adds node has first in cache
func (c *Cache) attach(n *Node) {
	c.first.next = n
	n.pre = c.first
	c.first = n
}

// detach a node from the linked list
func (*Cache) detach(n *Node) {
	// link next node to the previous node
	if n.pre != nil {
		n.pre.next = n.next
	}
	// link previous node to the next node
	if n.next != nil {
		n.next.pre = n.pre
	}
	// remove links
	n.next = nil
	n.pre = nil
}

func (c *Cache) IsEmpty() bool {
	return c.first == nil && c.last == nil
}

func (c *Cache) IsFull() bool {
	return c.n == c.c
}

func (c *Cache) Size() int {
	return c.n
}
