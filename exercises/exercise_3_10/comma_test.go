package exercise310

import "testing"

func TestComma(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "Given I have 3 character long string",
			input:  "123",
			output: "123",
		},
		{
			name:   "Given I have 4 character long string",
			input:  "1234",
			output: "1,234",
		},
		{
			name:   "Given I have 2 character long string",
			input:  "12",
			output: "12",
		},
		{
			name:   "Given I have 9 character long string",
			input:  "123456789",
			output: "123,456,789",
		},
		{
			name:   "Given I have 11 character long string",
			input:  "12345678901",
			output: "12,345,678,901",
		},
	}

	for _, tst := range tests {
		r := comma(tst.input)
		if r != tst.output {
			t.Errorf("expected %s, got %s", tst.output, r)
		}
	}

}
