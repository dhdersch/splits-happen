package main

import (
	"fmt"
	"testing"
)

func TestSplitsScore(t *testing.T) {

	testCases := []struct {
		input          string
		expectedResult int
	}{
		{input: "XXXXXXXXXXXX", expectedResult: 300}, // edge case
		{input: "9-9-9-9-9-9-9-9-9-9-", expectedResult: 90},
		{input: "5/5/5/5/5/5/5/5/5/5/5", expectedResult: 150},
		{input: "X7/9-X-88/-6XXX81", expectedResult: 167},
		{input: "239--1X4/-/9/527/71", expectedResult: 111},
		{input: "--------------------", expectedResult: 0}, // edge case
		{input: "-/X-/-/-/-/-/-/-/X9/", expectedResult: 140},
		{input: "------------------9/X", expectedResult: 20},
		{input: "-123456-12345-12345-", expectedResult: 51},
		{input: "-/X9/5--/-/-/-/-/-/-", expectedResult: 120},
	}

	var errors []error
	for _, testCase := range testCases {
		actualResult := SplitsScore(testCase.input)
		if actualResult != testCase.expectedResult {
			errors = append(errors, fmt.Errorf("input: '%s', expected score: %d, actual score: %d",
				testCase.input, testCase.expectedResult, actualResult))
		}
	}
	if len(errors) > 0 {
		t.Logf("%d test cases failed, %d test cases passed", len(errors), len(testCases)-len(errors))
		for _, err := range errors {
			t.Log(err.Error())
		}
		t.Fail()
	} else {
		t.Logf("%d test cases passed", len(testCases))
	}

}
