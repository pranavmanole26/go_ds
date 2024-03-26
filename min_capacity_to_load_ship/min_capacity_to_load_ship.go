package mincapacitytoloadship

import "fmt"

func Find(loads []int, maxDays int) int {
	l := loads[0]
	r := 0
	for _, load := range loads {
		r += load
	}
	cap := r
	for {
		if l > r {
			fmt.Println(cap)
			return cap
		}

		m := (l + r) / 2
		days := getDaysForCapacity(loads, m)
		if days <= maxDays {
			r = m - 1
			cap = m
		} else {
			l = m + 1
		}

	}

}

func getDaysForCapacity(loads []int, cap int) int {
	days := 1
	loadSum := 0

	for _, load := range loads {

		if loadSum+load > cap {
			days += 1
			loadSum = load
		} else {
			loadSum += load
		}
	}
	return days
}
