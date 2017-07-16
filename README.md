# go19-map-bench
Benchmark result for Go 1.9 concurrency map (sync.Map).

```
lnquy@rain:~/workspace/projects/go/src/github.com/lnquy/go19-map-bench$ go test -bench=. -cpu=1
goos: linux
goarch: amd64
pkg: github.com/lnquy/go19-map-bench
BenchmarkNMapIntInt_Write       100000000               16.2 ns/op
BenchmarkNMapIntInt_Read        100000000               15.9 ns/op
BenchmarkNMapStrInt_Write       10000000               242 ns/op
BenchmarkNMapStrInt_Read        10000000               242 ns/op
BenchmarkNMapIntStr_Write       10000000               206 ns/op
BenchmarkNMapIntStr_Read        100000000               17.7 ns/op
BenchmarkNMapStrStr_Write        5000000               251 ns/op
BenchmarkNMapStrStr_Read        10000000               250 ns/op
BenchmarkCMap_Write             10000000               120 ns/op
BenchmarkCMap_Read              30000000                41.3 ns/op
PASS
ok      github.com/lnquy/go19-map-bench 19.477s
lnquy@rain:~/workspace/projects/go/src/github.com/lnquy/go19-map-bench$ go run main.go
Normal map executed in 323.796539ms
Normal map size: 1000000
Sync map 1.9 executed in 726.494685ms
Sync map size: 1000000
```
