package myutils_test

import (
	"sync"
	"testing"

	mypackage "github.com/unbearablelightnessofbeing/myutils"
)

type CacheTestCase[T interface{}] struct {
	Key      string
	Expected T
	Ok       bool
}

func TestCache(t *testing.T) {
	cache := mypackage.NewCache[int]()

	cache.Set("one", 1)
	cache.Set("two", 2)

	cases := []CacheTestCase[int]{
		{
			Key:      "one",
			Expected: 1,
			Ok:       true,
		},
		{
			Key:      "two",
			Expected: 2,
			Ok:       true,
		},
		{
			Key:      "three",
			Expected: 0,
			Ok:       false,
		},
	}

  for _, testCase := range cases {
    v, ok := cache.Get(testCase.Key) 

    if testCase.Expected != v || testCase.Ok != ok {
      t.Fatalf("got wrong value:\n expected: \n\t%v\ngot: \n\t%v", testCase.Expected, v)
    }
  }
}

func TestAsyncCache(t *testing.T) {
  cache := mypackage.NewCache[string]()

  wg := &sync.WaitGroup{}
  wg.Add(2)

  go func() {
    defer wg.Done()
    cache.Set("aboba", "aboba")
    cache.Set("zeliboba", "zeliboba")
  }()

  go func() {
    defer wg.Done()
    cache.Set("amogus", "amogus")
    cache.Set("zelibobus", "zelibobus")
  }()

  wg.Wait()

  cases := []CacheTestCase[string]{
		{
			Key:      "aboba",
			Expected: "aboba",
			Ok:       true,
		},
		{
			Key:      "zeliboba",
			Expected: "zeliboba",
			Ok:       true,
		},
		{
			Key:      "amogus",
			Expected: "amogus",
			Ok:       true,
		},
		{
			Key:      "zelibobus",
			Expected: "zelibobus",
			Ok:       true,
		},
		{
			Key:      "not_exists",
			Expected: "",
			Ok:       false,
		},
  }

  for _, testCase := range cases {
    v, ok := cache.Get(testCase.Key)
    if testCase.Expected != v || testCase.Ok != ok {
      t.Fatalf("got wrong value:\n expected: \n\t%v\ngot: \n\t%v", testCase.Expected, v)
    }
  }
}
