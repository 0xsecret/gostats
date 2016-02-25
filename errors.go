package gostats

type gostatsError struct {
  error string
}

func (e gostatsError) Error() string {
  return e.error
}

var (                                           
  EmptySlice = gostatsError{"slice is empty"}
)
