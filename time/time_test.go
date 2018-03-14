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
	
	sub := Time {
		Hour: 0,
		Minute: 0,
		Second: 0,
	}
	
	t.Run("Addition", func(t *testing.T) {
		result := time1.Plus(&time2)
		if !reflect.DeepEqual(*result, add) {
			t.Errorf("Time addition was incorrect, got: %d, want: %d.", *result, add)
		}
	})
	
	t.Run("Subtraction", func(t *testing.T) {
		result := time1.Minus(&time2)
		if !reflect.DeepEqual(*result, sub) {
			t.Errorf("Time subtraction was incorrect, got: %d, want: %d.", *result, add)
		}
	})
}

func BenchmarkTime(b *testing.B) {
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
	
	time1.Plus(&time2)
	time1.Minus(&time2)
}