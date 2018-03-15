package datetime

import (
  "testing"
	"reflect"
)

func TestTime(t *testing.T) {

	invalidTime := Time {
		Hour: -1,
		Minute: 20,
		Second: 1,
	}
	
	time1 := Time {
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
	
	// test validate resulting in false.
	t.Run("Validate: {-1, 20, 1}", func(t *testing.T) {
		yn, err := invalidTime.Validate()
		expectedErr := &TimeError{ Message: "invalid time {-1 20 1}." }
		if yn {
			t.Errorf("Invalid Time validation should Fail.")
		}
		if !reflect.DeepEqual(err, expectedErr) {
			t.Errorf("Time error message was incorrect, got: %s, want: %s.", err, expectedErr)
		}
	})
	
	// test validate resulting in true.
	t.Run("Validate: {1, 1, 1}", func(t *testing.T) {
		yn, err := time1.Validate()
		expectedErr := &TimeError{ Message: "no errors." }
		if !yn {
			t.Errorf("Valid Time validation should Succeed.")
		}
		if !reflect.DeepEqual(err, expectedErr) {
			t.Errorf("Time error message was incorrect, got: %s, want: %s.", err, expectedErr)
		}
	})
	
	// test addition resulting in zero date.
	t.Run("Addition: {1, 1, 1} + {1, 1, 1}", func(t *testing.T) {
		result, _ := time1.Plus(&time1)
		if !reflect.DeepEqual(*result, add1) {
			t.Errorf("Time addition was incorrect, got: %d, want: %d.", *result, add1)
		}
	})
	
	// test subtraction resulting in zero date.
	t.Run("Subtraction: {1, 1, 1} - {1, 1, 1}", func(t *testing.T) {
		result, _ := time1.Minus(&time1)
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
  
	// test addition resulting in non-zero date.
  t.Run("Addition: {23, 19, 1} + {0, 39, 59}", func(t *testing.T) {
		result, _ := time3.Plus(&time4)
		if !reflect.DeepEqual(*result, add2) {
			t.Errorf("Time addition was incorrect, got: %d, want: %d.", *result, add2)
		}
	})
  
	// test subtraction resulting in non-zero date.
	t.Run("Subtraction: {23, 19, 1} + {0, 39, 59}", func(t *testing.T) {
		result, _ := time3.Minus(&time4)
		if !reflect.DeepEqual(*result, sub2) {
			t.Errorf("Time subtraction was incorrect, got: %d, want: %d.", *result, sub2)
		}
	})
	
	// test addition exception.
	t.Run("Addition: {23, 19, 1} + {-1, 20, 1}", func(t *testing.T) {
		result, err := time3.Plus(&invalidTime)
		expectedErr := &TimeError{Message: "invalid time {-1 20 1}."}
		if !reflect.DeepEqual(err, expectedErr) {
			t.Errorf("Time error message was incorrect, got: %s, want: %s.", err, expectedErr)
		}
		if result != nil {
			t.Errorf("The result of the operation should be nil.")
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