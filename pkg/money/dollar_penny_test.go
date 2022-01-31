package money

import (
	"fmt"
	"testing"
)

func Test_DollarToPenneis(t *testing.T) {

	tests := []struct {
		name    string
		amount  string
		expect  int64
		wantErr bool
	}{
		{"Normal case", "-15.93", -1593, false},
		{"Without 0 case", "-15.90", -1590, false},
		{"Error case", "364.9384", 0, true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := DollarToPennies(test.amount)

			if (err != nil) != test.wantErr {
				t.Errorf("DollarToPennies() err=%v, wantErr=%v", err, test.wantErr)
			}

			if result != test.expect {
				t.Errorf("DollarToPennies() result=%v, expect=%v", result, test.expect)
			}
		})
	}

}

func Test_PenniesToDollar(t *testing.T) {

}

func main() {

	inputAmt := "-15.93"
	fmt.Printf("input amount is %s dollars\n", inputAmt)

	pennies, err := DollarToPennies(inputAmt)
	if err != nil {
		panic(err)
	}

	fmt.Printf("input amount convert to %d pennies\n", pennies)

	// add 15 cents into last value
	pennies += 15
	dollars := PenniesToDollar(pennies)

	fmt.Printf("add 15 cents, now the value is %s\n", dollars)

}
