package gostats

type statisticsError struct {
  error string
}

func (e statisticsError) Error() string {
  return e.error
}

var (                                           
  EmptySlice = statisticsError{"slice is empty"}
)
