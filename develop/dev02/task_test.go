package main

import "testing"

func Test_extract(t *testing.T) {
	var table = []struct {
		input       string
		expectedOut string
		err         bool
	}{
		{
			input:       `a4bc2d5e`,
			expectedOut: `aaaabccddddde`,
			err:         false,
		},
		{
			input:       `abcd`,
			expectedOut: `abcd`,
			err:         false,
		},
		{
			input:       `qwe\45`,
			expectedOut: `qwe44444`,
			err:         false,
		},
		{
			input:       `45`,
			expectedOut: ``,
			err:         true,
		},
	}

	for _, test := range table {
		out, err := unpack(test.input)
		if out != test.expectedOut || err == nil && test.err || err != nil && !test.err {
			t.Error("Error in unpacking")
		}
	}
}
