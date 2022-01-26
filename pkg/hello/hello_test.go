package hello

import "testing"

func Test_Greet_Name(t *testing.T) {
	expect := "Hello! Annie."

	if result := Greet("Annie"); result != expect {
		t.Errorf("Greet(name string) = %s; want %s", result, expect)
	}
}
