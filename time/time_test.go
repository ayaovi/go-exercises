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
	
	add1 := Date {
		Day: 0,
    Month: 0,
    Year: 0,
    Time: &Time {
      Hour: 2,
      Minute: 2,
      Second: 2,
    },
	}
	
	sub1 := Time {
		Hour: 0,
		Minute: 0,
		Second: 0,
	}
	
	t.Run("Addition: {1, 1, 1} + {1, 1, 1}", func(t *testing.T) {
		result := time1.Plus(&time2)
		if !reflect.DeepEqual(*result, add1) {
			t.Errorf("Time addition was incorrect, got: %d, want: %d.", *result, add1)
		}
	})
	
	t.Run("Subtraction: {1, 1, 1} - {1, 1, 1}", func(t *testing.T) {
		result := time1.Minus(&time2)
		if !reflect.DeepEqual(*result, sub1) {
			t.Errorf("Time subtraction was incorrect, got: %d, want: %d.", *result, sub1)
		}
	})
  
  time3 := Time {
		Hour: 23,
		Minute: 20,
		Second: 1,
	}
	
	time4 := Time {
		Hour: 0,
		Minute: 39,
		Second: 59,
	}
	
	add2 := Date {
		Day: 1,
    Month: 0,
    Year: 0,
    Time: &Time {
      Hour: 0,
      Minute: 0,
      Second: 0,
    },
	}
	
	sub2 := Time {
		Hour: 22,
		Minute: 40,
		Second: 2,
	}
  
  t.Run("Addition: {23, 19, 1} + {0, 39, 59}", func(t *testing.T) {
		result := time3.Plus(&time4)
		if !reflect.DeepEqual(*result, add2) {
			t.Errorf("Time addition was incorrect, got: %d, want: %d.", *result, add2)
		}
	})
  
	t.Run("Subtraction: {23, 19, 1} + {0, 39, 59}", func(t *testing.T) {
		result := time3.Minus(&time4)
		if !reflect.DeepEqual(*result, sub2) {
			t.Errorf("Time subtraction was incorrect, got: %d, want: %d.", *result, sub2)
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