package hello

import (
	"testing"
)

func Test_Greet_Name(t *testing.T) {
	expect := "Hello! Annie."

	if result := Greet("Annie"); result != expect {
		t.Errorf("Greet(name string) = %s; want %s", result, expect)
	}
}

// gotests generation codes
func TestGreet(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "right case", args: args{name: "ak457"}, want: "Hello! ak457."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Greet(tt.args.name); got != tt.want {
				t.Errorf("Greet() = %v, want %v", got, tt.want)
			}
		})
	}
}
