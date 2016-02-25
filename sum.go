package gostats

import (
  "math"
)

//adds all the value of a slice
func Sum(slice []float64) (float64, error) {
  if len(slice) == 0 {
    return math.NaN(), EmptySlice
  }
  var sum float64 = 0
  for i := 0; i < len(slice); i++ {
    sum += slice[i]
  }
  return sum, nil
}
