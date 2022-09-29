package main

import "fmt"

func main() {
	day1 := &WednesDay{}

	day1WeekDays := &WeekDays{
		day: day1,
	}

	day1WeekDaysWeekends := &Weekends{
		day: day1WeekDays,
	}

	day1WeekDaysWeekendsChristmas := &Weekends{
		day: day1WeekDaysWeekends,
	}

	fmt.Printf("%v \n Guest Count - %v \n Profit - %v",
		day1WeekDaysWeekendsHolidaysChristmasChineseNewYear.getDescription(),
		day1WeekDaysWeekendsHolidaysChristmasChineseNewYear.getGuestCount(),
		day1WeekDaysWeekendsHolidaysChristmasChineseNewYear.getProfit())

		day2 := &Friday{}

		day2ChineseNewYear := &ChineseNewYear{
			day: day2,
	}

	day2ChineseNewYear := &Weekends{
		day: day2ChineseNewYear,
	}

	fmt.Println("\n-----------------------")
	fmt.Printf("%v \n Guest Count - %v \n Profit - %v",
	day2ChineseNewYearWeekends.getDescription(),
	day2ChineseNewYearWeekends.getGuestCount(),
	day2ChineseNewYearWeekends.getProfit())
}
