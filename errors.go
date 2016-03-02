package gostats

type gostatsError struct {
  error string
}

func (e gostatsError) Error() string {
  return e.error
}

var (
  EmptySliceError      = gostatsError{"slice is empty"}                        
  ZeroElementError     = gostatsError{"slice must not contain zero(0)"}        
  NegativeElementError = gostatsError{"slice must not contain negative number"}
)
