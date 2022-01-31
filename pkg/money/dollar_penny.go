package money

import (
	"errors"
	"strconv"
	"strings"
)

// DollarToPennies takes a dollar amount
// as a string, i.e. 1.00, 55.12 etc and converts it into an int64
func DollarToPennies(amount string) (int64, error) {

	// check if amount can be convert to valid float
	_, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return 0, err
	}

	grps := strings.Split(amount, ".")
	r := grps[0]

	// after "."
	cents := ""

	if len(grps) == 2 {
		cents = grps[1]
	}

	if len(cents) > 2 {
		return 0, errors.New("invalid cents format")
	}

	/*
		r := ""
		if len(grps) == 2 {
			if len(grps[1]) != 2 {
				return 0, errors.New("invalid cents")
			}
			r = grps[1]
			if len(r) > 2 {
				r = r[:2]
			}
		}
	*/

	// pad with 0, this will be 2 0's if there was no .
	for len(cents) < 2 {
		cents += "0"
	}

	r += cents
	return strconv.ParseInt(r, 10, 64)

}

// PenniesToDollar takes a penny amount as an int64 and returns
// a dollar string representation
func PenniesToDollar(amount int64) string {

	r := strconv.FormatInt(amount, 10)
	isNeg := false

	if r[0] == '-' {
		r = r[1:]
		isNeg = true
	}

	// left pad with 0 if we're passed in value < 100
	for len(r) < 3 {
		r = "0" + r
	}

	l := len(r)
	r = r[0:l-2] + "." + r[l-2:]

	// left pad '-' if negnative as checked
	if isNeg {
		r = "-" + r
	}

	return r // Re, return the string value in dollars

}
