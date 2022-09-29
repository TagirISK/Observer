package main

type Day interface {
	getGuestCount() float64
	getDescription() string
	getProfit() float64
}
