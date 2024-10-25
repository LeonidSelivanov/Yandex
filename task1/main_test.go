package task

import (
	"errors"
	"testing"
)

type Test struct {
	in  string
	out float64
	err error
}

func TestCalc(t *testing.T) {
	tests := []Test{
		{
			in:  "1+1",
			out: 2,
			err: nil,
		},
		{
			in:  "5*5-(5*5)",
			out: 0,
			err: nil,
		},
		{
			in:  "",
			out: 0,
			err: errors.New("ErrFoo Emmpty expression"),
		},
	}
	for _, test := range tests {
		gotOut, gotErr := test.out, test.err
		expectedOut, expectedOutErr := Calc(test.in)
		if gotOut != expectedOut || gotErr != expectedOutErr {
			t.Fatalf("Error!")
		}
	}
}
