package mypackage_test

import (
	"slices"
	"testing"

	mypackage "github.com/unbearablelightnessofbeing/my-package"
)

func TestSet(t *testing.T) {
  set := mypackage.NewSet[int]()

  set.Add(1)
  set.Add(3)

  if !set.Has(1) {
    t.Fatalf("set is expected to have value '1'")
  }
  if !set.Has(3) {
    t.Fatalf("set is expected to have value '3'")
  }

  if !slices.Equal(set.GetElements(), []int{1,3}) {
    t.Fatalf("set is expected to have values '{1,3}'")
  }

  set.Add(4, 5)
  if !slices.Equal(set.GetElements(), []int{1,3,4,5}) {
    t.Fatalf("set is expected to have values '{1,3,4,5}'\n got: %v", set.GetElements())
  }

  set.Add(5,5,5)
  if !slices.Equal(set.GetElements(), []int{1,3,4,5}) {
    t.Fatalf("set is expected to have values '{1,3,4,5}'")
  }
}

// func TestAsyncSet(t *testing.T) {
//   set := mypackage.NewSet[string]()
// 
//   go func() {
//     set.Add("test_1")
//     set.Add("test_2")
//   }()
// }
