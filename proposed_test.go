package main

import (
	"fmt"
	"testing"
)

func TestProposed(t *testing.T) {
	tests := []struct {
		name     string
		input    int64
		reversed int64
		splits   [][2]int64
		sums     []int64
		isCranky bool
	}{
		{
			name:     "Provided 351",
			input:    351,
			reversed: 153,
			splits: [][2]int64{
				[2]int64{35, 1},
				[2]int64{3, 51},
			},
			sums: []int64{
				35,
				153,
			},
			isCranky: true,
		},
		{
			name:     "Provided 425322",
			input:    425322,
			reversed: 223524,
			splits: [][2]int64{
				[2]int64{42532, 2},
				[2]int64{4253, 22},
				[2]int64{425, 322},
				[2]int64{42, 5322},
				[2]int64{4, 25322},
			},
			sums: []int64{
				42532 * 2,
				4253 * 22,
				425 * 322,
				42 * 5322,
				4 * 25322,
			},
			isCranky: true,
		},
		{
			name:     "Provided 57982563",
			input:    57982563,
			reversed: 36528975,
			splits: [][2]int64{
				[2]int64{5798256, 3},
				[2]int64{579825, 63},
				[2]int64{57982, 563},
				[2]int64{5798, 2563},
				[2]int64{579, 82563},
				[2]int64{57, 982563},
				[2]int64{5, 7982563},
			},
			sums: []int64{
				5798256 * 3,
				579825 * 63,
				57982 * 563,
				5798 * 2563,
				579 * 82563,
				57 * 982563,
				5 * 7982563,
			},
			isCranky: true,
		},
		{
			name:     "NonCranky 12345",
			input:    123456789,
			reversed: 987654321,
			splits: [][2]int64{
				[2]int64{12345678, 9},
				[2]int64{1234567, 89},
				[2]int64{123456, 789},
				[2]int64{12345, 6789},
				[2]int64{1234, 56789},
				[2]int64{123, 456789},
				[2]int64{12, 3456789},
				[2]int64{1, 23456789},
			},
			sums: []int64{
				12345678 * 9,
				1234567 * 89,
				123456 * 789,
				12345 * 6789,
				1234 * 56789,
				123 * 456789,
				12 * 3456789,
				1 * 23456789,
			},
			isCranky: false,
		},
	}

	for _, expected := range tests {
		t.Run(expected.name, func(t *testing.T) {
			num := proposed(expected.input)

			if val := num.reversed(); val != expected.reversed {
				t.Errorf("Reversed Number doesn't match. Got %d, Expected %d", val, expected.reversed)
			}

			splits := num.splits()
			if len(splits) != len(expected.splits) {
				t.Errorf("List of Splits doesn't match. Got %v, Expected %v", splits, expected.splits)
			} else {
				for row, split := range splits {
					if split[0] != expected.splits[row][0] || split[1] != expected.splits[row][1] {
						t.Errorf("Index %d  does not match between these splits. Got %v, Expected %v", row, splits, expected.splits)
					}
				}
			}

			sums := num.sums()
			if len(sums) != len(expected.sums) {
				t.Errorf("List of Sums doesn't match. Got %v, Expected %v", sums, expected.sums)
			} else {
				for row, sum := range sums {
					if sum != expected.sums[row] {
						t.Errorf("Index %d  does not match between these sums. Got %v, Expected %v", row, sums, expected.sums)
					}
				}
			}

			isCranky := num.isCranky()
			if isCranky != expected.isCranky {
				t.Errorf("Did not correctly evaluate whether cranky. Got %v, Expected %v", isCranky, expected.isCranky)
			}
		})
	}
}

func BenchmarkCranky(b *testing.B) {
	for n := 0; n < b.N; n++ {
		num := proposed(n)
		if num.isCranky() {
			fmt.Println(num)
		}
	}
}
