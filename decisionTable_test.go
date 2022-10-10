package main

import "testing"

func TestDecisionTable(t *testing.T) {
	testCase := []Case{
		{
			ID: 1,
			Input: Input{
				Sense:         -10,
				Applicability: 15,
				Time:          65,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 2,
			Input: Input{
				Sense:         70,
				Applicability: -10,
				Time:          60,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 3,
			Input: Input{
				Sense:         70,
				Applicability: 30,
				Time:          -20,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 4,
			Input: Input{
				Sense:         70,
				Applicability: 30,
				Time:          60,
			},
			Expected: "Rank C",
		},
		{
			ID: 5,
			Input: Input{
				Sense:         70,
				Applicability: 30,
				Time:          80,
			},
			Expected: "Rank C",
		},
		{
			ID: 6,
			Input: Input{
				Sense:         70,
				Applicability: 65,
				Time:          60,
			},
			Expected: "Rank C",
		},
		{
			ID: 7,
			Input: Input{
				Sense:         70,
				Applicability: 65,
				Time:          90,
			},
			Expected: "Rank C",
		},
		{
			ID: 8,
			Input: Input{
				Sense:         70,
				Applicability: 85,
				Time:          60,
			},
			Expected: "Rank C",
		},
		{
			ID: 9,
			Input: Input{
				Sense:         70,
				Applicability: 85,
				Time:          90,
			},
			Expected: "Rank C",
		},
		{
			ID: 10,
			Input: Input{
				Sense:         95,
				Applicability: 40,
				Time:          60,
			},
			Expected: "Rank C",
		},
		{
			ID: 11,
			Input: Input{
				Sense:         95,
				Applicability: 40,
				Time:          90,
			},
			Expected: "Rank C",
		},
		{
			ID: 12,
			Input: Input{
				Sense:         95,
				Applicability: 65,
				Time:          60,
			},
			Expected: "Rank C",
		},
		{
			ID: 13,
			Input: Input{
				Sense:         90,
				Applicability: 45,
				Time:          70,
			},
			Expected: "Rank C",
		},
		{
			ID: 14,
			Input: Input{
				Sense:         90,
				Applicability: 65,
				Time:          95,
			},
			Expected: "Rank B",
		},
		{
			ID: 15,
			Input: Input{
				Sense:         90,
				Applicability: 85,
				Time:          65,
			},
			Expected: "Rank C",
		},
		{
			ID: 16,
			Input: Input{
				Sense:         90,
				Applicability: 70,
				Time:          70,
			},
			Expected: "Rank B",
		},
		{
			ID: 17,
			Input: Input{
				Sense:         100,
				Applicability: 100,
				Time:          100,
			},
			Expected: "Rank A",
		},
		{
			ID: 18,
			Input: Input{
				Sense:         90,
				Applicability: 100,
				Time:          110,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 19,
			Input: Input{
				Sense:         90,
				Applicability: 110,
				Time:          20,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 20,
			Input: Input{
				Sense:         110,
				Applicability: 65,
				Time:          60,
			},
			Expected: "Input không hợp lệ",
		},
	}
	for _, tc := range testCase {
		result := SalaryRank(tc.Input.Sense, tc.Input.Applicability, tc.Input.Time)
		if result != tc.Expected {
			t.Errorf("Case: %v, Result: %q, Expected: %q", tc.ID, result, tc.Expected)
		}
	}
}
