package main

import "testing"

func TestEquivalence(t *testing.T) {
	testCase := []Case{
		{
			ID: 1,
			Input: Input{
				Sense:         -10,
				Applicability: -10,
				Time:          -20,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 2,
			Input: Input{
				Sense:         -10,
				Applicability: -10,
				Time:          60,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 3,
			Input: Input{
				Sense:         -10,
				Applicability: -10,
				Time:          90,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 4,
			Input: Input{
				Sense:         -10,
				Applicability: 40,
				Time:          -20,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 5,
			Input: Input{
				Sense:         -10,
				Applicability: 40,
				Time:          60,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 6,
			Input: Input{
				Sense:         -10,
				Applicability: 40,
				Time:          90,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 7,
			Input: Input{
				Sense:         -10,
				Applicability: 65,
				Time:          -20,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 8,
			Input: Input{
				Sense:         -10,
				Applicability: 65,
				Time:          60,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 9,
			Input: Input{
				Sense:         -10,
				Applicability: 65,
				Time:          90,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 10,
			Input: Input{
				Sense:         -10,
				Applicability: 85,
				Time:          -20,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 11,
			Input: Input{
				Sense:         -10,
				Applicability: 85,
				Time:          60,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 12,
			Input: Input{
				Sense:         -10,
				Applicability: 85,
				Time:          90,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 13,
			Input: Input{
				Sense:         70,
				Applicability: -10,
				Time:          -20,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 14,
			Input: Input{
				Sense:         70,
				Applicability: -10,
				Time:          60,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 15,
			Input: Input{
				Sense:         70,
				Applicability: -10,
				Time:          90,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 16,
			Input: Input{
				Sense:         70,
				Applicability: 40,
				Time:          -20,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 17,
			Input: Input{
				Sense:         70,
				Applicability: 40,
				Time:          60,
			},
			Expected: "Rank C",
		},
		{
			ID: 18,
			Input: Input{
				Sense:         70,
				Applicability: 40,
				Time:          90,
			},
			Expected: "Rank C",
		},
		{
			ID: 19,
			Input: Input{
				Sense:         70,
				Applicability: 65,
				Time:          -20,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 20,
			Input: Input{
				Sense:         70,
				Applicability: 65,
				Time:          60,
			},
			Expected: "Rank C",
		},
		{
			ID: 21,
			Input: Input{
				Sense:         70,
				Applicability: 65,
				Time:          90,
			},
			Expected: "Rank C",
		},
		{
			ID: 22,
			Input: Input{
				Sense:         70,
				Applicability: 85,
				Time:          -20,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 23,
			Input: Input{
				Sense:         70,
				Applicability: 85,
				Time:          60,
			},
			Expected: "Rank C",
		},
		{
			ID: 24,
			Input: Input{
				Sense:         70,
				Applicability: 85,
				Time:          90,
			},
			Expected: "Rank C",
		},
		{
			ID: 25,
			Input: Input{
				Sense:         90,
				Applicability: -10,
				Time:          -20,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 26,
			Input: Input{
				Sense:         90,
				Applicability: -10,
				Time:          60,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 27,
			Input: Input{
				Sense:         90,
				Applicability: -10,
				Time:          90,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 28,
			Input: Input{
				Sense:         90,
				Applicability: 40,
				Time:          -20,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 29,
			Input: Input{
				Sense:         90,
				Applicability: 40,
				Time:          60,
			},
			Expected: "Rank C",
		},
		{
			ID: 30,
			Input: Input{
				Sense:         90,
				Applicability: 40,
				Time:          90,
			},
			Expected: "Rank C",
		},
		{
			ID: 31,
			Input: Input{
				Sense:         90,
				Applicability: 65,
				Time:          -20,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 32,
			Input: Input{
				Sense:         90,
				Applicability: 65,
				Time:          60,
			},
			Expected: "Rank C",
		},
		{
			ID: 33,
			Input: Input{
				Sense:         90,
				Applicability: 65,
				Time:          90,
			},
			Expected: "Rank B",
		},
		{
			ID: 34,
			Input: Input{
				Sense:         90,
				Applicability: 85,
				Time:          -20,
			},
			Expected: "Input không hợp lệ",
		},
		{
			ID: 35,
			Input: Input{
				Sense:         90,
				Applicability: 85,
				Time:          60,
			},
			Expected: "Rank C",
		},
		{
			ID: 36,
			Input: Input{
				Sense:         90,
				Applicability: 85,
				Time:          90,
			},
			Expected: "Rank A",
		},
	}
	for _, tc := range testCase {
		result := SalaryRank(tc.Input.Sense, tc.Input.Applicability, tc.Input.Time)
		if result != tc.Expected {
			t.Errorf("Case: %v, Result: %q, Expected: %q", tc.ID, result, tc.Expected)
		}
	}
}
