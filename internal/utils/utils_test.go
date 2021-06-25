package utils

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		args struct {
			a, b int
		}
		expect int
	}{
		{
			args:   struct{ a, b int }{a: 1, b: 2},
			expect: 3,
		},
	}

	for _, test := range tests {
		if test.expect != Add(test.args.a, test.args.b) {
			t.Errorf("Test Add function error")
		}
	}
}
