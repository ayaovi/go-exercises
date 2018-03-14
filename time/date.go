package datetime

import ()

type Date struct {
  Day, Month, Year int
  Time *Time
}

func (d1 *Date) Plus(d2 *Date) *Date {
  // need to validate d1 and d2?
  return &Date {
    Day: d1.Day + d2.Day,
    Month: d1.Month + d2.Month,
    Year: d1.Year + d2.Year,
  }
}