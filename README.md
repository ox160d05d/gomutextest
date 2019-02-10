Just an example of common gopher mistake - using gorutines on loop iterator variable; Probably each gorutine will use
the same loop value (set after the end of loop but before the gorutines started execution).
See https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables

Test can be run by `go test . -v -failfast -count=1` Flag `-count` is used here to prevent caching of test results.

Using go version go1.11.5 (linux/amd64) and Intel® Core™ i5-8300H CPU @ 2.30GHz × 8 the test is often failed by unclear 
reason. But it works well with different values of `-cpu=N` flag. Examples:

```
go test . -v -failfast -count=1
=== RUN   TestMemoryCache_GetConcurrently
--- FAIL: TestMemoryCache_GetConcurrently (2.85s)
    memory_cache_test.go:43: Cache items number expected to be 10, but it is 12 (cache: &{{0 0} map[test7:test7 test9:test9 test3:test3 test10:test10 test2:test2 test5:test5 test8:test8 test9u:test9 test1t:test1 test1:test1 test4:test4 test6:test6]})
FAIL
FAIL	.../gomutextest	2.852s
```

The test failed because of some "magic" unexpectable keys like "test2t", "test9u". But if `-cpu` is set 
explicitly to 1 or 2 the test always (?) pass:

```
go test . -v -failfast -count=1 -cpu=1
=== RUN   TestMemoryCache_GetConcurrently
--- PASS: TestMemoryCache_GetConcurrently (5.54s)
PASS
ok  	.../gomutextest	5.545s
```

```
go test . -v -failfast -count=1 -cpu=2
=== RUN   TestMemoryCache_GetConcurrently
--- PASS: TestMemoryCache_GetConcurrently (2.66s)
PASS
ok  	.../gomutextest	2.667s
```

