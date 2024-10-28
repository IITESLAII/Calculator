package main

import (
	"testing"
)

func TestCalc(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     string
		expectedResult float64
	}{
		{
			name:           "simple",
			expression:     "1+1",
			expectedResult: 2,
		},
		{
			name:           "priority",
			expression:     "(2+2)*2",
			expectedResult: 8,
		},
		{
			name:           "mixed operations",
			expression:     "2+2*2",
			expectedResult: 6,
		},
		{
			name:           "division",
			expression:     "1/2",
			expectedResult: 0.5,
		},
	}

	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := Calc(testCase.expression)
			if err != nil {
				t.Fatalf("successful case %s returns error: %v", testCase.expression, err)
			}
			if val != testCase.expectedResult {
				t.Fatalf("%f should be equal %f", val, testCase.expectedResult)
			}
		})
	}

	testCasesFail := []struct {
		name       string
		expression string
	}{
		{
			name:       "missing operand",
			expression: "1+1*",
		},
		{
			name:       "incorrect syntax",
			expression: "((2+2-*(2",
		},
		{
			name:       "consecutive operators",
			expression: "2++2",
		},
		{
			name:       "division by zero",
			expression: "1/0",
		},
		{
			name:       "unmatched parentheses",
			expression: "(1+2))",
		},
	}

	for _, testCase := range testCasesFail {
		t.Run(testCase.name, func(t *testing.T) {
			_, err := Calc(testCase.expression)
			if err == nil {
				t.Fatalf("expected an error for input: %s", testCase.expression)
			}
		})
	}
}
