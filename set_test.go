package mypackage_test

import (
	"fmt"
	"testing"

	mypackage "github.com/unbearablelightnessofbeing/my-package"
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

    elemsStr := fmt.Sprintf("%v", set.GetElements())
    expectedStr := fmt.Sprintf("%v", testCase.ExpectedValues)
    if elemsStr != expectedStr {
      t.Fatalf("values in set do not match:\nexpected:\n\t%v\ngot:\n\t%v", elemsStr, expectedStr)
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

// func TestAsyncSet(t *testing.T) {
//   set := mypackage.NewSet[string]()
//
//   go func() {
//     set.Add("test_1")
//     set.Add("test_2")
//   }()
// }
