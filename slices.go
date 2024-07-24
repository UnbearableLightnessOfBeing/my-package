package mypackage

func Filter[T interface{}](s []T, compareFunc func(T) bool) []T {
  var res []T
  for _, v := range s {
    if (compareFunc(v)) {
      res = append(res, v)
    }
  } 
  return res
}

func Map[T interface{}, R interface{}](s []T, callback func(T) R) []R {
  var result []R 
  for _, v := range s {
    result = append(result, callback(v))
  }
  return result
}
