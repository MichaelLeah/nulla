package nulla

import (
	"testing"
)

type testSuite struct {
	Input          int
	ExpectedResult string
}

var testCases = []testSuite{
	{1, "I"},
	{5, "V"},
	{10, "X"},
	{20, "XX"},
	{3999, "MMMCMXCIX"},
}

func TestNewGeneratesNullaStructWithMinAndMax(t *testing.T) {
	var min, max int = 1, 100
	n := New(min, max)

	if n.Minimum != min {
		t.Error("given minimum:", min, "got:", n.Minimum)
	}

	if n.Maximum != max {
		t.Error("given maximum:", max, "got:", n.Maximum)
	}
}

func TestGenerateGetsExpectedResultsWithinValidRange(t *testing.T) {
	n := New(1, 3999)
	for i := range testCases {
		numeral, err := n.Generate(testCases[i].Input)
		if err != nil {
			t.Errorf("given input: %v, expected: %v, got: %v", testCases[i].Input, testCases[i].ExpectedResult, err)
		}

		if numeral != testCases[i].ExpectedResult {
			t.Errorf("given input: %v, expected: %v, got: %v", testCases[i].Input, testCases[i].ExpectedResult, numeral)
		}
	}
}

func TestGenerateFailsEarlyWhenGivenOutOfRangeValue(t *testing.T) {
	n := New(100, 101)
	for i := range testCases {
		numeral, err := n.Generate(testCases[i].Input)
		if err == nil {
			t.Errorf("given input: %v, expected: %v, got: %v", testCases[i].Input, "Out of bounds error", numeral)
		}
	}
}
