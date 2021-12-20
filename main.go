package main

import (
	"github.com/davecgh/go-spew/spew"
)

type Change struct {
	Quarters int
	Dimes    int
	Nickels  int
	Pennies  int
}

func main() {
	spew.Dump(moneySorter(59))
}

// divide gives you the quotient
// modulus gives you the remainder
func moneySorter(cents int) Change {
	totalChange := Change{}

	//calculate quarters
	if cents/25 != 0 {
		quarters := cents / 25
		totalChange.Quarters = quarters
		cents = cents - quarters*25
	}
	if cents == 0 {
		return totalChange
	}

	//calculate dimes
	if cents/10 != 0 {
		dimes := cents / 10
		totalChange.Dimes = dimes
		cents = cents - dimes*10
	}
	if cents == 0 {
		return totalChange
	}

	//calculate nickels
	if cents/5 != 0 {
		nickels := cents / 5
		totalChange.Nickels = nickels
		cents = cents - nickels*5
	}

	if cents == 0 {
		return totalChange
	}

	//pennies are whatever is left over
	totalChange.Pennies = cents
	return totalChange

}
