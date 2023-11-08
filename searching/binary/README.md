# 二分法查找

## 运行

在项目根目录下运行：

```
$ go test -v ./searching/binary/
=== RUN   TestSearch
=== RUN   TestSearch/zero
=== RUN   TestSearch/general-01
=== RUN   TestSearch/general-02
=== RUN   TestSearch/not-exists
--- PASS: TestSearch (0.00s)
    --- PASS: TestSearch/zero (0.00s)
    --- PASS: TestSearch/general-01 (0.00s)
    --- PASS: TestSearch/general-02 (0.00s)
    --- PASS: TestSearch/not-exists (0.00s)
PASS
ok      github.com/cloudfstrife/algs/searching/binary   0.178s
```

## 基准测试

```
$ go test -v -bench . -run=^$ -cpu=2,4,6,8 ./searching/binary/
goos: windows
goarch: amd64
pkg: github.com/cloudfstrife/algs/searching/binary
BenchmarkSearch
BenchmarkSearch-2       34377846                34.8 ns/op
BenchmarkSearch-4       34377550                34.4 ns/op
BenchmarkSearch-6       35433140                33.7 ns/op
BenchmarkSearch-8       35356719                34.0 ns/op
PASS
ok      github.com/cloudfstrife/algs/searching/binary   6.391s
```