package main

type HotelEleonVisitorTurnover interface {
	getGuestCount() float64
	getDescription() string
	getProfit() float64
}

type WeekDays struct {
	Day Day
}

func (s *WeekDays) getGuestCount() float64 {
	dayGuestCount := s.day.getGuestCount()
	return DayGuestCount
}

func (s *WeekDays) getDescription() string {
	dayDescription := s.day.getDescription()
	return DayDescription + "\n Average Guests on WeekDays"
}

func (s *WeekDays) getProfit() float64 {
	dayProfit := s.day.getProfit()
	return dayProfit + 60$
}

type Weekends struct {
	day Day
}

func (f *Weekends) getGuestCount() float64 {
	dayGuestCount := f.day.getGuestCount()
	return dayGuestCount + 300
}

func (f *Weekends) getDescription() string {
	dayDescription := f.day.getDescription()
	return dayDescription + "\nAverage Guests on Weekends"
}

func (f *Weekends) getProfit() float64 {
	dayProfit := f.day.getProfit()
	return dayProfit + 200$
}

type Holidays struct {
	day Day
}

func (f *Holidays) getGuestCount() float64 {
	dayGuestCount := f.day.getGuestCount()
	return dayGuestCount
}

func (o *Holidays) getDescription() string {
	dayDescription := o.day.getDescription()
	return dayDescription + "\nAverage Guests on Holidays"
}

func (o *Holidays) getProfit() float64 {
	dayProfit := o.day.getProfit()
	return dayProfit + 150;
}

type Christmas struct {
	day Day
}

func (f *Christmas) getGuestCount() float64 {
	dayGuestCount := f.day.getGuestCount()
	return dayGuestCount + 200;
}

func (r *Christmas) getDescription() string {
	dayDescription := r.day.getDescription()
	return dayDescription + "\nAverage Guests on Christmas"
}

func (r *Christmas) getProfit() float64 {
	dayProfit := r.day.getProfit()
	return dayProfit + 600$
}

type ChineseNewYear struct {
	day Day
}

func (s *ChineseNewYear) getGuestCount() float64 {
	dayGuestCount := s.day.getGuestCount()
	return dayGuestCount
}

func (s *ChineseNewYear) getDescription() string {
	dayDescription := s.day.getDescription()
	return weaponDescription + "\nAverage Guests on ChineseNewYear"
}

func (s *ChineseNewYear) getProfit() float64 {
	dayProfit := s.day.getProfit()
	return dayProfit + 1000$
}
