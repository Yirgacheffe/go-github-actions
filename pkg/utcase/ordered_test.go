package utcase

import "testing"

func Test_IsPalindrome(t *testing.T) {

	tests := []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"desserts", false},
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}

}
