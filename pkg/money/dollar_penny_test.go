package money

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	tests := []struct {
		name   string
		amount int64
		expect string
	}{
		{"Normal case", -1593, "-15.93"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := PenniesToDollar(test.amount)
			assert.Equal(t, result, test.expect)
		})
	}
}
