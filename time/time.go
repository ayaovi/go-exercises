package main

import (
  "testing"
)

type Date struct {
  Day, Month, Year int
  Time *Time
}

type Time struct {
  Hour, Minute, Second int
}

func (c1 *Time) Plus(c2 * Time) *Date {
  // need to validate c1 and c2?
  hour := c1.Hour + c2.Hour
  minute := c1.Minute + c2.Minute
  second := c1.Second + c2.Second
  day := 0
  
  if (second >= 60) {
    second %= 60
    minute++
  }
  
  if (minute >= 60) {
    minute %= 60
    hour++
  }
  
  if (hour >= 24) {
    second %= 24
    day++
  }
  
  return &Date {
    Day: day,
    Month: 0,
    Year: 0,
    Time: &Time {
      Hour: hour,
      Minute: minute,
      Second: second,
    },
  }
}

func (d1 *Date) Plus(d2 *Date) *Date {
  // need to validate d1 and d2?
  return &Date {
    Day: d1.Day + d2.Day,
    Month: d1.Month + d2.Month,
    Year: d1.Year + d2.Year,
  }
}

//func TestTime(t *testing.T) {}

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
  
  if *result != date3 {
    t.Errorf("Date addition was incorrect, got: %d, want: %d.", *result, date3)
  }
}

func main() {}