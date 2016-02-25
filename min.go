package gostats

import (
  "math"
)

//return the smallest value in a slice
func Min(slice []float64) (float64, error) {
  if len(slice) == 0 {
    return math.NaN(), EmptySlice
  }
  min := slice[0]
  for i := 1; i < len(slice); i++ {
    if slice[i] < min {
      min = slice[i]
    }
  }
  return min, nil
}
