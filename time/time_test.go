package datetime

import (
  "testing"
	"reflect"
)

func TestTime(t *testing.T) {
	time1 := Time {
		Hour: 1,
		Minute: 1,
		Second: 1,
	}
	
	time2 := Time {
		Hour: 1,
		Minute: 1,
		Second: 1,
	}
	
	add := Date {
		Day: 0,
    Month: 0,
    Year: 0,
    Time: &Time {
      Hour: 2,
      Minute: 2,
      Second: 2,
    },
	}
	
	result := time1.Plus(&time2)
	
	if !reflect.DeepEqual(*result, add) {
    t.Errorf("Time addition was incorrect, got: %d, want: %d.", *result, add)
  }
}