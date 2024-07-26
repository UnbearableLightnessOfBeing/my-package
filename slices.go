package myutils

// Filter filters out elements in a slice. An element is preserved when
//the callback function called with the element as a parameter returns true.
//Filter returns a new slice.
func Filter[T interface{}](s []T, callback func(T) bool) []T {
  var res []T
  for _, v := range s {
    if (callback(v)) {
      res = append(res, v)
    }
  } 
  return res
}

// Map transforms every element of the passed slice into another type,
//specified as a return value of the callback function. The callback function
//takes an element as a parameter and returns another value calculated based on
// the passed element. Map return a new slice.
func Map[T interface{}, R interface{}](s []T, callback func(T) R) []R {
  var result []R 
  for _, v := range s {
    result = append(result, callback(v))
  }
  return result
}
