package myutils_test

import (
	"strings"
	"testing"

	mypackage "github.com/unbearablelightnessofbeing/myutils"
)

type FilterTestCase[T interface{}] struct {
	input       []T
	compareFunc func(T) bool
	expected    []T
}

func checkFilterCases[T comparable](t *testing.T, cases []FilterTestCase[T]) {
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
			input: []int{},
			compareFunc: func(val int) bool {
				return true
			},
			expected: []int{},
		},
		{
			input: []int{1, 2, 3, 4, 5, 6},
			compareFunc: func(val int) bool {
				return val%2 == 0
			},
			expected: []int{2, 4, 6},
		},
		{
			input: []int{111, 22, 334, 123, 342, 56, 100},
			compareFunc: func(val int) bool {
				return val >= 100
			},
			expected: []int{111, 334, 123, 342, 100},
		},
	}

	checkFilterCases(t, casesInt)

	// test string
	casesStr := []FilterTestCase[string]{
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
				if bytes[len(bytes)-1] == byte('2') {
					return false
				}
				return true
			},
			expected: []string{"test_1", "test_3"},
		},
	}

	checkFilterCases(t, casesStr)
}

type MapTestCase[T interface{}, R interface{}] struct {
	input    []T
	callback func(T) R
	expected []R
}

func checkMapCases[T interface{}, R comparable](t *testing.T, cases []MapTestCase[T, R]) {
	for _, testCase := range cases {
		res := mypackage.Map(testCase.input, testCase.callback)
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

func TestMap(t *testing.T) {
	casesStrToInt := []MapTestCase[string, int]{
		{
			input: []string{"aboba", "zeliboba"},
			callback: func(s string) int {
				return len(s)
			},
			expected: []int{5, 8},
		},
		{
			input: []string{},
			callback: func(s string) int {
				return 123
			},
			expected: []int{},
		},
	}

	checkMapCases(t, casesStrToInt)

	casesStrToBool := []MapTestCase[string, bool]{
		{
			input: []string{"aboba", "zeliboba", "amogus"},
			callback: func(s string) bool {
				return strings.Contains(s, "boba")
			},
			expected: []bool{true, true, false},
		},
	}

	checkMapCases(t, casesStrToBool)

	type TestStruct struct {
		Name string
		Age  uint
	}

	type TestStructModified struct {
		Name    string
		Age     uint
		IsAdult bool
	}

  casesModifiedStruct := []MapTestCase[TestStruct, TestStructModified]{
    {
      input: []TestStruct{{Name: "Jack", Age: 17}, {Name: "Gabriel", Age: 20}},
      callback: func(s TestStruct) TestStructModified {
        return TestStructModified{
          Name: s.Name,
          Age: s.Age,
          IsAdult: s.Age >= 18,
        }
      },
      expected: []TestStructModified{
        { Name: "Jack", Age: 17, IsAdult: false }, 
        { Name: "Gabriel", Age: 20, IsAdult: true },
      },
    },
  }

  checkMapCases(t, casesModifiedStruct)
}
