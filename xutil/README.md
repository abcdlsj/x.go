# X Util

- [x] [clone](##clone)
- [x] [batch](##batch)


## clone
*for clone from a to b*

*thanks to chatGPT*

```shell
» go test -bench=.
goos: darwin
goarch: amd64
pkg: play
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkCloneStruct-12         17105265                67.86 ns/op
BenchmarkCloneStr-12            15476632                79.17 ns/op
BenchmarkCloneInt-12            13610522                75.16 ns/op
PASS
ok      play    5.070s
```

## batch

just a batch loop abstract.

```go
...
	batch := NewBatch(0, 100, 12) // start, end, step
	for batch.HasNext() {
		t.Log(batch.Next())
	}
...
```