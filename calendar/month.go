package calendar

import "time"

func GetPreviousMonth(currentMonth time.Month) time.Month {
	previosMonthMap := map[time.Month]time.Month{
		time.January:   time.December,
		time.February:  time.January,
		time.March:     time.February,
		time.April:     time.March,
		time.May:       time.April,
		time.June:      time.May,
		time.July:      time.June,
		time.August:    time.July,
		time.September: time.August,
		time.October:   time.September,
		time.November:  time.October,
		time.December:  time.November,
	}

	return previosMonthMap[currentMonth]
}

func GetMonthInRussian(month time.Month) string {
	monthMap := map[time.Month]string{
		time.January:   "Январь",
		time.February:  "Февраль",
		time.March:     "Март",
		time.April:     "Апрель",
		time.May:       "Май",
		time.June:      "Июнь",
		time.July:      "Июль",
		time.August:    "Август",
		time.September: "Сентябрь",
		time.October:   "Октябрь",
		time.November:  "Ноябрь",
		time.December:  "Декабрь",
	}

	return monthMap[month]
}
