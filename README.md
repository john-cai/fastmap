# Fastmap
A concurrent access map with locking per key


## Summary
When the number of items in a map grows very large, having a single mutex to lock access to the whole map is slow. This map has a mutex for every
item in the map, hence providing faster concurrent writes


## Usage

```
m := fastmap.New()

```

Set
```
m.Set("key","value")
```

Get
```
val, ok := m.Get("key")
```

Delete
```
m.Del("key")
```


## Benchmark results
```
BenchmarkLockedValMap-8  5000000               924 ns/op             145 B/op          1 allocs/op
ok      github.com/john-cai/fastmap     5.189s
```
