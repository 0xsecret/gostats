package gostats

import (
  "math"
)

//return the arithmetic mean of a slice
func Mean(slice []float64) (float64, error) {
  return ArithmeticMean(slice)
}

//return the arithmetic mean of a slice
func ArithmeticMean(slice []float64) (float64, error) {
  if len(slice) == 0 {
    return math.NaN(), EmptySliceError
  }
  sum, _ := Sum(slice)
  return sum / float64(len(slice)), nil
}

//return the geometric mean of a slice
func GeometricMean(slice []float64) (float64, error) {
  if len(slice) == 0 {
    return math.NaN(), EmptySliceError
  }
  var p float64 = 1
  for i := 0; i < len(slice); i++ {
    if slice[i] == 0 {
      return math.NaN(), ZeroElementError
    } else if slice[i] < 0 {
      return math.NaN(), NegativeElementError
    }
    p *= slice[i]
  }
  return math.Pow(p, 1/float64(len(slice))), nil
}

//return the harmonic mean of a slice
func HarmonicMean(slice []float64) (float64, error) {
  if len(slice) == 0 {
    return math.NaN(), EmptySliceError
  }
  var s float64 = 0
  for i := 0; i < len(slice); i++ {
    if slice[i] == 0 {
      return math.NaN(), ZeroElementError
    } else if slice[i] < 0 {
      return math.NaN(), NegativeElementError
    }
    s += (1 / slice[i])
  }
  return float64(len(slice)) / s, nil
}
