#Big O [![Circle CI](https://circleci.com/bb/_paulo/bigo.svg?style=svg)](https://circleci.com/bb/_paulo/bigo)

Go examples to better understand Big O notation

```
$ GOMAXPROCS=1 go test . -v -bench=. -benchmem
BenchmarkIsFirstZero10                  2000000000               0.37 ns/op            0 B/op          0 allocs/op
BenchmarkIsFirstZero1M                  2000000000               0.37 ns/op            0 B/op          0 allocs/op
BenchmarkContains10                     200000000                7.25 ns/op            0 B/op          0 allocs/op
BenchmarkContains1M                         3000            508857 ns/op               0 B/op          0 allocs/op
BenchmarkContainsDuplicates10           20000000                57.1 ns/op             0 B/op          0 allocs/op
BenchmarkContainsDuplicates1000             3000            564392 ns/op               0 B/op          0 allocs/op
BenchmarkFibonacci10                     5000000               320 ns/op               0 B/op          0 allocs/op
BenchmarkFibonacci20                       50000             40093 ns/op               0 B/op          0 allocs/op
```
