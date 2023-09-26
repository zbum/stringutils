package stringutils

import "testing"

func TestLower(t *testing.T) {

	tests := []struct {
		name string
		args string
		want string
	}{
		{
			"all Upper",
			"TEST",
			"test",
		},
		{
			"upper first character",
			"Test",
			"test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Lower(tt.args); got != tt.want {
				t.Errorf("Lower() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpper(t *testing.T) {

	tests := []struct {
		name string
		args string
		want string
	}{
		{
			"all Upper",
			"TEST",
			"TEST",
		},
		{
			"all Lower",
			"test",
			"TEST",
		},
		{
			"lower first character",
			"tEST",
			"TEST",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Upper(tt.args); got != tt.want {
				t.Errorf("Upper() = %v, want %v", got, tt.want)
			}
		})
	}
}
