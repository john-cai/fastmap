# Fastmap
A concurrent access map with locking per key


## Summary
When the number of items in a map grows very large, having a single mutex to lock access to the whole map is slow. This map has a mutex for every
item in the map, hence providing faster concurrent writes
