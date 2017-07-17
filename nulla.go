package nulla

import (
	"fmt"
	"strings"
)

// RomanNumeral is a representation of a Roman Numeral's string representation and its value.
type RomanNumeral struct {
	Numeral string
	Value   int
}

// Numerals contains an array of RomanNumeral instances to act as a conversion table.
var Numerals = []RomanNumeral{
	{"M", 1000},
	{"CM", 900},
	{"D", 500},
	{"CD", 400},
	{"C", 100},
	{"XC", 90},
	{"L", 50},
	{"XL", 40},
	{"X", 10},
	{"IX", 9},
	{"V", 5},
	{"IV", 4},
	{"I", 1},
}

// nulla implements the RomanNumeralGenerator interface and allows conversion of ints to the roman numeral string representation.
// It also allows a minimum and maximum range to be specified if desired.
type nulla struct {
	Minimum int
	Maximum int
}

// New is a basic constructor method to enforce minimum and maximum ranges to be set on the nulla struct
func New(min, max int) nulla {
	return nulla{min, max}
}

// checkRestrictions ensures that a given number is within the parameter range specified on the nulla struct
func (n *nulla) checkRestrictions(number int) error {
	if number < n.Minimum || number > n.Maximum {
		return fmt.Errorf("number is out of bounds, accepted range [%v - %v], given number: %v", n.Minimum, n.Maximum, number)
	}

	return nil
}

// Generate converts an integer into the roman numeral string representation.
func (n *nulla) Generate(number int) (string, error) {
	if err := n.checkRestrictions(number); err != nil {
		return "", fmt.Errorf("error when using input: %v", err)
	}
	var output string
	for i := range Numerals {
		count := number / Numerals[i].Value
		output = strings.Join([]string{output, strings.Repeat(Numerals[i].Numeral, count)}, "")
		number %= Numerals[i].Value
	}

	return output, nil
}
