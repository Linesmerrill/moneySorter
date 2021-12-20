package main

import "testing"

func TestWorkingExampleWith59(t *testing.T) {
	testChange := moneySorter(59)

	expectedChange := Change{
		Quarters: 2,
		Dimes:    0,
		Nickels:  1,
		Pennies:  4,
	}

	if testChange != expectedChange {
		t.Logf("Test results:\nexpected %v\ngot %v", expectedChange, testChange)
		t.Fail()
	}
}

func TestFailureExampleWith59(t *testing.T) {
	testChange := moneySorter(59)

	expectedChange := Change{
		Quarters: 2,
		Dimes:    0,
		Nickels:  1,
		Pennies:  3,
	}

	if testChange == expectedChange {
		t.Logf("Test results:\nexpected %v\ngot %v", expectedChange, testChange)
		t.Fail()
	}
}

func TestDimes(t *testing.T) {
	testChange := moneySorter(21)

	expectedChange := Change{
		Quarters: 0,
		Dimes:    2,
		Nickels:  0,
		Pennies:  1,
	}

	if testChange != expectedChange {
		t.Logf("Test results:\nexpected %v\ngot %v", expectedChange, testChange)
		t.Fail()
	}
}

func TestExactQuarter(t *testing.T) {
	testChange := moneySorter(25)

	expectedChange := Change{
		Quarters: 1,
		Dimes:    0,
		Nickels:  0,
		Pennies:  0,
	}

	if testChange != expectedChange {
		t.Logf("Test results:\nexpected %v\ngot %v", expectedChange, testChange)
		t.Fail()
	}
}

func TestExactDime(t *testing.T) {
	testChange := moneySorter(10)

	expectedChange := Change{
		Quarters: 0,
		Dimes:    1,
		Nickels:  0,
		Pennies:  0,
	}

	if testChange != expectedChange {
		t.Logf("Test results:\nexpected %v\ngot %v", expectedChange, testChange)
		t.Fail()
	}
}

func TestExactNickel(t *testing.T) {
	testChange := moneySorter(5)

	expectedChange := Change{
		Quarters: 0,
		Dimes:    0,
		Nickels:  1,
		Pennies:  0,
	}

	if testChange != expectedChange {
		t.Logf("Test results:\nexpected %v\ngot %v", expectedChange, testChange)
		t.Fail()
	}
}
