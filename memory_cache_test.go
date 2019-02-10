package main

import (
	"sync"
	"testing"
)

const run = 1000000

func TestMemoryCache_GetConcurrently(t *testing.T) {
	data := map[string]string{
		"test1":  "test1",
		"test2":  "test2",
		"test3":  "test3",
		"test4":  "test4",
		"test5":  "test5",
		"test6":  "test6",
		"test7":  "test7",
		"test8":  "test8",
		"test9":  "test9",
		"test10": "test10",
	}
	var wg sync.WaitGroup
	cache := NewMemoryCache()

	for key, value := range data {
		cache.Set(key, value)
	}

	for i := 1; i <= run; i++ {
		wg.Add(len(data))
		for key, value := range data {
			// Common gopher mistake - using gorutines on loop iterator variable - probably each gorutine will use
			// the same loop value (set after the end of loop but before the gorutines started execution)
			// See https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
			// Can be fixed declaring and using new variables in loop body as they
			// will not be shared between iterations:
			//   key, value := key, value
			// We can fix it also passing (key, value) as parameters to the closure, so they will be placed
			// on the stack for the gorutine, and they will not be shared in loop body:
			// go func(key, value string) {cache.Set(key, value); wg.Done()}(key, value)
			go func() {cache.Set(key, value); wg.Done()}()
		}
	}

	wg.Wait()

	if cache.Len() != uint32(len(data)) {
		t.Fatalf("Cache items number expected to be %d, but it is %d (cache: %v)", len(data), cache.Len(), cache)
	}

	for key, value := range data {
		v, ok := cache.Get(key)

		if v != value && !ok {
			t.Fatalf("Value cache[%s] expected to be %s, but it is %s", key, value, v)
		}
	}
}
