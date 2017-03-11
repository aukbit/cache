# Cache replacement policies [![Circle CI](https://circleci.com/gh/aukbit/cache.svg?style=svg)](https://circleci.com/gh/aukbit/cache)

Go examples about general cache algorithms

## First In First Out (FIFO)
Using this algorithm the cache behaves in the same way as a Queue.
As in a line or queue at a ticket stand, items are
removed from the data structure in the same order that
they are added.

## Last In First Out (LIFO)
Using this algorithm the cache behaves in the same way as a Stack.
That is, as in a stack of dinner plates, the most recent item added
to the stack is the first item to be removed.

## Least Recently Used (LRU)
Discards the least recently used items first. This algorithm requires
keeping track of what was used when, which is expensive if one wants to
make sure the algorithm always discards the least recently used item.
General implementations of this technique require keeping "age bits"
for cache-lines and track the "Least Recently Used" cache-line based on age-bits.

## Least-Frequently Used (LFU)
Counts how often an item is needed. Those that are used least often are
discarded first. This works very similar to LRU except that instead of
storing the value of how recently a block was accessed, we store the value
of how many times it was accessed. So of course while running an access
sequence we will replace a block which was used least number of times
from our cache.
E.g., if A was used (accessed) 5 times and B was used 3 times
and others C and D were used 10 times each, we will replace B.
If theres a draw we use FIFO.
