package main

import (
	"testing"
)

type step struct {
	whichSubset       string
	valueToAdd        int
	expectedToSucceed bool
	expectedDigits    []digit
}

// TestPropose tests number.proposeFirst() and number.proposeSecond()
func TestPropose(t *testing.T) {

	tests := []struct {
		length   int
		slicePos int
		steps    []step
	}{
		{
			length:   3,
			slicePos: 1,
			steps: []step{
				{
					whichSubset:       "first",
					valueToAdd:        3,
					expectedToSucceed: true,
					expectedDigits:    []digit{{value: 3, known: true}, {known: false}, {known: false}},
				},
				{
					whichSubset:       "second",
					valueToAdd:        1,
					expectedToSucceed: true,
					expectedDigits:    []digit{{value: 3, known: true}, {known: false}, {value: 1, known: true}},
				},
				{
					whichSubset:       "second",
					valueToAdd:        5,
					expectedToSucceed: true,
					expectedDigits:    []digit{{value: 3, known: true}, {value: 5, known: true}, {value: 1, known: true}},
				},
			},
		},
		{
			length:   3,
			slicePos: 1,
			steps: []step{
				{
					whichSubset:       "first",
					valueToAdd:        3,
					expectedToSucceed: true,
					expectedDigits:    []digit{{value: 3, known: true}, {known: false}, {known: false}},
				},
				{
					whichSubset:       "second",
					valueToAdd:        1,
					expectedToSucceed: true,
					expectedDigits:    []digit{{value: 3, known: true}, {known: false}, {value: 1, known: true}},
				},
				{
					whichSubset:       "second",
					valueToAdd:        6,
					expectedToSucceed: false,
					expectedDigits:    []digit{{value: 3, known: true}, {value: 6, known: true}, {value: 1, known: true}},
				},
			},
		},
	}

	for _, test := range tests {

		num := newNumber(test.length)
		num.setSlicePos(test.slicePos)

		for j, step := range test.steps {

			var succeed bool
			if step.whichSubset == "first" {
				succeed = num.proposeFirst(step.valueToAdd)
			} else {
				succeed = num.proposeSecond(step.valueToAdd)
			}

			if succeed != step.expectedToSucceed {
				t.Errorf("Step #%d: Proposing '%d' doesn't match expected outcome. Got %t, expected %t.", j, step.valueToAdd, succeed, step.expectedToSucceed)
			}

			for k, expected := range step.expectedDigits {

				if expected.known {

					if num.digits[k].known == false {
						t.Errorf("Digit #%d should not be known", k)
					} else if num.digits[k].value != expected.value {
						t.Errorf("Digit #%d value did not match. For value, Got %d, Expected %d", k, num.digits[0].value, expected.value)
					}

				} else {

					// Expected not to be known
					if num.digits[k].known {
						t.Errorf("Digit #%d should not be known", k)
					}

				}
			}

			if t.Failed() {
				t.FailNow()
			}
		}

	}
}

func TestProductSoFar(t *testing.T) {
	tests := []struct {
		name     string
		first    []int
		second   []int
		expected []int
	}{
		{name: "Only First",
			first: []int{3}, second: []int{}, expected: []int{}},
		{name: "Only Second",
			first: []int{}, second: []int{4}, expected: []int{}},
		{name: "Single Digit Product",
			first: []int{2}, second: []int{3}, expected: []int{6}},
		{name: "Multi Digit Product expecting single digit",
			first: []int{6}, second: []int{7}, expected: []int{2}},
		{name: "Multi Digit Product expecting two digits",
			first: []int{7, 2}, second: []int{5, 3}, expected: []int{5, 4}},
		{name: "Take the smaller",
			first: []int{6, 5, 4, 3, 2, 1}, second: []int{9, 8, 7}, expected: []int{4, 8, 7}},
	}

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			num := newNumber(14)
			num.setSlicePos(7)

			for _, i := range test.first {
				num.proposeFirst(i)
			}
			for _, i := range test.second {
				num.proposeSecond(i)
			}

			res := num.productSoFar()
			if len(res) != len(test.expected) {
				t.Errorf("product does not match expected. Got %v, expected %v", res, test.expected)
			} else {
				for i, val := range res {
					if val != test.expected[i] {
						t.Errorf("product does not match expected. Got %v, expected %v", res, test.expected)
					}
				}
			}
		})
	}
}

func TestFirstInt(t *testing.T) {

	tests := []struct {
		input    []int
		expected int
	}{
		{input: []int{}, expected: 0},
		{input: []int{0}, expected: 0},
		{input: []int{1}, expected: 1},
		{input: []int{1, 2}, expected: 21},
	}

	for _, test := range tests {

		num := newNumber(14)
		num.setSlicePos(7)

		for _, newNum := range test.input {
			num.proposeFirst(newNum)
		}

		if res := num.firstInt(); res != test.expected {
			t.Errorf("Resulting integer does not match expected. Got %d, Expected %d", res, test.expected)
		}

	}
}

func TestSecondInt(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{input: []int{}, expected: 0},
		{input: []int{0}, expected: 0},
		{input: []int{1}, expected: 1},
		{input: []int{1, 2}, expected: 21},
	}

	for _, test := range tests {

		num := newNumber(14)
		num.setSlicePos(7)

		for _, newNum := range test.input {
			num.proposeSecond(newNum)
		}

		if res := num.secondInt(); res != test.expected {
			t.Errorf("Resulting integer does not match expected. Got %d, Expected %d", res, test.expected)
		}

	}
}

func TestEstablishedDigits(t *testing.T) {
	tests := []struct {
		firstInput  []int
		secondInput []int
		expected    int
	}{
		{firstInput: []int{}, secondInput: []int{}, expected: 0},
		{firstInput: []int{1, 2}, secondInput: []int{}, expected: 0},
		{firstInput: []int{}, secondInput: []int{3, 2, 1}, expected: 0},
		{firstInput: []int{1, 2, 3, 4}, secondInput: []int{5, 4, 3, 2, 1}, expected: 4},
	}

	for _, test := range tests {

		num := newNumber(20)
		num.setSlicePos(10)

		for _, newNum := range test.firstInput {
			num.proposeFirst(newNum)
		}
		for _, newNum := range test.secondInput {
			num.proposeSecond(newNum)
		}

		if res := num.establishedDigits(); res != test.expected {
			t.Errorf("Number of effective digits does not match expected. Got %d, Expected %d", res, test.expected)
		}

	}
}
