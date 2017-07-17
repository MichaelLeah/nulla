package nulla

// RomanNumeralGenerator contains one method which is for converting an integer to the string representation of a Roman Numeral.
type RomanNumeralGenerator interface {
	Generate(number int) string
}
