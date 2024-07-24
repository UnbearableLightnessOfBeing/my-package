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
