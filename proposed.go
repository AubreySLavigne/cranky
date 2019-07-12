package main

type proposed int64

// 12345
// Reversed => 54321
// Splits => [(1,2345), (12,345), (123,45), (1234,5)]
// Sum of Splits => [2345, 4140, 5535, 6170]
// Is Cranky? => false

func (p proposed) reversed() int64 {
	tmp := int64(p)
	res := int64(0)
	for tmp > 0 {
		res *= 10
		i := tmp % 10
		res += i
		tmp /= 10
	}
	return res
}

func (p proposed) splits() [][2]int64 {

	res := [][2]int64{}
	tmp := 10

	for tmp <= int(p) {
		res = append(res, [2]int64{int64(int(p) / tmp), int64(int(p) % tmp)})
		tmp *= 10
	}

	return res
}

func (p proposed) sums() []int64 {
	splits := p.splits()
	res := []int64{}
	for _, split := range splits {
		res = append(res, split[0]*split[1])
	}
	return res
}

func (p proposed) isCranky() bool {
	rev := p.reversed()
	sums := p.sums()
	for _, sum := range sums {
		if sum == rev {
			return true
		}
	}
	return false
}
