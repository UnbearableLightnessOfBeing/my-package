package myutils_test

import (
	"slices"
	"sync"
	"testing"

	mypackage "github.com/unbearablelightnessofbeing/myutils"
)

type SetTestCase[T comparable] struct {
	Add            []T
  Has T
	ShouldHave     bool
	ExpectedValues []T
}

func checkSetCases[T comparable](t *testing.T, cases []SetTestCase[T]) {
  set := mypackage.NewSet[T]()

	for _, testCase := range cases {
    if len(testCase.Add) != 0 {
      set.Add(testCase.Add...) 
    }

    if set.Has(testCase.Has) != testCase.ShouldHave {
      t.Fatalf("set should have element %v, set: %v", testCase.Has, set)
    }

    res := set.GetElements()
    if len(testCase.ExpectedValues) != len(res) {
      t.Fatalf("values in set do not match:\nexpected:\n\t%v\ngot:\n\t%v", res, testCase.ExpectedValues)
    }

    for _, v := range res {
      if !slices.Contains(testCase.ExpectedValues, v) {
        t.Fatalf("values in set do not match:\nexpected:\n\t%v\ngot:\n\t%v", res, testCase.ExpectedValues)
      }
    }
	}
}

func TestSet(t *testing.T) {
	cases := []SetTestCase[int]{
    {
      Add: []int{},
      Has: 1,
      ShouldHave: false,
      ExpectedValues: []int{},
    },
    {
      Add: []int{1},
      Has: 1,
      ShouldHave: true,
      ExpectedValues: []int{1},
    },
    {
      Add: []int{2},
      Has: 2,
      ShouldHave: true,
      ExpectedValues: []int{1, 2},
    },
    {
      Add: []int{},
      Has: 3,
      ShouldHave: false,
      ExpectedValues: []int{1, 2},
    },
    {
      Add: []int{3, 4},
      Has: 3,
      ShouldHave: true,
      ExpectedValues: []int{1, 2, 3, 4},
    },
	}

  checkSetCases(t, cases)
}

func TestAsyncSet(t *testing.T) {
  set := mypackage.NewSet[string]()

  values := []string{"test_1","test_2","test_3","test_4"}

  wg := &sync.WaitGroup{}
  wg.Add(4)

  for _, v := range values {
    go func() {
      defer wg.Done()
      set.Add(v)
    }()
  }

  wg.Wait()

  result := set.GetElements()

  for _, v := range result {
    if !slices.Contains(values, v) {
      t.Fatalf("values in set do not match:\nexpected:\n\t%v\ngot:\n\t%v", result, values)
    }
  }
}
