package main

import (
	"fmt"
	"math"
)

// the colleciton of digits starts with the most significant digit.
//
// []digit{1,2,3} represents 123, not 321
type number struct {
	slicePos int
	digits   []digit
}

func newNumber(length int) number {
	return number{
		digits: make([]digit, length),
	}
}

func (n *number) setSlicePos(x int) {
	n.slicePos = x
}

// proposeFirst attempts to add the provided number (x) into the least
// significant digit of the first subset
//
// Returns true if adding this number causes no conflict, or if there is not
// enough information to start evaluating this number (e.g. this might happen
// if we've started to propose only one of the two subsets
//
// Returns false if this number would cause a contradiction moving forward
func (n *number) proposeFirst(x int) bool {
	i := n.slicePos - 1
	for {
		if i < 0 {
			return false
		}

		if !n.digits[i].known {
			n.digits[i] = n.digits[i].set(x)
			return n.fillWithProduct()
		}

		i--
	}
}

// proposeSecond attempts to add the provided number (x) into the least
// significant digit of the second subset
//
// Returns true if adding this number causes no conflict, or if there is not
// enough information to start evaluating this number (e.g. this might happen
// if we've started to propose only one of the two subsets
//
// Returns false if this number would cause a contradiction moving forward
func (n *number) proposeSecond(x int) bool {
	i := len(n.digits) - 1
	for {
		if i < n.slicePos {
			return false
		}

		if !n.digits[i].known {
			n.digits[i] = n.digits[i].set(x)
			return n.fillWithProduct()
		}

		i--
	}
}

func (n *number) fillWithProduct() bool {
	product := n.productSoFar()
	if len(product) > len(n.digits) {
		// Product is larger than the results being multiplied, so we can
		// assume this is a bad combination
		return false
	}

	for i := 0; i < len(product); i++ {
		// The product can't end in a 0, because that would been the reversed
		// product can't begin with a 0 (which would change the number of characters
		if i == 0 && product[i] == 0 {
			return false
		}
		if n.digits[i].known {
			if n.digits[i].value != product[i] {
				return false
			}

		} else {
			n.digits[i] = n.digits[i].set(product[i])
		}
	}

	return true
}

// productSoFar returns the product of both subsets, but limited to the
// number of digits of the minimum side (if both sides still have unknown values)
//
// For example, if the two subsets are '1234' and '56' are the two subsets,
// the result will be '04'. The product of the two is 69104, and limited to
// two digits, we get '04'
func (n *number) productSoFar() []int {

	first := n.firstInt()
	second := n.secondInt()

	digitsToReturn := n.establishedDigits()
	var resArray []int
	if digitsToReturn == 0 {
		return resArray
	}

	total := first * second
	for i := 0; i < digitsToReturn; i++ {
		val := total % 10
		total /= 10
		resArray = append(resArray, val)
	}

	return resArray
}

func (n *number) productInt() int {
	return n.firstInt() * n.secondInt()
}

// check returns true if the two subsets multiply into each other resulting
// in the reverse order of the digits (i.e. the number is cranky)
func (n *number) check() bool {
	if n.hasMore() {
		return false
	}

	return n.fillWithProduct()
}

func (n *number) firstInt() int {
	res := 0
	i := n.slicePos - 1
	exp := 0
	for i-exp >= 0 && n.digits[i-exp].known {
		res += n.digits[i-exp].value * int(math.Pow10(exp))
		exp++
	}
	return res
}

func (n *number) secondInt() int {
	res := 0
	i := len(n.digits) - 1
	exp := 0
	for i-exp >= n.slicePos && n.digits[i-exp].known {
		res += n.digits[i-exp].value * int(math.Pow10(exp))
		exp++
	}
	return res
}

// Returns the minimum number of digits represented by either subset.
//
// There are a couple rules that determine what value will be found here:
//  - If either side still has unknown values, return the minimum length of
//    the two known subsets
//  - If one side has no more unknown values, but the other does, return the
//    size of the subset with more to explore
//  - If both subsets are complete, return the length of the full number
func (n *number) establishedDigits() int {

	firstCount := 0
	for i := n.slicePos - 1; i >= 0 && n.digits[i].known; i-- {
		firstCount++
	}
	firstHasMore := (firstCount != n.slicePos)

	secondCount := 0
	for i := len(n.digits) - 1; i >= n.slicePos && n.digits[i].known; i-- {
		secondCount++
	}
	secondHasMore := (secondCount != len(n.digits)-n.slicePos)

	if !firstHasMore && !secondHasMore {
		return len(n.digits)
	}

	if firstHasMore && !secondHasMore {
		return firstCount
	}

	if !firstHasMore && secondHasMore {
		return secondCount
	}

	if firstCount < secondCount {
		return firstCount
	}

	return secondCount
}

func (n *number) summary() string {

	var first string
	sub1 := n.digits[0:n.slicePos]

	for i := 0; i < len(sub1); i++ {
		if sub1[i].known {
			first += fmt.Sprintf("%d", sub1[i].value)
		} else {
			first += "x"
		}
	}

	var second string
	sub2 := n.digits[n.slicePos:len(n.digits)]

	for i := 0; i < len(sub2); i++ {
		if sub2[i].known {
			second += fmt.Sprintf("%d", sub2[i].value)
		} else {
			second += "x"
		}
	}

	return fmt.Sprintf("[%s][%s]", first, second)
}

func (n *number) propose(x int) bool {

	firstCount := 0
	for i := n.slicePos - 1; i >= 0 && n.digits[i].known; i-- {
		firstCount++
	}
	firstHasMore := (firstCount != n.slicePos)

	secondCount := 0
	for i := len(n.digits) - 1; i >= n.slicePos && n.digits[i].known; i-- {
		secondCount++
	}
	secondHasMore := (secondCount != len(n.digits)-n.slicePos)

	if firstHasMore && secondHasMore {
		if firstCount < secondCount {
			return n.proposeFirst(x)
		}

		return n.proposeSecond(x)
	}

	if !firstHasMore {
		return n.proposeSecond(x)
	}

	return n.proposeFirst(x)
}

func (n *number) copy() number {
	var newDigits []digit
	for _, d := range n.digits {
		newDigits = append(newDigits, d)
	}
	return number{
		slicePos: n.slicePos,
		digits:   newDigits,
	}
}

func (n *number) hasMore() bool {
	for _, d := range n.digits {
		if !d.known {
			return true
		}
	}
	return false
}

// returns the original, full number as an integer
func (n *number) asInt() int {
	num := 0

	for i := 0; i < len(n.digits); i++ {
		num *= 10
		num += n.digits[i].value
	}

	return num
}
