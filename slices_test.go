package mypackage_test

import (
	"testing"

	mypackage "github.com/unbearablelightnessofbeing/my-package"
)

type FilterTestCase[T interface{}] struct {
	input       []T
	compareFunc func(T) bool
	expected    []T
}

func checkCases[T comparable](t *testing.T, cases []FilterTestCase[T]) {

	for _, testCase := range cases {
	  res := mypackage.Filter(testCase.input, testCase.compareFunc)
		if len(res) != len(testCase.expected) {
			t.Fatalf(
				"length doesn't match:\nexpected: \n\t%v\ngot: \n\t%v",
				testCase.expected,
				res,
			)
		}

		for i, v := range res {
			if v != testCase.expected[i] {
				t.Fatalf(
					"elements do not match:\nexpected: \n\t%v\ngot: \n\t%v",
					testCase.expected,
					res,
				)
			}
		}
	}
}

func TestFilter(t *testing.T) {
	// test int
	casesInt := []FilterTestCase[int]{
		{
			input:  []int{1, 2, 3, 4, 5, 6},
      compareFunc: func(val int) bool {
        return val%2 == 0
      },
			expected: []int{2, 4, 6},
		},
		{
			input:  []int{111, 22, 334, 123, 342, 56, 100},
      compareFunc: func(val int) bool {
        return val >= 100
      },
			expected: []int{111, 334, 123, 342, 100},
		},
	}

	checkCases(t, casesInt)

  // test string
  casesStr := []FilterTestCase[string] {
    {
      input: []string{"aboba", "amogus"},
      compareFunc: func(s string) bool {
        return s != "aboba"
      },
      expected: []string{"amogus"},
    },
    {
      input: []string{"test_1", "test_2", "test_3"},
      compareFunc: func(s string) bool {
        bytes := []byte(s)
        if (bytes[len(bytes)-1] == byte('2')) {
          return false
        }
        return true
      },
      expected: []string{"test_1", "test_3"},
    },
  }

  checkCases(t, casesStr)
}
