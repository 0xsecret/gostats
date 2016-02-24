package statistics

import (
  "math"
)

func Max(slice []float64) (float64, error) {
  if len(slice) == 0 {
    return math.NaN(), EmptySlice
  }
  max := slice[0]
  for i := 1; i < len(slice); i++ {
    if slice[i] > max {
      max = slice[i]
    }
  }
  return max, nil
}
