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
    hour %= 24
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

func (t1 *Time) Minus(t2 *Time) *Time {
  t := Time {
		Hour: t1.Hour - t2.Hour,
		Minute: t1.Minute - t2.Minute,
		Second: t1.Second - t2.Second,
	}
  r := 0
  
  if t1.Second < t2.Second {
    t.Second = (t1.Second + 60) - t2.Second
    r = -1
  }
  
  if (t1.Minute + r) < t2.Minute {
    t.Minute = (t1.Minute + r + 60) - t2.Minute
    r = -1
  }
  
  t.Hour += r
  
	return &t
}