package solve

import "testing"

type inputParameters struct {
	productName string
	amount      float64
	phoneNumber string
	duration    int
}

type testCase struct {
	description string
	input       inputParameters
	expected    float64
}

var testCases = []testCase{
	// for phone
	{
		description: "Test 1 -> Phone",
		input:       inputParameters{"phone", 1000, "+992 93 9965556", 18},
		expected:    1060,
	},
	{
		description: "Test 2 -> Phone",
		input:       inputParameters{"phone", 750, "+992 93 9965556", 20},
		expected:    817.5,
	},
	{
		description: "Test 3 -> Phone",
		input:       inputParameters{"phone", 20250, "+992 93 9965556", 4},
		expected:    20857.5,
	},

	// for computer
	{
		description: "Test 1 -> Computer",
		input:       inputParameters{"computer", 1000, "+992 93 9965556", 18},
		expected:    1080,
	},
	{
		description: "Test 2 -> Computer",
		input:       inputParameters{"computer", 750, "+992 93 9965556", 20},
		expected:    810,
	},
	{
		description: "Test 3 -> Computer",
		input:       inputParameters{"computer", 20250, "+992 93 9965556", 4},
		expected:    21060,
	},

	// for television
	{
		description: "Test 1 -> Television",
		input:       inputParameters{"television", 1000, "+992 93 9965556", 18},
		expected:    1050,
	},
	{
		description: "Test 2 -> Computer",
		input:       inputParameters{"television", 750, "+992 93 9965556", 20},
		expected:    825,
	},
	{
		description: "Test 3 -> Computer",
		input:       inputParameters{"television", 20250, "+992 93 9965556", 4},
		expected:    21262.5,
	},

	// wrong name
	{
		description: "Test 1 -> Wrong",
		input:       inputParameters{"", 20250, "+992 93 9965556", 4},
		expected:    0,
	},
	{
		description: "Test 2 -> Wrong",
		input:       inputParameters{"computer58653", 20250, "+992 93 9965556", 4},
		expected:    0,
	},
	{
		description: "Test 3 -> Wrong",
		input:       inputParameters{"phone", -100, "+992 93 9965556", 4},
		expected:    0,
	},
	{
		description: "Test 4 -> Wrong",
		input:       inputParameters{"phone", 100, "+992 93 9965556", -4},
		expected:    0,
	},
}

func TestTotalAmount(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := TotalAmount(tc.input.productName, tc.input.amount, tc.input.phoneNumber, tc.input.duration)
			if got != tc.expected {
				t.Fatalf("TotalAmount function failed!\ngot: %f\nwant: %f", got, tc.expected)
			}
		})
	}
}
