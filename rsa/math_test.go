package rsa

import (
	"testing"
)

func TestGcd(t *testing.T) {
	testCases := []struct {
		x        int
		y        int
		expected int
	}{
		{
			2, 3, 1,
		},
		{
			2, 4, 2,
		},
		{
			7, 15, 1,
		},
		{
			7, 14, 7,
		},
		{
			16, 12, 4,
		},
		{
			0, 3, 3,
		},
		{
			3, 0, 3,
		},
		{
			0, 0, 0,
		},
	}

	for _, testCase := range testCases {
		val := gcd(testCase.x, testCase.y)
		if val != testCase.expected {
			t.Errorf("gcd(%d, %d) != %d; got=%d", testCase.x, testCase.y, testCase.expected, val)
		}
	}
}

func TestModInverse(t *testing.T) {
	testCases := []struct {
		x        int
		y        int
		expected int
	}{
		{
			17, 3120, 2753,
		},
		{
			3, 26, 9,
		},
		{
			16, 45, 31,
		},
	}

	for _, testCase := range testCases {
		val := modInverse(testCase.x, testCase.y)
		if val != testCase.expected {
			t.Errorf("modInverse(%d, %d) != %d; got=%d", testCase.x, testCase.y, testCase.expected, val)
		}
	}
}

func TestNthPrime(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{
			1, 2,
		},
		{
			2, 3,
		},
		{
			3, 5,
		},
		{
			9, 23,
		},
		{
			20, 71,
		},
	}

	for _, testCase := range testCases {
		val := NthPrime(testCase.n)
		if val != testCase.expected {
			t.Errorf("nthPrime(%d) != %d; got=%d", testCase.n, testCase.expected, val)
		}
	}
}

func TestIsPrime(t *testing.T) {
	testCases := []struct {
		n        int
		expected bool
	}{
		{
			-1, false,
		},
		{
			1, false,
		},
		{
			2, true,
		},
		{
			3, true,
		},
		{
			9, false,
		},
		{
			20, false,
		},
	}

	for _, testCase := range testCases {
		val := isPrime(testCase.n)
		if val != testCase.expected {
			t.Errorf("isPrime(%d) != %T; got=%T", testCase.n, testCase.expected, val)
		}
	}
}
