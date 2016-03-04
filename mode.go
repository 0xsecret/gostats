package gostats

import (
  "math"
  "sort"
)

//return the most common number from the slice
//if there is not exactly one most common nummber, return DuplicatedError
func Mode(slice []float64) (float64, error) {
  if len(slice) == 0 {
    return math.NaN(), EmptySliceError
  }
  cnt := 1
  maxCnt := 0
  var maxValue float64 = 0
  duplicated := 0
  tempSlice := make([]float64, len(slice))
  copy(tempSlice, slice)
  sort.Float64s(tempSlice)

  for i := 1; i < len(tempSlice); i++ {
    if tempSlice[i] == tempSlice[i-1] {
      cnt++
    } else {
      cnt = 1
    }
    if cnt > maxCnt {
      maxCnt = cnt
      duplicated = 0
      maxValue = tempSlice[i-1]
    } else if cnt == maxCnt {
      duplicated++
    }
  }
  if duplicated != 0 {
    return math.NaN(), DuplicatedError
  } else {
    return maxValue, nil
  }
}
