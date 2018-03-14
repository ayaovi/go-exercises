package datetime

import (
	"testing"
	"reflect"
)

func TestDate(t *testing.T) {
  date1 := Date {
    Day: 1,
    Month: 1,
    Year: 1,
    Time: &Time {
      Hour: 1,
      Minute: 1,
      Second: 1,
    },
  }
  
  date2 := Date {
    Day: 1,
    Month: 1,
    Year: 1,
    Time: &Time {
      Hour: 1,
      Minute: 1,
      Second: 1,
    },
  }
  
  date3 := Date {
    Day: 2,
    Month: 2,
    Year: 2,
    Time: nil,
  }
  
  result := date1.Plus(&date2)
  
  if !reflect.DeepEqual(*result, date3) {
    t.Errorf("Date addition was incorrect, got: %d, want: %d.", *result, date3)
  }
}

func BenchmarkDate(b *testing.B) {
	date1 := Date {
    Day: 1,
    Month: 1,
    Year: 1,
    Time: &Time {
      Hour: 1,
      Minute: 1,
      Second: 1,
    },
  }
	
	date2 := Date {
    Day: 1,
    Month: 1,
    Year: 1,
    Time: &Time {
      Hour: 1,
      Minute: 1,
      Second: 1,
    },
  }
	
	date1.Plus(&date2)
}