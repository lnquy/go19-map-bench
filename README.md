# go19-map-bench
Benchmark for Go 1.9 concurrency map (sync.Map)

```
lnquy@lnquy:~/workspace/projects/go/src/github.com/lnquy/go19-map-bench$ go test -bench=. -cpu=1
goos: linux
goarch: amd64
pkg: github.com/lnquy/go19-map-bench
BenchmarkNMapIntInt_Read        1000000000               2.81 ns/op
BenchmarkNMapIntInt_Write       200000000                8.19 ns/op
BenchmarkNMapStrInt_Read        10000000               146 ns/op
BenchmarkNMapStrInt_Write       10000000               155 ns/op
BenchmarkNMapIntStr_Read        500000000                3.20 ns/op
BenchmarkNMapIntStr_Write       10000000               153 ns/op
BenchmarkNMapStrStr_Read        10000000               147 ns/op
BenchmarkNMapStrStr_Write        5000000               280 ns/op
BenchmarkCMap_Read              200000000                9.13 ns/op
BenchmarkCMap_Write             20000000                93.8 ns/op
PASS
ok      github.com/lnquy/go19-map-bench 20.588s
```
