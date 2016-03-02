package gostats

type gostatsError struct {
  error string
}

func (e gostatsError) Error() string {
  return e.error
}

var (
  EmptySliceError = gostatsError{"slice is empty"}
)
