package datetime

import ()

type Time struct {
  Hour, Minute, Second int
}

func (t1 *Time) Plus(t2 * Time) *Date {
  // need to validate t1 and t2?
  hour := t1.Hour + t2.Hour
  minute := t1.Minute + t2.Minute
  second := t1.Second + t2.Second
  day := 0
  
  if second >= 60 {
    second %= 60
    minute++
  }
  
  if minute >= 60 {
    minute %= 60
    hour++
  }
  
  if hour >= 24 {
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