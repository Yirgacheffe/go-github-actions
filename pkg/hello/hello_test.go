package hello

import "testing"

func Test_Greet_Name(t *testing.T) {
	result := Greet("Annie")
	expect := "Hello! Annie."

	if result != expect {
		t.Errorf("Greet(name string) = %s; want %s", result, expect)
	}
}
