package lru

type Item struct {
	Key  string
	Data interface{}
	pre  *Item
	next *Item
}

type Cache struct {
	first *Item
	last  *Item
	// hash map with keys = items, values = location in linked list.
	hash map[string]*Item
	n    int
	c    int
}

func New(capacity int) *Cache {
	return &Cache{
		c:    capacity,
		hash: make(map[string]*Item),
	}
}

// Access operation inserts the item onto the data structure if it’s
// not already present.
func (c *Cache) Access(key string, data interface{}) *Item {
	i := c.get(key)
	if i == nil {
		i = c.set(key, data)
	}
	return i
}

// get operation shifts the item to the first position on Cache is it’s already present.
// Otherwise returns nil.
func (c *Cache) get(key string) *Item {
	// if not present returns nil
	i, ok := c.hash[key]
	if !ok {
		return nil
	}
	if c.Size() == 1 || c.first == i {
		return i
	}
	// delete item from the linked list
	next, _ := c.delete(i)
	if c.last == i {
		c.last = next
	}
	// reinsert item at the beginning of the linked list
	c.setFirst(i)
	//
	return i
}

// set operation inserts the item onto the data structure and removes the least
// frequently used if cache is full
func (c *Cache) set(key string, data interface{}) *Item {
	i := &Item{
		Key:  key,
		Data: data,
	}
	if c.IsEmpty() {
		c.first = i
		c.last = i
		c.hash[key] = i
		c.n++
		return i
	}
	if c.IsFull() {
		// remove the least frequently used from the linked list which is the on in last
		next, _ := c.delete(c.last)
		delete(c.hash, c.last.Key)
		c.last = next
		c.n--
	}
	c.setFirst(i)
	c.hash[key] = i
	c.n++
	return i
}

// setFirst adds item has first in cache
func (c *Cache) setFirst(i *Item) {
	c.first.next = i
	i.pre = c.first
	c.first = i
}

// Delete deletes the item from linked list
func (c *Cache) delete(i *Item) (next *Item, pre *Item) {
	// link next item to the previous item
	if i.pre != nil {
		i.pre.next = i.next
	}
	// link previous item to the next item
	if i.next != nil {
		i.next.pre = i.pre
	}
	next = i.next
	pre = i.pre
	// remove links
	i.next = nil
	i.pre = nil
	return next, pre
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
