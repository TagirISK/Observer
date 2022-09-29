package main

type Friday struct {
}

func (f *Friday) getGuestCount() float64 {
	return 400
}

func (f *Friday) getDescription() string {
	return "Friday"
}

func (f *Friday) getProfit() float64 {
	return 1600$
}
